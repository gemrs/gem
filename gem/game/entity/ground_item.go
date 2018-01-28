package entity

import "github.com/gemrs/gem/gem/game/item"

type GroundItem interface {
	Entity
	Expire()
	Expired() chan bool
	Definition() *item.Definition
	Item() *item.Stack
}
