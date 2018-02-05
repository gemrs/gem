package data

import (
	"fmt"
	"math"

	astar "github.com/beefsack/go-astar"
	"github.com/gemrs/gem/gem/runite/format/rt3"
)

type CollisionFlag int

const (
	ColFloorBlockswalk                      CollisionFlag = 0x200000
	ColFloordecoBlockswalk                  CollisionFlag = 0x40000
	ColObj                                  CollisionFlag = 0x100
	ColObjBlocksfly                         CollisionFlag = 0x20000
	ColObjBlockswalkAlternative             CollisionFlag = 0x40000000
	ColWallNorth                            CollisionFlag = 0x2
	ColWallEast                             CollisionFlag = 0x8
	ColWallSouth                            CollisionFlag = 0x20
	ColWallWest                             CollisionFlag = 0x80
	ColWallNorthBlocksfly                   CollisionFlag = 0x400
	ColWallEastBlocksfly                    CollisionFlag = 0x1000
	ColWallSouthBlocksfly                   CollisionFlag = 0x4000
	ColWallWestBlocksfly                    CollisionFlag = 0x10000
	ColWallNorthBlockswalkAlternative       CollisionFlag = 0x800000
	ColWallEastBlockswalkAlternative        CollisionFlag = 0x2000000
	ColWallSouthBlockswalkAlternative       CollisionFlag = 0x8000000
	ColWallWestBlockswalkAlternative        CollisionFlag = 0x20000000
	ColCornerNorthwest                      CollisionFlag = 0x1
	ColCornerNortheast                      CollisionFlag = 0x4
	ColCornerSoutheast                      CollisionFlag = 0x10
	ColCornerSouthwest                      CollisionFlag = 0x40
	ColCornerNorthwestBlocksfly             CollisionFlag = 0x200
	ColCornerNortheastBlocksfly             CollisionFlag = 0x800
	ColCornerSoutheastBlocksfly             CollisionFlag = 0x2000
	ColCornerSouthwestBlocksfly             CollisionFlag = 0x8000
	ColCornerNorthwestBlockswalkAlternative CollisionFlag = 0x400000
	ColCornerNortheastBlockswalkAlternative CollisionFlag = 0x1000000
	ColCornerSoutheastBlockswalkAlternative CollisionFlag = 0x4000000
	ColCornerSouthwestBlockswalkAlternative CollisionFlag = 0x10000000
)

var CollisionMap = map[int]*CollisionRegion{}

type CollisionRegion struct {
	RegionX, RegionY int
	Tiles            [4][rt3.RegionSize][rt3.RegionSize]*CollisionTile
	Loaded           bool
}

func GetCollisionTile(absX, absY, absZ int) *CollisionTile {
	region := GetRegion(absX/64, absY/64)
	if region == nil {
		return nil
	}

	return region.Tiles[absZ][absX%64][absY%64]
}

func (r *CollisionRegion) RegionNorth() *CollisionRegion {
	return GetRegion(r.RegionX, r.RegionY+1)
}

func (r *CollisionRegion) RegionSouth() *CollisionRegion {
	return GetRegion(r.RegionX, r.RegionY-1)
}

func (r *CollisionRegion) RegionEast() *CollisionRegion {
	return GetRegion(r.RegionX+1, r.RegionY)
}

func (r *CollisionRegion) RegionWest() *CollisionRegion {
	return GetRegion(r.RegionX-1, r.RegionY)
}

func (r *CollisionRegion) NorthOf(t *CollisionTile) *CollisionTile {
	x, y, z := t.X, t.Y, t.Z

	y++

	if y >= rt3.RegionSize {
		y = 0
		r = r.RegionNorth()
	}

	if r == nil {
		return nil
	}

	return r.Tiles[z][x][y]
}

func (r *CollisionRegion) SouthOf(t *CollisionTile) *CollisionTile {
	x, y, z := t.X, t.Y, t.Z

	y--

	if y < 0 {
		y = rt3.RegionSize - 1
		r = r.RegionSouth()
	}

	if r == nil {
		return nil
	}

	return r.Tiles[z][x][y]
}

func (r *CollisionRegion) WestOf(t *CollisionTile) *CollisionTile {
	x, y, z := t.X, t.Y, t.Z

	x--

	if x < 0 {
		x = rt3.RegionSize - 1
		r = r.RegionWest()
	}

	if r == nil {
		return nil
	}

	return r.Tiles[z][x][y]
}

func (r *CollisionRegion) EastOf(t *CollisionTile) *CollisionTile {
	x, y, z := t.X, t.Y, t.Z

	x++
	if x >= rt3.RegionSize {
		x = 0
		r = r.RegionEast()
	}

	if r == nil {
		return nil
	}

	return r.Tiles[z][x][y]
}

type CollisionTile struct {
	X, Y, Z    int
	AbsX, AbsY int
	Flag       CollisionFlag
	Region     *CollisionRegion
}

func (t *CollisionTile) String() string {
	return fmt.Sprintf("CollisionTile(%v, %v, %v, %v)", t.AbsX, t.AbsY, t.Z, t.Flag)
}

func (t *CollisionTile) North() *CollisionTile {
	return t.Region.NorthOf(t)
}

func (t *CollisionTile) South() *CollisionTile {
	return t.Region.SouthOf(t)
}

func (t *CollisionTile) East() *CollisionTile {
	return t.Region.EastOf(t)
}

func (t *CollisionTile) West() *CollisionTile {
	return t.Region.WestOf(t)
}

func (t *CollisionTile) NorthEast() *CollisionTile {
	north := t.North()
	if north == nil {
		return nil
	}
	return north.East()
}

func (t *CollisionTile) NorthWest() *CollisionTile {
	north := t.North()
	if north == nil {
		return nil
	}
	return north.West()
}

func (t *CollisionTile) SouthEast() *CollisionTile {
	south := t.South()
	if south == nil {
		return nil
	}
	return south.East()
}

func (t *CollisionTile) SouthWest() *CollisionTile {
	south := t.South()
	if south == nil {
		return nil
	}
	return south.West()
}

func (t *CollisionTile) Blocked() bool {
	return t.Flag&(ColObj|ColFloorBlockswalk) != 0
}

func (t *CollisionTile) PathNeighbors() []astar.Pather {
	validTiles := []astar.Pather{}

	// I'm sure there's an easier way to do this than a huge chain of logic, but
	// this works fine for now.

	north := t.North()
	northEast := t.NorthEast()
	east := t.East()
	southEast := t.SouthEast()
	south := t.South()
	southWest := t.SouthWest()
	west := t.West()
	northWest := t.NorthWest()

	if north != nil && t.Flag&ColWallNorth == 0 && !north.Blocked() {
		validTiles = append(validTiles, north)
	}

	if northEast != nil &&
		t.Flag&ColCornerNortheast == 0 && !northEast.Blocked() &&
		t.Flag&ColWallNorth == 0 && northEast.Flag&ColWallSouth == 0 &&
		t.Flag&ColWallEast == 0 && northEast.Flag&ColWallWest == 0 {
		valid := true

		if north == nil || north.Blocked() {
			valid = false
		}

		if east == nil || east.Blocked() {
			valid = false
		}

		if valid {
			validTiles = append(validTiles, northEast)
		}
	}

	if east != nil &&
		t.Flag&ColWallEast == 0 && !east.Blocked() {
		validTiles = append(validTiles, east)
	}

	if southEast != nil &&
		t.Flag&ColCornerSoutheast == 0 && !southEast.Blocked() &&
		t.Flag&ColWallSouth == 0 && southEast.Flag&ColWallNorth == 0 &&
		t.Flag&ColWallEast == 0 && southEast.Flag&ColWallWest == 0 {
		valid := true

		if south == nil || south.Blocked() {
			valid = false
		}

		if east == nil || east.Blocked() {
			valid = false
		}

		if valid {
			validTiles = append(validTiles, southEast)
		}
	}

	if south != nil &&
		t.Flag&ColWallSouth == 0 && !south.Blocked() {
		validTiles = append(validTiles, south)
	}

	if southWest != nil &&
		t.Flag&ColCornerSouthwest == 0 && !southWest.Blocked() &&
		t.Flag&ColWallSouth == 0 && southWest.Flag&ColWallNorth == 0 &&
		t.Flag&ColWallWest == 0 && southWest.Flag&ColWallEast == 0 {
		valid := true

		if south == nil || south.Blocked() {
			valid = false
		}

		if west == nil || west.Blocked() {
			valid = false
		}

		if valid {
			validTiles = append(validTiles, southWest)
		}
	}

	if west != nil &&
		t.Flag&ColWallWest == 0 && !west.Blocked() {
		validTiles = append(validTiles, west)
	}

	if northWest != nil &&
		t.Flag&ColCornerNorthwest == 0 && !northWest.Blocked() &&
		t.Flag&ColWallNorth == 0 && northWest.Flag&ColWallSouth == 0 &&
		t.Flag&ColWallWest == 0 && northWest.Flag&ColWallEast == 0 {
		valid := true

		if north == nil || north.Blocked() {
			valid = false
		}

		if west == nil || west.Blocked() {
			valid = false
		}

		if valid {
			validTiles = append(validTiles, northWest)
		}
	}

	return validTiles
}

func (t *CollisionTile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(*CollisionTile)
	dx := float64(t.AbsX - toT.AbsX)
	dy := float64(t.AbsY - toT.AbsY)
	cost := math.Ceil(math.Sqrt(dx*dx + dy*dy))
	return cost
}

func (t *CollisionTile) PathEstimatedCost(to astar.Pather) float64 {
	return t.PathNeighborCost(to)
}
