package game

import (
	"fmt"
	"net"
	"sync"

	"gem/encoding"
	"gem/log"
	"gem/runite"

	"github.com/qur/gopy/lib"
	tomb "gopkg.in/tomb.v2"
)

var logInit sync.Once
var logger *log.Module

type Index int

type context struct {
	conn   *GameConnection
	update *updateService
	game   *gameService
}

//go:generate gopygen -type Server -exclude "^[a-z].+" $GOFILE
type Server struct {
	py.BaseObject

	laddr string
	ln    net.Listener

	update    *updateService
	game      *gameService
	runite    *runite.Context
	clients   map[Index]*GameConnection
	nextIndex Index

	t tomb.Tomb
}

// Start creates the tcp listener and starts the connection handler in a goroutine
func (s *Server) Start(laddr string, ctx *runite.Context) error {
	var err error
	s.laddr = laddr
	s.runite = ctx
	s.clients = make(map[Index]*GameConnection)
	s.update = newUpdateService(ctx)
	s.game = newGameService(ctx)
	go s.update.processQueue()

	logInit.Do(func() {
		logger = log.New("game")
	})

	logger.Info("Starting game server...")

	s.ln, err = net.Listen("tcp", s.laddr)
	if err != nil {
		return fmt.Errorf("couldn't start game server: %v", err)
	}

	s.t.Go(s.run)
	return nil
}

// Stop signals that the listener thread should be stopped.
// Existing clients are forcefully disconnected. Blocks until all connections and
// the listener are closed.
func (s *Server) Stop() error {
	logger.Info("Stopping game server...")
	if s.t.Alive() {
		s.t.Kill(nil)
		return s.t.Wait()
	}
	return nil
}

// run is the main tcp handler thread. listens for new connections and starts a new goroutine
// for each connection to handle i/o
func (s *Server) run() error {
	logger.Noticef("Listening on %v", s.laddr)

	// Accept in a seperate goroutine
	accept := make(chan net.Conn, 16)
	go func() {
		for s.t.Alive() {
			conn, err := s.ln.Accept()
			if err != nil {
				if s.t.Alive() {
					logger.Errorf("error accepting client: %v", err)
				}
			}
			accept <- conn
		}
		close(accept)
	}()

	// Pull connections from the accept channel
	var wg sync.WaitGroup
	for s.t.Alive() {
		select {
		case conn := <-accept:
			wg.Add(1)
			go func() {
				defer wg.Done()
				s.handle(conn)
			}()
		case <-s.t.Dying():
			continue
		}
	}

	// Stop accepting
	s.ln.Close()

	// Close any existing connections
	for _, conn := range s.clients {
		conn.Disconnect()
	}
	wg.Wait()

	logger.Noticef("Shut down")
	return nil
}

// handle is the per-connection i/o goroutine.
// buffers data into readBuffer and flushes data from writeBuffer.
// if the disconnect channel is signalled, breaks the main loop and closes the connection
func (s *Server) handle(netConn net.Conn) {
	index := s.nextIndex
	s.nextIndex++

	conn := newConnection(index, netConn, logger)
	conn.decode = conn.handshake
	s.clients[index] = conn

	conn.Log.Info("accepted connection")

	connCtx := &context{
		conn:   conn,
		update: s.update,
		game:   s.game,
	}

	go conn.fillReadBuffer()

	// Main client loop
L:
	for {
		select {
		case <-conn.disconnect:
			break L
		case <-conn.canRead:
			// at this point, the only error should be because we didn't have enough data
			// todo: formalize this and check for the right error
			err := conn.readBuffer.Try(func(b *encoding.Buffer) error {
				return conn.decode(connCtx, b)
			})
			if err == nil {
				// We handled some data, discard it
				conn.readBuffer.Trim()
				if conn.readBuffer.Len() > 0 {
					// there's still some data buffered
					conn.canRead <- 1
				}
			}
		case <-conn.canWrite:
			conn.flushWriteBuffer()
		}
	}

	conn.Log.Info("connection closed")
	conn.conn.Close()
	delete(s.clients, index)
}
