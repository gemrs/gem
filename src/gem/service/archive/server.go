package archive

import (
	"fmt"
	"net"
	"regexp"
	"sync"

	"gem/log"
	"gem/runite"
	"gem/runite/format/rt3"

	"bufio"
	"github.com/qur/gopy/lib"
	tomb "gopkg.in/tomb.v2"
)

var logInit sync.Once
var logger *log.Module
var requestRegexp = regexp.MustCompile("JAGGRAB /([a-z]+)[0-9\\-]+")

//go:generate gopygen -type Server -excfunc "^[a-z].+" $GOFILE
type Server struct {
	py.BaseObject

	laddr string
	ln    net.Listener

	archives *rt3.ArchiveFS

	t tomb.Tomb
}

func (s *Server) Start(laddr string, ctx *runite.Context) error {
	var err error
	s.laddr = laddr
	index, err := ctx.FS.Index(0)
	if err != nil {
		return err
	}
	s.archives = rt3.NewArchiveFS(index)

	logInit.Do(func() {
		logger = log.New("archive")
	})

	logger.Info("Starting archive server...")
	s.ln, err = net.Listen("tcp", s.laddr)
	if err != nil {
		return fmt.Errorf("couldn't start archive server: %v", err)
	}

	logger.Infof("Found %v archives", s.archives.FileCount())

	s.t.Go(s.run)
	return nil
}

func (s *Server) Stop() error {
	logger.Info("Stopping archive server...")
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
	io := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

Outer:
	for {
		var request string
		var err error
		if request, err = io.ReadString('\n'); err != nil {
			break
		}

		matches := requestRegexp.FindStringSubmatch(request)
		if matches == nil {
			logger.Errorf("invalid request: %v", request)
			break
		}

		if b, err := io.ReadByte(); err != nil || b != '\n' {
			logger.Errorf("missing newline in request: %v", request)
			break
		}

		archive, err := s.archives.ResolveArchive(matches[1])
		if err != nil {
			logger.Errorf("couldn't locate archive %v: %v", matches[1], err)
			break
		}

		toWrite := len(archive)
		for toWrite > 0 {
			n, err := io.Write(archive[len(archive)-toWrite:])
			if err != nil {
				break Outer
			}
			toWrite -= n
			io.Flush()
		}
	}
	conn.Close()
}
