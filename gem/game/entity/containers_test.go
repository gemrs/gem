package entity_test

import (
	"testing"

	"github.com/gemrs/gem/gem/game/entity"
	"github.com/gemrs/gem/gem/game/position"
)

type MockEntity struct {
	index int
	typ   entity.EntityType
}

func (e MockEntity) Region() *position.Region       { return nil }
func (e MockEntity) RegionChange()                  {}
func (e MockEntity) SectorChange()                  {}
func (e MockEntity) Position() *position.Absolute   { return nil }
func (e MockEntity) SetPosition(*position.Absolute) {}

func (e MockEntity) EntityType() entity.EntityType {
	return e.typ
}

func (e MockEntity) Index() int {
	return e.index
}

func compareLists(t *testing.T, name string, slice *entity.Set, tclist []MockEntity) {
	if slice.Size() != len(tclist) {
		t.Errorf("Invalid size in %v list: got %v expected %v", name, slice.Size(), len(tclist))
	}

	for _, a := range slice.Slice() {
		found := false
		for _, b := range tclist {
			if a.Index() == b.Index() {
				found = true
			}
		}
		if found == false {
			t.Errorf("Couldn't find entity %v in %v list", name, a.Index())
		}
	}
}

// Simulates a few entity add/remove/update cycles
func TestContainer(t *testing.T) {
	col := entity.NewCollection()

	e0, e1, e2 := MockEntity{0, entity.EntityType("a")}, MockEntity{1, entity.EntityType("a")}, MockEntity{2, entity.EntityType("b")}

	tcs := []struct {
		add      []MockEntity
		entities []MockEntity
		remove   []MockEntity
	}{
		{
			add:      []MockEntity{e0},
			entities: []MockEntity{e0},
			remove:   []MockEntity{},
		},
		{
			add:      []MockEntity{e1},
			entities: []MockEntity{e0, e1},
			remove:   []MockEntity{},
		},
		{
			add:      []MockEntity{},
			entities: []MockEntity{e0, e1},
			remove:   []MockEntity{},
		},
		{
			add:      []MockEntity{e2},
			entities: []MockEntity{e0, e1, e2},
			remove:   []MockEntity{e0},
		},
		{
			add:      []MockEntity{},
			entities: []MockEntity{e1, e2},
			remove:   []MockEntity{e1},
		},
		{
			add:      []MockEntity{},
			entities: []MockEntity{e2},
			remove:   []MockEntity{e2},
		},
		{
			add:      []MockEntity{},
			entities: []MockEntity{},
			remove:   []MockEntity{},
		},
	}

	for i, tc := range tcs {
		t.Logf("cycle %v", i)
		col.Update()

		for _, e := range tc.add {
			col.Add(e)
		}

		for _, e := range tc.remove {
			col.Remove(e)
		}

		compareLists(t, "adding", col.Adding(), tc.add)
		compareLists(t, "removing", col.Removing(), tc.remove)
		compareLists(t, "entities", col.Entities(), tc.entities)
	}
}

// Test slice filtering works correctly
func TestFilter(t *testing.T) {
	sl := entity.NewSet()

	e0, e1, e2 := MockEntity{0, entity.EntityType("a")}, MockEntity{1, entity.EntityType("a")}, MockEntity{2, entity.EntityType("b")}

	sl.Add(e0)
	sl.Add(e1)
	sl.Add(e2)

	var subset *entity.Set
	subset = sl.Filter("a")
	compareLists(t, "filtered", subset, []MockEntity{e0, e1})
	subset = sl.Filter("b")
	compareLists(t, "filtered", subset, []MockEntity{e2})
}
