package player

import (
	"github.com/gtank/isaac"
	"github.com/qur/gopy/lib"
)

//go:generate gopygen -type Session $GOFILE
type Session struct {
	py.BaseObject

	RandIn  isaac.ISAAC
	RandOut isaac.ISAAC
	RandKey []int32
}
