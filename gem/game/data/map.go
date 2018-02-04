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
	return nil
}
