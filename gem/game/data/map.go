package data

import (
	"github.com/gemrs/gem/gem/runite"
	"github.com/gemrs/gem/gem/runite/format/rt3"
)

var Map rt3.Map

//glua:bind
func LoadMap(runite *runite.Context) error {
	err := Map.Load(runite.FS, GetMapKeys)
	if err != nil {
		return err
	}

	logger.Notice("Loaded [%v] map regions", Map.NumRegions())

	return nil
}

func GetRegion(x, y int) *CollisionRegion {
	id := x<<8 | y
	r, ok := CollisionMap[id]
	if ok && r.Loaded {
		return r
	}

	mapRegion, err := Map.RegionById(id)
	if err != nil {
		return nil
	}

	// Ensure that we've got the all regions surrounding this initialized
	// Then load the target region with collision data

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			region, err := Map.Region(i, j)
			if err == nil {
				initRegion(region)
			}
		}
	}

	// We have to set the Loaded flag first, to avoid recursion
	CollisionMap[id].Loaded = true

	loadCollisionFromRegion(mapRegion)
	return CollisionMap[id]
}

func initRegion(region *rt3.Region) {
	if _, ok := CollisionMap[region.Region]; ok {
		return
	}

	colRegion := &CollisionRegion{
		RegionX: region.X,
		RegionY: region.Y,
	}

	for z := 0; z < 4; z++ {
		for x := 0; x < rt3.RegionSize; x++ {
			for y := 0; y < rt3.RegionSize; y++ {
				tile := &CollisionTile{
					AbsX:   region.AbsX + x,
					AbsY:   region.AbsY + y,
					X:      x,
					Y:      y,
					Z:      z,
					Region: colRegion,
				}

				colRegion.Tiles[z][x][y] = tile
			}
		}
	}

	CollisionMap[region.Region] = colRegion
}

func loadCollisionFromRegion(region *rt3.Region) {
	loadCollisionFromLandscape(region)
	loadCollisionFromLocations(region)
}

func loadCollisionFromLandscape(region *rt3.Region) {
	colRegion := CollisionMap[region.Region]

	landscape := region.Landscape.Tiles
	for z := 0; z < 4; z++ {
		for x := 0; x < rt3.RegionSize; x++ {
			for y := 0; y < rt3.RegionSize; y++ {
				if (landscape[z][x][y].RenderType & 1) == 1 {
					height := z
					if (landscape[1][x][y].RenderType & 2) == 2 {
						height--
					}

					if height >= 0 && height <= 3 {
						colRegion.Tiles[height][x][y].Flag |= ColFloorBlockswalk
					}
				}
			}
		}
	}
}

func loadCollisionFromLocations(region *rt3.Region) {
	locations := region.Locations.Locations
	landscape := region.Landscape.Tiles

	for _, obj := range locations {
		height := obj.LocalZ
		if (landscape[1][obj.LocalX][obj.LocalY].RenderType & 2) == 2 {
			height--
		}

		if height >= 0 && height <= 3 {
			addCollisionObject(region, &obj, height)
		}
	}

}

func addCollisionObject(region *rt3.Region, obj *rt3.MapLocation, height int) {
	colRegion := CollisionMap[region.Region]
	absX, absY := region.AbsX+obj.LocalX, region.AbsY+obj.LocalY
	id := obj.Id

	if id == -1 {
		return
	}

	definition, err := Config.Object(id)
	if err != nil {
		// Some maps just use objects that aren't in the cache.. shrug
		return
	}

	xLength := 0
	yLength := 0
	if obj.Orientation == 1 || obj.Orientation == 3 {
		xLength = definition.SizeX
		yLength = definition.SizeY
	} else {
		xLength = definition.SizeY
		yLength = definition.SizeX
	}

	if obj.Type == 22 {
		if definition.ClipType == 1 {
			colRegion.Tiles[height][obj.LocalX][obj.LocalY].Flag |= ColFloordecoBlockswalk
		}
	} else {
		if obj.Type != 10 && obj.Type != 11 {
			if obj.Type >= 12 {
				if definition.ClipType != 0 {
					addCollisionForSolidObject(definition, absX, absY, obj.LocalZ, xLength, yLength, definition.BlocksProjectile)
				}
			} else if obj.Type == 0 {
				if definition.ClipType != 0 {
					addCollisionForVariableObject(absX, absY, obj.LocalZ, obj.Type, obj.Orientation, definition.BlocksProjectile)
				}
			} else if obj.Type == 1 {
				if definition.ClipType != 0 {
					addCollisionForVariableObject(absX, absY, obj.LocalZ, obj.Type, obj.Orientation, definition.BlocksProjectile)
				}
			} else {
				if obj.Type == 2 {
					if definition.ClipType != 0 {
						addCollisionForVariableObject(absX, absY, obj.LocalZ, obj.Type, obj.Orientation, definition.BlocksProjectile)
					}
				} else if obj.Type == 3 {
					if definition.ClipType != 0 {
						addCollisionForVariableObject(absX, absY, obj.LocalZ, obj.Type, obj.Orientation, definition.BlocksProjectile)
					}
				} else if obj.Type == 9 {
					if definition.ClipType != 0 {
						addCollisionForSolidObject(definition, absX, absY, obj.LocalZ, xLength, yLength, definition.BlocksProjectile)
					}
				} else if obj.Type == 4 {
					// addBoundaryDecoration?
				} else if obj.Type == 6 {
					// addBoundaryDecoration?
				} else if obj.Type == 4 {
					// addBoundaryDecoration?
				} else if obj.Type == 7 {
					// addBoundaryDecoration?
				} else if obj.Type == 8 {
					// addBoundaryDecoration?
				}
			}

		} else {
			if definition.ClipType != 0 {
				addCollisionForSolidObject(definition, absX, absY, obj.LocalZ, xLength, yLength, definition.BlocksProjectile)
			}
		}
	}
}

func setCollisionFlags(x, y, z int, flags CollisionFlag) {
	tile := GetCollisionTile(x, y, z)
	if tile == nil {
		return
	}
	tile.Flag |= flags
}

func addCollisionForSolidObject(def *rt3.ObjectDefinition, x, y, z, xLength, yLength int, flag1 bool) {
	flags := ColObj

	if flag1 {
		flags |= ColObjBlocksfly
	}

	for i := x; i < x+xLength; i++ {
		for j := y; j < y+yLength; j++ {
			setCollisionFlags(i, j, z, flags)
		}
	}
}

func addCollisionForVariableObject(x, y, z, typ, orientation int, flag bool) {
	if typ == 0 {
		if orientation == 0 {
			setCollisionFlags(x, y, z, 0x80)
			setCollisionFlags(x-1, y, z, 0x8)
		}
		if orientation == 1 {
			setCollisionFlags(x, y, z, 0x2)
			setCollisionFlags(x, y+1, z, 0x20)
		}
		if orientation == 2 {
			setCollisionFlags(x, y, z, 8)
			setCollisionFlags(x+1, y, z, 0x80)
		}
		if orientation == 3 {
			setCollisionFlags(x, y, z, 0x20)
			setCollisionFlags(x, y-1, z, 0x2)
		}
	}

	if typ == 1 || typ == 3 {
		if orientation == 0 {
			setCollisionFlags(x, y, z, 0x1)
			setCollisionFlags(x-1, y+1, z, 0x10)
		}
		if orientation == 1 {
			setCollisionFlags(x, y, z, 0x4)
			setCollisionFlags(x+1, y+1, z, 0x40)
		}
		if orientation == 2 {
			setCollisionFlags(x, y, z, 16)
			setCollisionFlags(x+1, y-1, z, 0x1)
		}
		if orientation == 3 {
			setCollisionFlags(x, y, z, 64)
			setCollisionFlags(x-1, y-1, z, 0x4)
		}
	}

	if typ == 2 {
		if orientation == 0 {
			setCollisionFlags(x, y, z, 0x82)
			setCollisionFlags(x-1, y, z, 0x8)
			setCollisionFlags(x, y+1, z, 0x20)
		}
		if orientation == 1 {
			setCollisionFlags(x, y, z, 0xa)
			setCollisionFlags(x, y+1, z, 0x20)
			setCollisionFlags(x+1, y, z, 0x80)
		}
		if orientation == 2 {
			setCollisionFlags(x, y, z, 0x28)
			setCollisionFlags(x+1, y, z, 0x80)
			setCollisionFlags(x, y-1, z, 0x2)
		}
		if orientation == 3 {
			setCollisionFlags(x, y, z, 0xa0)
			setCollisionFlags(x, y-1, z, 0x2)
			setCollisionFlags(x-1, y, z, 0x8)
		}
	}

	if flag {
		if typ == 0 {
			if orientation == 0 {
				setCollisionFlags(x, y, z, 0x10000)
				setCollisionFlags(x-1, y, z, 0x1000)
			}
			if orientation == 1 {
				setCollisionFlags(x, y, z, 0x400)
				setCollisionFlags(x, y+1, z, 0x4000)
			}
			if orientation == 2 {
				setCollisionFlags(x, y, z, 0x1000)
				setCollisionFlags(x+1, y, z, 0x10000)
			}
			if orientation == 3 {
				setCollisionFlags(x, y, z, 0x4000)
				setCollisionFlags(x, y-1, z, 0x400)
			}
		}

		if typ == 1 || typ == 3 {
			if orientation == 0 {
				setCollisionFlags(x, y, z, 0x200)
				setCollisionFlags(x-1, y+1, z, 0x2000)
			}
			if orientation == 1 {
				setCollisionFlags(x, y, z, 0x800)
				setCollisionFlags(x+1, y+1, z, 0x8000)
			}
			if orientation == 2 {
				setCollisionFlags(x, y, z, 0x2000)
				setCollisionFlags(x+1, y-1, z, 0x200)
			}
			if orientation == 3 {
				setCollisionFlags(x, y, z, 0x8000)
				setCollisionFlags(x-1, y-1, z, 0x800)
			}
		}

		if typ == 2 {
			if orientation == 0 {
				setCollisionFlags(x, y, z, 0x10400)
				setCollisionFlags(x-1, y, z, 0x1000)
				setCollisionFlags(x, y+1, z, 0x4000)
			}
			if orientation == 1 {
				setCollisionFlags(x, y, z, 0x1400)
				setCollisionFlags(x, y+1, z, 0x4000)
				setCollisionFlags(x+1, y, z, 0x10000)
			}
			if orientation == 2 {
				setCollisionFlags(x, y, z, 0x5000)
				setCollisionFlags(x+1, y, z, 0x10000)
				setCollisionFlags(x, y-1, z, 0x400)
			}
			if orientation == 3 {
				setCollisionFlags(x, y, z, 0x14000)
				setCollisionFlags(x, y-1, z, 0x400)
				setCollisionFlags(x-1, y, z, 0x1000)
			}
		}
	}
}
