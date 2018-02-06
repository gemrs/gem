package rt3

import (
	"errors"

	"github.com/gemrs/gem/gem/core/encoding"
)

var (
	ErrNoObject = errors.New("no object with given id")
	ErrNoItem   = errors.New("no item with given id")
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
	objects        []*ObjectDefinition
	objectsArchive *Archive
	items          []*ItemDefinition
	itemsArchive   *Archive
}

func (c *Config) Load(fs *JagFS) error {
	index, err := fs.Index(IdxConfigs)
	if err != nil {
		return err
	}

	c.objectsArchive, err = index.Archive(CfgObject)
	if err != nil {
		return err
	}

	c.objects = make([]*ObjectDefinition, len(c.objectsArchive.Entries))

	c.itemsArchive, err = index.Archive(CfgItem)
	if err != nil {
		return err
	}

	c.items = make([]*ItemDefinition, len(c.itemsArchive.Entries))

	return nil
}

func (c *Config) ObjectsCount() int {
	return len(c.objects)
}

func (c *Config) Object(id int) (*ObjectDefinition, error) {
	if id >= len(c.objects) {
		return nil, ErrNoObject
	}

	if obj := c.objects[id]; obj != nil {
		return obj, nil
	}

	c.objects[id] = NewObjectDefinition(id)
	buf := encoding.NewBufferBytes(c.objectsArchive.Entries[id])
	c.objects[id].Decode(buf, nil)
	return c.objects[id], nil
}

func (c *Config) ItemsCount() int {
	return len(c.items)
}

func (c *Config) Item(id int) (*ItemDefinition, error) {
	if id >= len(c.items) {
		return nil, ErrNoItem
	}

	if obj := c.items[id]; obj != nil {
		return obj, nil
	}

	c.items[id] = NewItemDefinition(id)
	buf := encoding.NewBufferBytes(c.itemsArchive.Entries[id])
	c.items[id].Decode(buf, nil)
	return c.items[id], nil
}
