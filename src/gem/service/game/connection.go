package game

import (
	"io"
	"net"
	"runtime"
	"time"

	"gem/encoding"
	"gem/log"
	"gem/protocol"
	"gem/service/game/player"

	"github.com/qur/gopy/lib"
)

const (
	GameService   encoding.Int8 = 14
	UpdateService encoding.Int8 = 15
)

// encodeDecodeFunc is the function currently used for parsing the read stream and
// dealing with the incoming data.
// If an error is returned, it is assumed that we didn't have enough data, and
// the underlying buffer's read pointer is not altered.
type encodeDecodeFunc func(*context, *encoding.Buffer) error

//go:generate gopygen -type Connection -excfield "^[a-z].*" $GOFILE
// Connection is a network-level representation of the connection.
// It handles read/write buffering, and decodes data into game packets or update requests for processing
type Connection struct {
	py.BaseObject

	Index   Index
	Log     *log.Module
	Session *player.Session
	Profile *player.Profile

	conn        net.Conn
	readBuffer  *encoding.Buffer
	writeBuffer *encoding.Buffer
	read        chan encoding.Codable
	write       chan encoding.Codable
	disconnect  chan bool
	decode      encodeDecodeFunc
}

func newConnection(conn net.Conn, parentLogger *log.Module) *Connection {
	session, err := player.Session{}.Alloc()
	if err != nil {
		panic(err)
	}

	// FIXME: There's something nasty going on here.. Possibly a data race
	gameConn, err := Connection{
		Log:     parentLogger.SubModule(conn.RemoteAddr().String()),
		Session: session,

		conn:        conn,
		readBuffer:  encoding.NewBuffer(),
		writeBuffer: encoding.NewBuffer(),
		read:        make(chan encoding.Codable, 16),
		write:       make(chan encoding.Codable, 16),
		disconnect:  make(chan bool),
	}.Alloc()
	if err != nil {
		panic(err)
	}

	return gameConn
}

// disconnect signals to the connection loop that this connection should be, or has been closed
func (conn *Connection) Disconnect() {
	select {
	case <-conn.disconnect:
	default:
		close(conn.disconnect)
	}
}

func (conn *Connection) recover() {
	if err := recover(); err != nil {
		stack := make([]byte, 1024*10)
		runtime.Stack(stack, true)
		conn.Log.Criticalf("Recovered from panic in game client handler: %v", err)
		conn.Log.Debug(string(stack))
	}
}

// handshake reads the service selection byte and points the connection's decode func
// towards the decode func for the selected service
func (conn *Connection) handshake(ctx *context, b *encoding.Buffer) error {
	var svc protocol.ServiceSelect
	if err := svc.Decode(b, nil); err != nil {
		return err
	}

	switch svc.Service {
	case UpdateService:
		conn.write <- new(protocol.UpdateHandshakeResponse)

		conn.Log.Infof("new update client")
		conn.decode = ctx.update.decodeRequest
		return nil
	case GameService:
		conn.Log.Infof("new game client")
		conn.decode = ctx.game.handshake
		return nil
	default:
		conn.Log.Errorf("invalid service requested: %v", svc)
		conn.Disconnect()
	}

	return nil
}

// Write is a convenience wrapper around writeBuffer.Write(p)
func (conn *Connection) Write(p []byte) (n int, err error) {
	return conn.writeBuffer.Write(p)
}

// decodeToReadQueue is the goroutine handling the read buffer
// reads from the buffer, decodes Codables using conn.decode, which can choose
// to either handle the data or place a Codable into the read queue
func (conn *Connection) decodeToReadQueue(connCtx *context) {
	defer conn.recover()
	for {
		err := conn.fillReadBuffer()
		if err != nil {
			conn.Log.Debugf("read error: %v", err)
			conn.Disconnect()
			break
		}

		// at this point, the only error should be because we didn't have enough data
		// todo: formalize this and check for the right error
		canTrim := false
		for conn.readBuffer.Len() > 0 && err == nil {
			err = conn.readBuffer.Try(func(b *encoding.Buffer) error {
				return conn.decode(connCtx, b)
			})
			if err == nil {
				canTrim = true
			}
		}

		if err != nil && err != io.EOF {
			conn.Log.Criticalf("decode returned non EOF error")
		}

		if canTrim {
			// We handled some data, discard it
			conn.readBuffer.Trim()
		}
	}
}

// encodeFromWriteQueue is the goroutine handling the write buffer
// picks from conn.write, encodes the Codables, and flushes the write buffer
func (conn *Connection) encodeFromWriteQueue(connCtx *context) {
	defer conn.recover()
L:
	for {
		select {
		case codable, ok := <-conn.write:
			if ok {
				select {
				case <-conn.disconnect:
					ok = false
				default:
				}
			}

			if !ok {
				break L
			}

			err := codable.Encode(conn.writeBuffer, nil)
			if err == nil {
				err = conn.flushWriteBuffer()
			}

			if err != nil {
				conn.Log.Debugf("write error: %v", err)
				conn.Disconnect()
				break L
			}
		}
	}
}

// flushWriteBuffer drains the write buffer and ensures that all data is written to
// the connection. If conn.Write returns an error (timeout), the client is disconnected.
func (conn *Connection) flushWriteBuffer() error {
	for conn.writeBuffer.Len() > 0 {
		_, err := conn.writeBuffer.WriteTo(conn.conn)
		if err != nil {
			return err
		}
	}
	conn.writeBuffer.Trim()
	return nil
}

// fillReadBuffer pulls data from the connection and buffers it for decoding into the readBuffer
// launched in a goroutine by Server.handle
func (conn *Connection) fillReadBuffer() error {
	conn.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, err := conn.readBuffer.ReadFrom(conn.conn)
	return err
}
