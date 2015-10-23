package server

import (
	"io"
	"net"

	"gem/encoding"
	"gem/log"
	"gem/util/safe"

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

//go:generate gopygen -type Connection -excfield ".*" $GOFILE
// Connection is a network-level representation of the connection.
// It handles read/write buffering, and decodes data into game packets or update requests for processing
type Connection struct {
	py.BaseObject

	ReadBuffer     *encoding.Buffer
	WriteBuffer    *encoding.Buffer
	Read           chan encoding.Decodable
	Write          chan encoding.Encodable
	DisconnectChan chan bool

	log   log.Logger
	index int
	conn  net.Conn
}

func (c *Connection) Init(conn net.Conn, parentLogger *log.Module) error {
	c.ReadBuffer = encoding.NewBuffer()
	c.WriteBuffer = encoding.NewBuffer()
	c.Read = make(chan encoding.Decodable, 16)
	c.Write = make(chan encoding.Encodable, 16)
	c.DisconnectChan = make(chan bool)

	c.log = parentLogger.SubModule(conn.RemoteAddr().String())
	c.conn = conn
	return nil
}

func (c *Connection) Log() *log.Module {
	return c.log.(*log.Module)
}

// WaitForDisconnect blocks until the connection has been closed
func (conn *Connection) WaitForDisconnect() {
	<-conn.DisconnectChan
}

// IsDisconnecting checks whether the client should be disconnecting or not
func (conn *Connection) IsDisconnecting() bool {
	select {
	case <-conn.DisconnectChan:
		// client is disconnecting. discard
		return true
	default:
	}
	return false
}

// disconnect signals to the connection loop that this connection should be, or has been closed
func (conn *Connection) Disconnect() {
	select {
	case <-conn.DisconnectChan:
	default:
		close(conn.DisconnectChan)
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

// decodeToReadQueue is the goroutine handling the read buffer
// reads from the buffer, decodes Codables using conn.decode, which can choose
// to either handle the data or place a Codable into the read queue
func decodeToReadQueue(client Client) {
	conn := client.Conn()
	defer safe.Recover(conn.Log())
	for {
		err := conn.fillReadBuffer()
		if err != nil {
			conn.Log().Debugf("read error: %v", err)
			conn.Disconnect()
			break
		}

		// at this point, the only error should be because we didn't have enough data
		// todo: formalize this and check for the right error
		canTrim := false
		toRead := conn.ReadBuffer.Len()
		for toRead > 0 && err == nil {
			err = conn.ReadBuffer.Try(func(b *encoding.Buffer) error {
				return client.Decode()
			})
			if err == nil {
				canTrim = true
			}
		}

		if err != nil && err != io.EOF {
			conn.Log().Criticalf("decode returned non EOF error: %v", err)
		}

		if canTrim {
			// We handled some data, discard it
			conn.ReadBuffer.Trim()
		}
	}
}

// encodeFromWriteQueue is the goroutine handling the write buffer
// picks from conn.write, encodes the Codables, and flushes the write buffer
func encodeFromWriteQueue(client Client) {
	conn := client.Conn()
	defer safe.Recover(conn.Log())
L:
	for {
		select {
		case codable, ok := <-conn.Write:
			if ok {
				select {
				case <-conn.DisconnectChan:
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
				conn.Log().Debugf("write error: %v", err)
				conn.Disconnect()
				break L
			}
		}
	}
}

// flushWriteBuffer drains the write buffer and ensures that all data is written to
// the connection. If conn.Write returns an error (timeout), the client is disconnected.
func (conn *Connection) flushWriteBuffer() error {
	for conn.WriteBuffer.Len() > 0 {
		_, err := conn.WriteBuffer.WriteTo(conn.conn)
		if err != nil {
			return err
		}
	}
	conn.WriteBuffer.Trim()
	return nil
}

// fillReadBuffer pulls data from the connection and buffers it for decoding into the readBuffer
// launched in a goroutine by Server.handle
func (conn *Connection) fillReadBuffer() error {
	_, err := conn.ReadBuffer.ReadFrom(conn.conn)
	return err
}
