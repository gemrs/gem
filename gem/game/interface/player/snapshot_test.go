package player

import (
	"testing"

	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/position"
)

func TestSnapshot(t *testing.T) {
	dummyPos := position.NewAbsolute(3500, 3500, 1)

	dummyPlayer := &PlayerSnapshot{
		flags:        entity.MobFlagWalkUpdate,
		loadedRegion: dummyPos.RegionOf(),
		profile: &ProfileSnapshot{
			username: "a player",
			password: "hunter2",
			rights:   RightsPlayer,
			pos:      dummyPos,

			skills: &SkillsSnapshot{
				combatLevel: 8,
			},
			appearance: &AppearanceSnapshot{
				gender:   1,
				headIcon: 1,
				models:   map[BodyPart]int{},
				colors:   map[BodyPart]int{},
			},
		},
		animations: &AnimationsSnapshot{
			anims: map[Anim]int{},
		},
		waypointQueue: &WaypointQueueSnapshot{
			currentDirection: 1,
			lastDirection:    2,
		},
	}

	snapshot := Snapshot(dummyPlayer)

	if !comparePlayers(dummyPlayer, snapshot) {
		t.Error("snapshotted player didn't match")
	}
}
