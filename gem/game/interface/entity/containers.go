package entity

import (
	"github.com/qur/gopy/lib"
)

// Collection is an efficient, cycle based entity collection.
// The underlying collection is transactional and is updated at a fixed interval.
// The Update method should be called to commit Register/Unregistered entities.
type Collection struct {
	py.BaseObject

	entities   *List
	register   *Slice // the set of entities to add this cycle
	unregister *Slice // the set of entities to remove this cycle
}

func (c *Collection) Init() {
	c.entities = NewList()
	c.register = NewSlice()
	c.unregister = NewSlice()
}

// Add requests a new entity be added to the collection.
// The new entity goes into the tracking list, and to the adding list
func (c *Collection) Add(entity Entity) {
	c.register.Add(entity)
	c.entities.Add(entity)
}

// Remove requests an entity be removed from the collection
// The entity is removed from the tracking list, and added to the removing list
func (c *Collection) Remove(entity Entity) {
	c.unregister.Add(entity)
	c.entities.Remove(entity)
}

// Update cycles the collection. Both adding and removing lists are emptied.
func (c *Collection) Update() {
	c.register.Empty()
	c.unregister.Empty()
}

// Adding returns a slice of entities being added this cycle
func (c *Collection) Adding() *Slice {
	return c.register
}

// Removing returns a slice of entities being removed this cycle
func (c *Collection) Removing() *Slice {
	return c.unregister
}

// Entities returns a slice of all entities being tracked
func (c *Collection) Entities() *Slice {
	return c.entities.Slice()
}

// Slice is a slice of entities.
// Slices can be added to and emptied, but not removed from. They are intended for buffering
// entities for addition to a List.
type Slice struct {
	py.BaseObject

	s []Entity
}

func (s *Slice) Init() {
	s.Empty()
}

func (s *Slice) Empty() {
	s.s = make([]Entity, 0)
}

// Filter returns a new slice which contains the subset of entities with the given type
func (s *Slice) Filter(typ EntityType) *Slice {
	slice := NewSlice()
	for _, e := range s.Slice() {
		if e.EntityType() == typ {
			slice.Add(e)
		}
	}
	return slice
}

func (s *Slice) Add(e Entity) {
	s.s = append(s.s, e)
}

func (s *Slice) Slice() []Entity {
	return s.s
}

func (s *Slice) Size() int {
	return len(s.s)
}

// List is implemented as a map[Index]Entity for efficiency; lookup can utilize
// the underlying hash-table lookup of the map type
type List struct {
	py.BaseObject

	m map[int]Entity
}

func (l *List) Init() {
	l.m = make(map[int]Entity)
}

// Slice converts the List to a Slice
func (list *List) Slice() *Slice {
	slice := NewSlice()
	for _, e := range list.m {
		slice.Add(e)
	}
	return slice
}

// Add inserts an entity into the list
func (list *List) Add(e Entity) {
	list.m[e.Index()] = e
}

// Remove removes an entity from the list
func (list *List) Remove(e Entity) {
	delete(list.m, e.Index())
}

// Add inserts a list of entities into the list
func (list *List) AddAll(slice *Slice) {
	for _, e := range slice.Slice() {
		list.Add(e)
	}
}

// Remove removes a list of entities from the list
func (list *List) RemoveAll(slice *Slice) {
	for _, e := range slice.Slice() {
		list.Remove(e)
	}
}

func (list *List) Size() int {
	return len(list.m)
}
