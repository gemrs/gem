package impl

import (
	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/protocol"
)

//glua:bind
type GroundItem struct {
	index      int
	position   *position.Absolute
	item       *item.Stack
	sector     protocol.Sector
	expireChan chan bool
}

//glua:bind constructor GroundItem
func NewGroundItem(item *item.Stack, pos *position.Absolute, world protocol.World) *GroundItem {
	entity := &GroundItem{
		index:      entity.NextIndex(),
		position:   pos,
		item:       item,
		sector:     world.Sector(pos.Sector()),
		expireChan: make(chan bool),
	}
	entity.activate()
	return entity
}

func (item *GroundItem) activate() {
	item.sector.Add(item)
}

//glua:Bind
func (item *GroundItem) Expire() {
	close(item.expireChan)
	item.sector.Remove(item)
}

func (item *GroundItem) Expired() chan bool {
	return item.expireChan
}

//glua:bind
func (item *GroundItem) Index() int {
	return item.index
}

//glua:bind
func (item *GroundItem) Item() *item.Stack {
	return item.item
}

//glua:bind
func (item *GroundItem) Definition() *item.Definition {
	return item.item.Definition()
}

// Position returns the absolute position of the item
//glua:bind
func (item *GroundItem) Position() *position.Absolute {
	return item.position
}

// SetPosition warps the item to a given location
func (item *GroundItem) SetPosition(pos *position.Absolute) {
	panic("ground items cannot be repositioned")
}

// EntityType identifies what kind of entity this entity is
func (item *GroundItem) EntityType() entity.EntityType {
	return entity.GroundItemType
}
