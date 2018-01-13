package entity

// Collection is an efficient, cycle based entity collection.
// The underlying collection is transactional and is updated at a fixed interval.
// The Update method should be called to commit Register/Unregistered entities.
type Collection struct {
	entities   *Set
	register   *Set // the set of entities to add this cycle
	unregister *Set // the set of entities to remove this cycle
}

func NewCollection() *Collection {
	return &Collection{
		entities:   NewSet(),
		register:   NewSet(),
		unregister: NewSet(),
	}
}

// Clone a collection
func (c *Collection) Clone() *Collection {
	return &Collection{
		entities:   c.entities.Clone(),
		register:   c.register.Clone(),
		unregister: c.unregister.Clone(),
	}
}

// Add requests a new entity be added to the collection.
// The new entity goes into the tracking list, and to the adding list
func (c *Collection) Add(entity Entity) {
	if entity == nil {
		panic("added nil entity to collection")
	}

	if c.entities.Contains(entity) {
		return
	}

	c.register.Add(entity)
	c.entities.Add(entity)
}

// Remove requests an entity be removed from the collection
// The entity is removed from the tracking list, and added to the removing list
func (c *Collection) Remove(entity Entity) {
	if entity == nil {
		panic("attempted to remove nil entity to collection")
	}

	c.unregister.Add(entity)
}

// Add all tracked entities from another collection
func (c *Collection) AddAll(other *Collection) {
	for _, e := range other.entities.Slice() {
		c.Add(e)
	}
}

// Remove all tracked entities from another collection
func (c *Collection) RemoveAll(other *Collection) {
	for _, e := range other.entities.Slice() {
		c.Remove(e)
	}
}

// Update cycles the collection. Both adding and removing lists are emptied.
func (c *Collection) Update() {
	c.register.Empty()
	c.entities.RemoveAll(c.unregister)
	c.unregister.Empty()
}

// Adding returns a slice of entities being added this cycle
func (c *Collection) Adding() *Set {
	return c.register
}

// Removing returns a slice of entities being removed this cycle
func (c *Collection) Removing() *Set {
	return c.unregister
}

// Entities returns a slice of all entities being tracked
func (c *Collection) Entities() *Set {
	return c.entities
}

// Size returns the total number of entities (adding,removing, or tracking) in the collection
func (c *Collection) Size() int {
	return c.entities.Size() + c.unregister.Size()
}

// Set is an ordered set of entities.
// it is implemented as a map[Index]Entity for efficiency; lookup can utilize
// the underlying hash-table lookup of the map type
type Set struct {
	m     map[int]Entity
	order []int
}

func NewSet() *Set {
	return &Set{
		m:     make(map[int]Entity),
		order: make([]int, 0),
	}
}

// Clone a set
func (list *Set) Clone() *Set {
	new := &Set{
		m:     make(map[int]Entity),
		order: make([]int, len(list.order)),
	}
	for k, v := range list.m {
		new.m[k] = v
	}
	copy(new.order, list.order)
	return new
}

// Empty the set
func (list *Set) Empty() {
	list.m = make(map[int]Entity)
	list.order = make([]int, 0)
}

// Slice converts the Set to a Slice
func (list *Set) Slice() []Entity {
	s := make([]Entity, len(list.order))
	for i, index := range list.order {
		s[i] = list.m[index]
	}
	return s
}

// Filter returns a new slice which contains the subset of entities with the given type
func (list *Set) Filter(typ EntityType) *Set {
	set := NewSet()
	for _, e := range list.Slice() {
		if e.EntityType() == typ {
			set.Add(e)
		}
	}
	return set
}

// Add inserts an entity into the list
func (list *Set) Add(e Entity) {
	list.m[e.Index()] = e
	list.order = append(list.order, e.Index())
}

// Remove removes an entity from the list
func (list *Set) Remove(e Entity) {
	delete(list.m, e.Index())

	pos := -1
	for i, index := range list.order {
		if index == e.Index() {
			pos = i
			break
		}
	}

	if pos == -1 {
		return
	}

	if pos == 0 {
		list.order = list.order[1:]
	} else if pos == len(list.order)-1 {
		list.order = list.order[:pos]
	} else {
		list.order = append(list.order[:pos], list.order[pos+1:]...)
	}
}

// Add inserts a list of entities into the list
func (list *Set) AddAll(slice *Set) {
	for _, e := range slice.Slice() {
		list.Add(e)
	}
}

// Remove removes a list of entities from the list
func (list *Set) RemoveAll(slice *Set) {
	for _, e := range slice.Slice() {
		list.Remove(e)
	}
}

// Returns true if the set contains e
func (list *Set) Contains(e Entity) bool {
	_, ok := list.m[e.Index()]
	return ok
}

// Returns the size of the set
func (list *Set) Size() int {
	return len(list.m)
}
