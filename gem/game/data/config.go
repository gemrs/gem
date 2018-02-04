package data

import (
	"github.com/gemrs/gem/gem/runite"
	"github.com/gemrs/gem/gem/runite/format/rt3"
)

var Config *rt3.Config

//glua:bind
func LoadConfig(runite *runite.Context) error {
	Config = rt3.NewConfig()
	err := Config.Load(runite.FS)
	if err != nil {
		return err
	}

	logger.Notice("Loaded [%v] object definitions", len(Config.Objects))
	return nil
}
