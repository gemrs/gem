// Generated by gopygen; DO NOT EDIT
package game

import (
	"fmt"
	"github.com/sinusoids/gem/gem/encoding"
	"github.com/sinusoids/gem/gem/game/player"
	"github.com/sinusoids/gem/gem/game/position"
	"github.com/sinusoids/gem/gem/game/server"

	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

var PlayerDef = py.Class{
	Name:    "Player",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*Player)(nil),
}

// Registers this type with a python module
func RegisterPlayer(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = PlayerDef.Create(); err != nil {
		return err
	}

	if err = module.AddObject("Player", class); err != nil {
		return err
	}

	return nil
}

// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func NewPlayer(arg_0 *server.Connection, arg_1 *GameService) (*Player, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := PlayerDef.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*Player)
	err = alloc.Init(arg_0, arg_1)
	return alloc, err
}

func (obj *Player) PyInit(_args *py.Tuple, kwds *py.Dict) error {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 2 {
		return fmt.Errorf("(Player) PyInit: parameter length mismatch")
	}

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "*server.Connection")
	if err != nil {
		return err
	}

	args[1].Incref()
	in_1, err := gopygen.TypeConvIn(args[1], "*GameService")
	if err != nil {
		return err
	}

	err = obj.Init(in_0.(*server.Connection), in_1.(*GameService))

	args[0].Decref()

	args[1].Decref()

	return err
}

func (client *Player) Py_Session(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Session: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := client.Session()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "player.Session")
	if err != nil {
		return nil, err
	}
	out_0.Incref()

	return out_0, nil

}

func (client *Player) Py_Profile(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Profile: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := client.Profile()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "player.Profile")
	if err != nil {
		return nil, err
	}
	out_0.Incref()

	return out_0, nil

}

func (client *Player) Py_Conn(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Conn: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := client.Conn()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "*server.Connection")
	if err != nil {
		return nil, err
	}
	out_0.Incref()

	return out_0, nil

}

func (client *Player) Py_Decode(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Decode: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := client.Decode()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "error")
	if err != nil {
		return nil, err
	}
	out_0.Incref()

	return out_0, nil

}

func (client *Player) Py_Position(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Position: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := client.Position()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "*position.Absolute")
	if err != nil {
		return nil, err
	}
	out_0.Incref()

	return out_0, nil

}

func (client *Player) Py_WalkDirection(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_WalkDirection: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0, res1 := client.WalkDirection()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "int")
	if err != nil {
		return nil, err
	}
	out_0.Incref()

	out_1, err := gopygen.TypeConvOut(res1, "int")
	if err != nil {
		return nil, err
	}
	out_1.Incref()

	return py.PackTuple(out_0, out_1)

}

func (client *Player) Py_Flags(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Flags: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := client.Flags()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "entity.Flags")
	if err != nil {
		return nil, err
	}
	out_0.Incref()

	return out_0, nil

}

func (client *Player) Py_Region(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_Region: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	res0 := client.Region()

	// Remove local references

	out_0, err := gopygen.TypeConvOut(res0, "*position.Region")
	if err != nil {
		return nil, err
	}
	out_0.Incref()

	return out_0, nil

}

func (client *Player) Py_SetPosition(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_SetPosition: parameter length mismatch")
	}
	// Convert parameters

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "*position.Absolute")
	if err != nil {
		return nil, err
	}

	// Make the function call

	client.SetPosition(in_0.(*position.Absolute))

	// Remove local references

	args[0].Decref()

	py.None.Incref()
	return py.None, nil

}

func (client *Player) Py_SetAppearance(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_SetAppearance: parameter length mismatch")
	}
	// Convert parameters

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "player.Appearance")
	if err != nil {
		return nil, err
	}

	// Make the function call

	client.SetAppearance(in_0.(player.Appearance))

	// Remove local references

	args[0].Decref()

	py.None.Incref()
	return py.None, nil

}

func (client *Player) Py_AppearanceUpdated(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 0 {
		return nil, fmt.Errorf("Py_AppearanceUpdated: parameter length mismatch")
	}
	// Convert parameters

	// Make the function call

	client.AppearanceUpdated()

	// Remove local references

	py.None.Incref()
	return py.None, nil

}

func (client *Player) Py_Encode(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_Encode: parameter length mismatch")
	}
	// Convert parameters

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "encoding.Encodable")
	if err != nil {
		return nil, err
	}

	// Make the function call

	res0 := client.Encode(in_0.(encoding.Encodable))

	// Remove local references

	args[0].Decref()

	out_0, err := gopygen.TypeConvOut(res0, "error")
	if err != nil {
		return nil, err
	}
	out_0.Incref()

	return out_0, nil

}

func (client *Player) Py_SendMessage(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != 1 {
		return nil, fmt.Errorf("Py_SendMessage: parameter length mismatch")
	}
	// Convert parameters

	args[0].Incref()
	in_0, err := gopygen.TypeConvIn(args[0], "string")
	if err != nil {
		return nil, err
	}

	// Make the function call

	client.SendMessage(in_0.(string))

	// Remove local references

	args[0].Decref()

	py.None.Incref()
	return py.None, nil

}
