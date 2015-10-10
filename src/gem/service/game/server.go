package game

import (
	"fmt"
	"net"
	"sync"

	"gem/auth"
	"gem/crypto"
	"gem/encoding"
	"gem/log"
	"gem/protocol"
	"gem/runite"

	"github.com/qur/gopy/lib"
	tomb "gopkg.in/tomb.v2"
)

var logInit sync.Once
var logger *log.Module

type Index int

//go:generate gopygen -type Server -excfunc "^[a-z].+" -excfield "^[a-z].+" $GOFILE
type Server struct {
	py.BaseObject

	laddr string
	ln    net.Listener

	update    *updateService
	game      *gameService
	runite    *runite.Context
	nextIndex chan Index

	m       sync.Mutex
	clients map[Index]*Connection

	t tomb.Tomb
}

// Start creates the tcp listener and starts the connection handler in a goroutine
func (s *Server) Start(laddr string, ctx *runite.Context, rsaKeyPath string, auth auth.Provider) error {
	logInit.Do(func() {
		logger = log.New("game")
	})

	logger.Info("Starting game server...")

	var err error
	var key *crypto.Keypair
	key, err = crypto.LoadPrivateKey(rsaKeyPath)
	if err != nil {
		return err
	}
	logger.Infof("Loaded RSA keypair %v", rsaKeyPath)

	s.laddr = laddr
	s.runite = ctx
	s.clients = make(map[Index]*Connection)
	s.update = newUpdateService(ctx)
	s.game = newGameService(ctx, key, auth)
	go s.update.processQueue()

	s.ln, err = net.Listen("tcp", s.laddr)
	if err != nil {
		return fmt.Errorf("couldn't start game server: %v", err)
	}

	s.nextIndex = make(chan Index)
	go s.generateNextIndex()

	s.t.Go(s.run)
	return nil
}

func (s *Server) generateNextIndex() {
	index := Index(0)
	for {
		s.nextIndex <- index
		index++
	}
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

	s.m.Lock()
	// Close any existing connections
	for _, conn := range s.clients {
		conn.Disconnect()
	}
	s.m.Unlock()
	wg.Wait()

	logger.Noticef("Shut down")
	return nil
}

// handle is the per-connection i/o goroutine.
// buffers data into readBuffer and flushes data from writeBuffer.
// if the disconnect channel is signalled, breaks the main loop and closes the connection
func (s *Server) handle(netConn net.Conn) {
	conn := newConnection(netConn, logger)
	conn.decode = s.handshake
	s.registerClient(conn)
	defer s.unregisterClient(conn)

	defer conn.recover()

	conn.Log.Info("accepted connection")

	connCtx := &context{
		conn:   conn,
		update: s.update,
		game:   s.game,
	}

	go conn.encodeFromWriteQueue(connCtx)
	go conn.decodeToReadQueue(connCtx)

	// Block this thread until disconnect
	<-conn.disconnect

	// ensure any pending data is flushed before disconnecting
	conn.flushWriteBuffer()

	close(conn.read)
	close(conn.write)
	conn.conn.Close()
	conn.Log.Info("connection closed")
}

// registerClient adds a connection to the clients map
func (s *Server) registerClient(conn *Connection) {
	s.m.Lock()
	defer s.m.Unlock()

	index := <-s.nextIndex
	s.clients[index] = conn
	conn.Index = index
}

// unregisterClient removes a connection to the clients map
func (s *Server) unregisterClient(conn *Connection) {
	s.m.Lock()
	defer s.m.Unlock()

	delete(s.clients, conn.Index)
	index := <-s.nextIndex
	s.clients[index] = conn
}

// handshake reads the service selection byte and points the connection's decode func
// towards the decode func for the selected service
func (s *Server) handshake(ctx *context, b *encoding.Buffer) error {
	var svc protocol.ServiceSelect
	if err := svc.Decode(b, nil); err != nil {
		return err
	}

	conn := ctx.conn

	switch svc.Service {
	case UpdateService:
		conn.Log.Infof("new update client")
		conn.decode = ctx.update.decodeRequest
		conn.encode = conn.encodeCodable

		conn.write <- new(protocol.UpdateHandshakeResponse)
		return nil
	case GameService:
		conn.Log.Infof("new game client")
		conn.decode = ctx.game.handshake
		conn.encode = conn.encodeCodable
		return nil
	default:
		conn.Log.Errorf("invalid service requested: %v", svc)
		conn.Disconnect()
	}

	return nil
}
