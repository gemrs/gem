package server

import (
	"fmt"
	"io"
	"net"
	"runtime"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/util/safe"
	"github.com/gemrs/willow/log"
)

// Connection is a network-level representation of the connection.
// It handles read/write buffering, and decodes data into game packets or update requests for processing
type Connection struct {
	ReadBuffer     *encoding.Buffer
	WriteBuffer    *encoding.Buffer
	Read           chan encoding.Decodable
	Write          chan encoding.Encodable
	DisconnectChan chan bool

	log   log.Log
	index int
	conn  net.Conn
}

func NewConnection(conn net.Conn, parentLogger log.Log) *Connection {
	return &Connection{
		ReadBuffer:     encoding.NewBuffer(),
		WriteBuffer:    encoding.NewBuffer(),
		Read:           make(chan encoding.Decodable, 16),
		Write:          make(chan encoding.Encodable, 16),
		DisconnectChan: make(chan bool),
		log:            parentLogger.Child("connection", log.MapContext{"addr": conn.RemoteAddr().String()}),
		conn:           conn,
	}
}

func (c *Connection) NetConn() net.Conn {
	return c.conn
}

func (c *Connection) Expired() chan bool {
	return c.DisconnectChan
}

func (c *Connection) Expire() {
	close(c.DisconnectChan)
}

func (c *Connection) Log() log.Log {
	return c.log.(log.Log)
}

// WaitForDisconnect blocks until the connection has been closed
func (conn *Connection) WaitForDisconnect() {
	<-conn.Expired()
}

// IsDisconnecting checks whether the client should be disconnecting or not
func (conn *Connection) IsDisconnecting() bool {
	select {
	case <-conn.Expired():
		// client is disconnecting. discard
		return true
	default:
	}
	return false
}

// disconnect signals to the connection loop that this connection should be, or has been closed
func (conn *Connection) Disconnect() {
	select {
	case <-conn.Expired():
	default:
		conn.Expire()
	}
}

// decodeToReadQueue is the goroutine handling the read buffer
// reads from the buffer, decodes Codables using conn.decode, which can choose
// to either handle the data or place a Codable into the read queue
func decodeToReadQueue(client GameClient) {
	conn := client.Conn()
	defer safe.Recover(conn.Log())
	for {
		err := conn.fillReadBuffer()
		if err != nil {
			conn.Log().Debug("read error: %v", err)
			conn.Disconnect()
			break
		}

		// at this point, the only error should be because we didn't have enough data
		// todo: formalize this and check for the right error
		canTrim := false
		toRead := conn.ReadBuffer.Len()
		stack := make([]byte, 1024*10)
		for toRead > 0 && err == nil {
			err = conn.ReadBuffer.Try(func(b *encoding.Buffer) (err error) {
				defer func() {
					if e := recover(); e != nil {
						if e, ok := e.(error); ok {
							err = e
						} else {
							err = fmt.Errorf("%v", e)
						}
						runtime.Stack(stack, true)
					}
				}()
				client.Decode()
				return nil
			})

			if err == nil {
				canTrim = true
			}
		}

		if err != nil && err != io.EOF {
			conn.Log().Error("decode returned non EOF error: %v", err)
			conn.Log().Error(string(stack))
		}

		if canTrim {
			// We handled some data, discard it
			conn.ReadBuffer.Trim()
		}
	}
}

// encodeFromWriteQueue is the goroutine handling the write buffer
// picks from conn.write, encodes the Codables, and flushes the write buffer
func encodeFromWriteQueue(client GameClient) {
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
				conn.Log().Debug("write error: %v", err)
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
