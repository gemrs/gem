package game

import (
	"net"
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

//go:generate gopygen -type GameConnection -excfield "^[a-z].*" $GOFILE
// GameConnection is a network-level representation of the connection.
// It handles read/write buffering, and decodes data into game packets or update requests for processing
type GameConnection struct {
	py.BaseObject

	Index   Index
	Log     *log.Module
	Session *player.Session
	Profile *player.Profile

	conn        net.Conn
	read        chan encoding.Codable
	write       chan encoding.Codable
	readBuffer  *encoding.Buffer
	writeBuffer *encoding.Buffer
	decode      encodeDecodeFunc
	disconnect  chan int
	canDecode   chan int
	active      bool
}

func newConnection(index Index, conn net.Conn, parentLogger *log.Module) *GameConnection {
	session, err := player.Session{}.Alloc()
	if err != nil {
		panic(err)
	}

	// FIXME: There's something nasty going on here.. Possibly a data race
	gameConn, err := GameConnection{
		Log:     parentLogger.SubModule(conn.RemoteAddr().String()),
		Index:   index,
		Session: session,

		conn:        conn,
		read:        make(chan encoding.Codable, 16),
		write:       make(chan encoding.Codable, 16),
		readBuffer:  encoding.NewBuffer(),
		writeBuffer: encoding.NewBuffer(),
		disconnect:  make(chan int, 2),
		canDecode:   make(chan int, 2),
		active:      true,
	}.Alloc()
	if err != nil {
		panic(err)
	}

	return gameConn
}

// disconnect signals to the connection loop that this connection should be, or has been closed
func (conn *GameConnection) Disconnect() {
	conn.active = false
	conn.disconnect <- 1
}

// handshake reads the service selection byte and points the connection's decode func
// towards the decode func for the selected service
func (conn *GameConnection) handshake(ctx *context, b *encoding.Buffer) error {
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
func (conn *GameConnection) Write(p []byte) (n int, err error) {
	return conn.writeBuffer.Write(p)
}

// flushWriteBuffer drains the write buffer and ensures that all data is written to
// the connection. If conn.Write returns an error (timeout), the client is disconnected.
func (conn *GameConnection) flushWriteBuffer() {
	for conn.writeBuffer.Len() > 0 {
		_, err := conn.writeBuffer.WriteTo(conn.conn)
		if err != nil {
			conn.Log.Debug("write error")
			conn.Disconnect()
			break
		}
	}
	conn.writeBuffer.Trim()
}

// fillReadBuffer pulls data from the connection and buffers it for decoding into the readBuffer
// launched in a goroutine by Server.handle
func (conn *GameConnection) fillReadBuffer() {
	for {
		conn.conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		_, err := conn.readBuffer.ReadFrom(conn.conn)
		if err != nil {
			conn.Log.Debugf("read error: %v", err)
			conn.Disconnect()
			break
		}

		conn.canDecode <- 1
	}
}
