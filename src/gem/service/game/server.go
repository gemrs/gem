package game

import (
	"fmt"
	"net"
	"sync"

	"gem/log"
	"gem/runite"

	"github.com/qur/gopy/lib"
	tomb "gopkg.in/tomb.v2"
)

var logInit sync.Once
var logger *log.Module

//go:generate gopygen -type Server -exclude "^[a-z].+" $GOFILE
type Server struct {
	py.BaseObject

	laddr string
	ln    net.Listener

	runite *runite.Context

	t tomb.Tomb
}

func (s *Server) Start(laddr string, ctx *runite.Context) error {
	var err error
	s.laddr = laddr
	s.runite = ctx

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

func (s *Server) Stop() error {
	logger.Info("Stopping game server...")
	if s.t.Alive() {
		s.t.Kill(nil)
		return s.t.Wait()
	}
	return nil
}

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

	// Wait for existing connections to close
	wg.Wait()

	logger.Noticef("Shut down")
	return nil
}

func (s *Server) handle(conn net.Conn) {
	logger.Debugf("accepted connection from %v", conn)
	conn.Close()
}
