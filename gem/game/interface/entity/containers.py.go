package entity

import (
	"github.com/qur/gopy/lib"

	"github.com/sinusoids/gem/pybind"
)

// These modules belong to gem.game.entity, registered in gem/game/entity/pymodule.go

var CollectionDef = pybind.Define("EntityCollection", (*Collection)(nil))
var RegisterCollection = pybind.GenerateRegisterFunc(CollectionDef)
var NewCollection = pybind.GenerateConstructor(CollectionDef).(func() *Collection)

func (c *Collection) Py_add(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(c.Add)
	return fn(args, kwds)
}

func (c *Collection) Py_remove(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(c.Remove)
	return fn(args, kwds)
}

func (c *Collection) Py_update(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(c.Update)
	return fn(args, kwds)
}

func (c *Collection) PyGet_adding() (py.Object, error) {
	fn := pybind.Wrap(c.Adding)
	return fn(nil, nil)
}

func (c *Collection) PyGet_removing() (py.Object, error) {
	fn := pybind.Wrap(c.Removing)
	return fn(nil, nil)
}

func (c *Collection) PyGet_entities() (py.Object, error) {
	fn := pybind.Wrap(c.Entities)
	return fn(nil, nil)
}

var SliceDef = pybind.Define("EntitySlice", (*Slice)(nil))
var RegisterSlice = pybind.GenerateRegisterFunc(SliceDef)
var NewSlice = pybind.GenerateConstructor(SliceDef).(func() *Slice)

func (s *Slice) Py_add(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(s.Add)
	return fn(args, kwds)
}

func (s *Slice) Py_empty(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(s.Empty)
	return fn(args, kwds)
}

func (s *Slice) Py_filter(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(s.Filter)
	return fn(args, kwds)
}

func (s *Slice) PyGet_list() (py.Object, error) {
	slice := s.Slice()
	list, err := py.NewList(int64(len(slice)))
	if err != nil {
		return list, err
	}
	for i, e := range slice {
		list.SetItem(int64(i), e.(py.Object))
	}
	return list, nil
}

func (s *Slice) PyGet_size() (py.Object, error) {
	fn := pybind.Wrap(s.Size)
	return fn(nil, nil)
}

var ListDef = pybind.Define("EntityList", (*List)(nil))
var RegisterList = pybind.GenerateRegisterFunc(ListDef)
var NewList = pybind.GenerateConstructor(ListDef).(func() *List)

func (l *List) Py_add(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(l.Add)
	return fn(args, kwds)
}

func (l *List) Py_remove(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(l.Remove)
	return fn(args, kwds)
}

func (l *List) Py_add_all(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(l.AddAll)
	return fn(args, kwds)
}

func (l *List) Py_remove_all(args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	fn := pybind.Wrap(l.RemoveAll)
	return fn(args, kwds)
}

func (l *List) PyGet_slice() (py.Object, error) {
	fn := pybind.Wrap(l.Slice)
	return fn(nil, nil)
}

func (l *List) PyGet_size() (py.Object, error) {
	fn := pybind.Wrap(l.Size)
	return fn(nil, nil)
}
