package archive

import (
	"sync"
	"fmt"
	"net"

	"gem/log"

	tomb "gopkg.in/tomb.v2"
)

var logInit sync.Once
var logger *log.Module

type ArchiveServer struct {
	laddr string

	t tomb.Tomb
}

func NewServer(laddr string) *ArchiveServer{
	logInit.Do(func() {
		logger = log.New("archive")
	})

	return &ArchiveServer{
		laddr: laddr,
	}
}

func (s *ArchiveServer) Start() error {
	logger.Info("Starting archive server...")
	s.t.Go(s.run)
	return nil
}

func (s *ArchiveServer) Stop() error {
	logger.Info("Stopping archive server...")
	s.t.Kill(nil)
	return s.t.Wait()
}

func (s *ArchiveServer) run() error {
	ln, err := net.Listen("tcp", s.laddr)
	if err != nil {
		err = fmt.Errorf("couldn't start archive server: %v", err)
		logger.Critical(err.Error())
		return err
	}

	logger.Noticef("Started on %v", s.laddr)

	// Accept in a seperate goroutine
	accept := make(chan net.Conn, 16)
	go func() {
		for s.t.Alive() {
			conn, err := ln.Accept()
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
	ln.Close()

	// Wait for existing connections to close
	wg.Wait()

	logger.Noticef("Shut down")
	return nil
}

func (s *ArchiveServer) handle(conn net.Conn) {

	conn.Close()
}
