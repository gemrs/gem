package position

const (
	AreaSize   int = 8
	RegionSize int = 13
)

// Positionable is an object which has an absolute position in the world
type Positionable interface {
	Position() *Absolute
	SetPosition(*Absolute)
}
