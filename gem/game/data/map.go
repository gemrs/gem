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

	logger.Notice("Loaded [%v] map regions", len(Map.Regions))
	buildCollisionMap(Map)
	logger.Notice("Constructed collision data")

	return nil
}

func buildCollisionMap(m rt3.Map) {
	for _, region := range m.Regions {
		loadCollisionFromRegion(region)
	}
}

func loadCollisionFromRegion(region *rt3.Region) {
	colRegion := &CollisionRegion{
		RegionX: region.X,
		RegionY: region.Y,
	}

	landscape := region.Landscape.Tiles
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
	CollisionMap[region.Region] = colRegion
}
