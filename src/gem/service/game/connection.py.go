// Generated by gopygen; DO NOT EDIT
package game

import (
	"fmt"
	"gem/encoding"
	"gem/log"
	"gem/service/game/player"

	"github.com/qur/gopy/lib"

	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

var ConnectionDef = py.Class{
	Name:    "Connection",
	Pointer: (*Connection)(nil),
}

// Registers this type with a python module
func RegisterConnection(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = ConnectionDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("Connection", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func (obj Connection) Alloc() (*Connection, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := ConnectionDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*Connection)
	// Copy fields

	alloc.Index = obj.Index

	alloc.Log = obj.Log

	alloc.Session = obj.Session

	alloc.Profile = obj.Profile

	alloc.conn = obj.conn

	alloc.readBuffer = obj.readBuffer

	alloc.writeBuffer = obj.writeBuffer

	alloc.read = obj.read

	alloc.write = obj.write

	alloc.disconnect = obj.disconnect

	alloc.decode = obj.decode

	return alloc, nil
}

func (obj *Connection) PyGet_Index() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Index, "Index")
}

func (obj *Connection) PySet_Index(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "Index")
	if err != nil {
		return err
	}
	obj.Index = val.(Index)
	return nil
}

func (obj *Connection) PyGet_Log() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Log, "*log.Module")
}

func (obj *Connection) PySet_Log(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "*log.Module")
	if err != nil {
		return err
	}
	obj.Log = val.(*log.Module)
	return nil
}

func (obj *Connection) PyGet_Session() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Session, "*player.Session")
}

func (obj *Connection) PySet_Session(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "*player.Session")
	if err != nil {
		return err
	}
	obj.Session = val.(*player.Session)
	return nil
}

func (obj *Connection) PyGet_Profile() (py.Object, error) {
	return gopygen.TypeConvOut(obj.Profile, "*player.Profile")
}

func (obj *Connection) PySet_Profile(arg py.Object) error {
	val, err := gopygen.TypeConvIn(arg, "*player.Profile")
	if err != nil {
		return err
	}
	obj.Profile = val.(*player.Profile)
	return nil
}

func (conn *Connection) Py_Disconnect(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Disconnect: parameter length mismatch")
	}

	conn.Disconnect()

	py.None.Incref()
	return py.None, nil

}

func (conn *Connection) Py_recover(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_recover: parameter length mismatch")
	}

	conn.recover()

	py.None.Incref()
	return py.None, nil

}

func (conn *Connection) Py_handshake(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 2 {
		return nil, fmt.Errorf("Py_handshake: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "*context")
	if err != nil {
		return nil, err
	}

	in_1, err := gopygen.TypeConvIn(args[1], "*encoding.Buffer")
	if err != nil {
		return nil, err
	}

	res0 := conn.handshake(in_0.(*context), in_1.(*encoding.Buffer))

	out_0, err := gopygen.TypeConvOut(res0, "error")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (conn *Connection) Py_Write(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Write: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "[]byte")
	if err != nil {
		return nil, err
	}

	res0, res1 := conn.Write(in_0.([]byte))

	out_0, err := gopygen.TypeConvOut(res0, "int")
	if err != nil {
		return nil, err
	}

	out_1, err := gopygen.TypeConvOut(res1, "error")
	if err != nil {
		return nil, err
	}

	return py.PackTuple(out_0, out_1)

}

func (conn *Connection) Py_decodeToReadQueue(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_decodeToReadQueue: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "*context")
	if err != nil {
		return nil, err
	}

	conn.decodeToReadQueue(in_0.(*context))

	py.None.Incref()
	return py.None, nil

}

func (conn *Connection) Py_encodeFromWriteQueue(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_encodeFromWriteQueue: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "*context")
	if err != nil {
		return nil, err
	}

	conn.encodeFromWriteQueue(in_0.(*context))

	py.None.Incref()
	return py.None, nil

}

func (conn *Connection) Py_flushWriteBuffer(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_flushWriteBuffer: parameter length mismatch")
	}

	res0 := conn.flushWriteBuffer()

	out_0, err := gopygen.TypeConvOut(res0, "error")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (conn *Connection) Py_fillReadBuffer(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_fillReadBuffer: parameter length mismatch")
	}

	res0 := conn.fillReadBuffer()

	out_0, err := gopygen.TypeConvOut(res0, "error")
	if err != nil {
		return nil, err
	}

	return out_0, nil

}

func (conn *Connection) Py_SendMessage(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_SendMessage: parameter length mismatch")
	}

	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	conn.SendMessage(in_0.(string))

	py.None.Incref()
	return py.None, nil

}