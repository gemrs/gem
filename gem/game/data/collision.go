package data

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"os"

	astar "github.com/beefsack/go-astar"
	"github.com/gemrs/gem/gem/game/position"
)

type CollisionFlag int

const (
	ColFloorBlockswalk                      CollisionFlag = 0x200000
	ColFloordecoBlockswalk                                = 0x40000
	ColObj                                                = 0x100
	ColObjBlocksfly                                       = 0x20000
	ColObjBlockswalkAlternative                           = 0x40000000
	ColWallNorth                                          = 0x2
	ColWallEast                                           = 0x8
	ColWallSouth                                          = 0x20
	ColWallWest                                           = 0x80
	ColWallNorthBlocksfly                                 = 0x400
	ColWallEastBlocksfly                                  = 0x1000
	ColWallSouthBlocksfly                                 = 0x4000
	ColWallWestBlocksfly                                  = 0x10000
	ColWallNorthBlockswalkAlternative                     = 0x800000
	ColWallEastBlockswalkAlternative                      = 0x2000000
	ColWallSouthBlockswalkAlternative                     = 0x8000000
	ColWallWestBlockswalkAlternative                      = 0x20000000
	ColCornerNorthwest                                    = 0x1
	ColCornerNortheast                                    = 0x4
	ColCornerSoutheast                                    = 0x10
	ColCornerSouthwest                                    = 0x40
	ColCornerNorthwestBlocksfly                           = 0x200
	ColCornerNortheastBlocksfly                           = 0x800
	ColCornerSoutheastBlocksfly                           = 0x2000
	ColCornerSouthwestBlocksfly                           = 0x8000
	ColCornerNorthwestBlockswalkAlternative               = 0x400000
	ColCornerNortheastBlockswalkAlternative               = 0x1000000
	ColCornerSoutheastBlockswalkAlternative               = 0x4000000
	ColCornerSouthwestBlockswalkAlternative               = 0x10000000
)

const regionSize = position.RegionSize * position.SectorSize

type CollisionData [regionSize][regionSize]CollisionTile

type CollisionTile struct {
	X, Y   int
	Flag   CollisionFlag
	Region *CollisionData
}

func (t CollisionTile) North() *CollisionTile {
	if t.Y+1 >= regionSize {
		return nil
	}

	return &t.Region[t.X][t.Y+1]
}

func (t CollisionTile) South() *CollisionTile {
	if t.Y-1 < 0 {
		return nil
	}

	return &t.Region[t.X][t.Y-1]
}

func (t CollisionTile) East() *CollisionTile {
	if t.X+1 >= regionSize {
		return nil
	}

	return &t.Region[t.X+1][t.Y]
}

func (t CollisionTile) West() *CollisionTile {
	if t.X-1 < 0 {
		return nil
	}

	return &t.Region[t.X-1][t.Y]
}

func (t CollisionTile) NorthEast() *CollisionTile {
	north := t.North()
	if north == nil {
		return nil
	}
	return north.East()
}

func (t CollisionTile) NorthWest() *CollisionTile {
	north := t.North()
	if north == nil {
		return nil
	}
	return north.West()
}

func (t CollisionTile) SouthEast() *CollisionTile {
	south := t.South()
	if south == nil {
		return nil
	}
	return south.East()
}

func (t CollisionTile) SouthWest() *CollisionTile {
	south := t.South()
	if south == nil {
		return nil
	}
	return south.West()
}

func (t CollisionTile) PathNeighbors() []astar.Pather {
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

	if north != nil && t.Flag&ColWallNorth == 0 && north.Flag&ColObj == 0 {
		validTiles = append(validTiles, *north)
	}

	if northEast != nil &&
		t.Flag&ColCornerNortheast == 0 && northEast.Flag&ColObj == 0 &&
		t.Flag&ColWallNorth == 0 && northEast.Flag&ColWallSouth == 0 &&
		t.Flag&ColWallEast == 0 && northEast.Flag&ColWallWest == 0 {
		valid := true

		if north != nil && north.Flag&ColObj != 0 {
			valid = false
		}

		if east != nil && east.Flag&ColObj != 0 {
			valid = false
		}

		if valid {
			validTiles = append(validTiles, *northEast)
		}
	}

	if east != nil &&
		t.Flag&ColWallEast == 0 && east.Flag&ColObj == 0 {
		validTiles = append(validTiles, *east)
	}

	if southEast != nil &&
		t.Flag&ColCornerSoutheast == 0 && southEast.Flag&ColObj == 0 &&
		t.Flag&ColWallSouth == 0 && southEast.Flag&ColWallNorth == 0 &&
		t.Flag&ColWallEast == 0 && southEast.Flag&ColWallWest == 0 {
		valid := true

		if south != nil && south.Flag&ColObj != 0 {
			valid = false
		}

		if east != nil && east.Flag&ColObj != 0 {
			valid = false
		}

		if valid {
			validTiles = append(validTiles, *southEast)
		}
	}

	if south != nil &&
		t.Flag&ColWallSouth == 0 && south.Flag&ColObj == 0 {
		validTiles = append(validTiles, *south)
	}

	if southWest != nil &&
		t.Flag&ColCornerSouthwest == 0 && southWest.Flag&ColObj == 0 &&
		t.Flag&ColWallSouth == 0 && southWest.Flag&ColWallNorth == 0 &&
		t.Flag&ColWallWest == 0 && southWest.Flag&ColWallEast == 0 {
		valid := true

		if south != nil && south.Flag&ColObj != 0 {
			valid = false
		}

		if west != nil && west.Flag&ColObj != 0 {
			valid = false
		}

		if valid {
			validTiles = append(validTiles, *southWest)
		}
	}

	if west != nil &&
		t.Flag&ColWallWest == 0 && west.Flag&ColObj == 0 {
		validTiles = append(validTiles, *west)
	}

	if northWest != nil &&
		t.Flag&ColCornerNorthwest == 0 && northWest.Flag&ColObj == 0 &&
		t.Flag&ColWallNorth == 0 && northWest.Flag&ColWallSouth == 0 &&
		t.Flag&ColWallWest == 0 && northWest.Flag&ColWallEast == 0 {
		valid := true

		if north != nil && north.Flag&ColObj != 0 {
			valid = false
		}

		if west != nil && west.Flag&ColObj != 0 {
			valid = false
		}

		if valid {
			validTiles = append(validTiles, *northWest)
		}
	}

	return validTiles
}

func (t CollisionTile) PathNeighborCost(to astar.Pather) float64 {
	toT := to.(CollisionTile)
	dx := float64(t.X - toT.X)
	dy := float64(t.Y - toT.Y)
	return math.Ceil(math.Sqrt(dx*dx + dy*dy))
}

func (t CollisionTile) PathEstimatedCost(to astar.Pather) float64 {
	return t.PathNeighborCost(to)
}

var TestCollision CollisionData

func init() {
	var data [regionSize][regionSize]int
	fd, err := os.Open("data/collision.json")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	jsonData, err := ioutil.ReadAll(fd)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(jsonData, &data)

	for x, s := range data {
		for y, flag := range s {
			TestCollision[x][y].X = x
			TestCollision[x][y].Y = y
			TestCollision[x][y].Flag = CollisionFlag(flag)
			TestCollision[x][y].Region = &TestCollision
		}
	}
}
