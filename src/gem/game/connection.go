package game

import (
	"io"
	"net"
	"runtime"

	"gem/encoding"
	"gem/log"

	"github.com/qur/gopy/lib"
)

const (
	GameServiceType   encoding.Int8 = 14
	UpdateServiceType encoding.Int8 = 15
)

// Client is a common interface to game/update clients
type Client interface {
	Conn() *Connection
	Decode() error
	Encode(encoding.Encodable) error
	Disconnect()
	Index() int
	SetIndex(index int)
}

//go:generate gopygen -type Connection -excfield "^[a-z].*" $GOFILE
// Connection is a network-level representation of the connection.
// It handles read/write buffering, and decodes data into game packets or update requests for processing
type Connection struct {
	py.BaseObject

	Log *log.Module

	index       int
	conn        net.Conn
	readBuffer  *encoding.Buffer
	writeBuffer *encoding.Buffer
	read        chan encoding.Decodable
	write       chan encoding.Encodable
	disconnect  chan bool
}

func newConnection(conn net.Conn, parentLogger *log.Module) *Connection {
	gameConn, err := Connection{
		Log: parentLogger.SubModule(conn.RemoteAddr().String()),

		conn:        conn,
		readBuffer:  encoding.NewBuffer(),
		writeBuffer: encoding.NewBuffer(),
		read:        make(chan encoding.Decodable, 16),
		write:       make(chan encoding.Encodable, 16),
		disconnect:  make(chan bool),
	}.Alloc()
	if err != nil {
		panic(err)
	}

	return gameConn
}

// WaitForDisconnect blocks until the connection has been closed
func (conn *Connection) WaitForDisconnect() {
	<-conn.disconnect
}

// IsDisconnecting checks whether the client should be disconnecting or not
func (conn *Connection) IsDisconnecting() bool {
	select {
	case <-conn.disconnect:
		// client is disconnecting. discard
		return true
	default:
	}
	return false
}

// disconnect signals to the connection loop that this connection should be, or has been closed
func (conn *Connection) Disconnect() {
	select {
	case <-conn.disconnect:
	default:
		close(conn.disconnect)
	}
}

// Index returns the connection's unique index
func (conn *Connection) Index() int {
	return conn.index
}

// Index sets the connection's unique index
func (conn *Connection) SetIndex(index int) {
	conn.index = index
}

// WriteEncodable implements encoding.Writer
func (conn *Connection) WriteEncodable(e encoding.Encodable) {
	conn.write <- e
}

// recover captures panics in the game client handler and prints a stack trace
func (conn *Connection) recover() {
	if err := recover(); err != nil {
		stack := make([]byte, 1024*10)
		runtime.Stack(stack, true)
		conn.Log.Criticalf("Recovered from panic in game client handler: %v", err)
		conn.Log.Debug(string(stack))
	}
}

// decodeToReadQueue is the goroutine handling the read buffer
// reads from the buffer, decodes Codables using conn.decode, which can choose
// to either handle the data or place a Codable into the read queue
func decodeToReadQueue(client Client) {
	conn := client.Conn()
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
		toRead := conn.readBuffer.Len()
		for toRead > 0 && err == nil {
			err = conn.readBuffer.Try(func(b *encoding.Buffer) error {
				return client.Decode()
			})
			if err == nil {
				canTrim = true
			}
		}

		if err != nil && err != io.EOF {
			conn.Log.Criticalf("decode returned non EOF error: %v", err)
		}

		if canTrim {
			// We handled some data, discard it
			conn.readBuffer.Trim()
		}
	}
}

// encodeFromWriteQueue is the goroutine handling the write buffer
// picks from conn.write, encodes the Codables, and flushes the write buffer
func encodeFromWriteQueue(client Client) {
	conn := client.Conn()
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

			err := client.Encode(codable)
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
	_, err := conn.readBuffer.ReadFrom(conn.conn)
	return err
}
