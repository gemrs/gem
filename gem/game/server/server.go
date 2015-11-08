package server

import (
	"fmt"
	"net"
	"sync"

	"github.com/sinusoids/gem/gem/log"
	"github.com/sinusoids/gem/gem/protocol"
	"github.com/sinusoids/gem/gem/util/id"
	"github.com/sinusoids/gem/gem/util/safe"

	"github.com/qur/gopy/lib"
	tomb "gopkg.in/tomb.v2"
)

var logger = log.New("server", log.NilContext)

// Server is the listener object and its associated state
type Server struct {
	py.BaseObject

	laddr string
	ln    net.Listener

	nextIndex <-chan int

	m        sync.Mutex
	clients  map[int]Client
	services map[int]Service

	t tomb.Tomb
}

// A Service is capable of creating Clients specific to each service (game/update)
type Service interface {
	NewClient(conn *Connection, service int) Client
}

func (s *Server) Init(laddr string) {
	s.laddr = laddr
	s.clients = make(map[int]Client)
}

// SetService registers a service with it's selector id (see protocol.InboundServiceSelect)
func (s *Server) SetService(selector int, service Service) {
	if s.services == nil {
		s.services = make(map[int]Service)
	}
	s.services[selector] = service
}

// Start creates the tcp listener and starts the connection handler in a goroutine
func (s *Server) Start() (err error) {
	logger.Info("Starting game server...")

	s.ln, err = net.Listen("tcp", s.laddr)
	if err != nil {
		return fmt.Errorf("couldn't start game server: %v", err)
	}

	s.nextIndex = id.Generator()

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
	logger.Notice("Listening on [%v]", s.laddr)

	// Accept in a seperate goroutine
	accept := make(chan net.Conn, 16)
	go func() {
		for s.t.Alive() {
			conn, err := s.ln.Accept()
			if err != nil {
				if s.t.Alive() {
					logger.Error("error accepting client: %v", err)
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

	logger.Info("Shut down")
	return nil
}

// handle is the per-connection i/o goroutine.
// buffers data into readBuffer and flushes data from writeBuffer.
// if the disconnect channel is signalled, breaks the main loop and closes the connection
func (s *Server) handle(netConn net.Conn) {
	conn := NewConnection(netConn, logger)
	client, err := s.handshake(conn)
	if err == nil && client != nil {
		s.registerClient(client)
		defer s.unregisterClient(client)

		defer safe.Recover(conn.Log())

		conn.Log().Info("accepted connection")

		go encodeFromWriteQueue(client)
		go decodeToReadQueue(client)

		// Block this thread until disconnect
		<-conn.DisconnectChan

		// ensure any pending data is flushed before disconnecting
		conn.flushWriteBuffer()
	}

	close(conn.Read)
	close(conn.Write)
	conn.conn.Close()
	conn.Log().Info("connection closed")
}

// registerClient adds a connection to the clients map
func (s *Server) registerClient(client Client) {
	s.m.Lock()
	defer s.m.Unlock()

	index := <-s.nextIndex
	s.clients[index] = client
	client.SetIndex(index)
}

// unregisterClient removes a client from the clients map
func (s *Server) unregisterClient(client Client) {
	s.m.Lock()
	defer s.m.Unlock()

	delete(s.clients, client.Index())
}

// handshake reads the service selection byte and points the connection's decode func
// towards the decode func for the selected service
func (s *Server) handshake(conn *Connection) (Client, error) {
	var serviceSelect protocol.InboundServiceSelect
	if err := serviceSelect.Decode(conn.conn, nil); err != nil {
		return nil, err
	}

	selector := int(serviceSelect.Service)

	service, ok := s.services[selector]
	if !ok {
		err := fmt.Errorf("invalid service requested: %v", serviceSelect)
		conn.Log().Error("%v", err)
		conn.Disconnect()
		return nil, err
	}

	client := service.NewClient(conn, selector)

	return client, nil
}
