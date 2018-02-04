package rt3

import (
	"github.com/gemrs/gem/gem/core/encoding"
)

const (
	CfgArea            = 35
	CfgEnum            = 8
	CfgHitbar          = 33
	CfgHitmark         = 32
	CfgIdentkit        = 3
	CfgItem            = 10
	CfgInv             = 5
	CfgNpc             = 9
	CfgObject          = 6
	CfgOverlay         = 4
	CfgParams          = 11
	CfgSequence        = 12
	CfgSpotanim        = 13
	CfgStruct          = 34
	CfgUnderlay        = 1
	CfgVarbit          = 14
	CfgVarclient       = 19
	CfgVarclientstring = 15
	CfgVarplayer       = 16
)

type Config struct {
	Objects []*ObjectDefinition
}

func (c *Config) Load(fs *JagFS) error {
	index, err := fs.Index(IdxConfigs)
	if err != nil {
		return err
	}

	objectsArchive, err := index.Archive(CfgObject)
	if err != nil {
		return err
	}

	c.Objects = make([]*ObjectDefinition, len(objectsArchive.Entries))
	for id, data := range objectsArchive.Entries {
		c.Objects[id] = NewObjectDefinition(id)
		buf := encoding.NewBufferBytes(data)
		c.Objects[id].Decode(buf, nil)
	}

	return nil
}
