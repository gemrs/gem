package player

import (
	"github.com/sinusoids/gem/gem/game/entity"
	"github.com/sinusoids/gem/gem/game/position"
	"github.com/sinusoids/gem/gem/log"
)

// Snapshot clones a player to a static, read-only implementation of the player interfaces
// A snapshot is generally performed every tick, so that when the update packet is being built
// every player is syncing with the same state, and so that the state doesn't change during synchronization
func Snapshot(player Player) Player {
	srcProfile := player.Profile()
	srcSession := player.Session()
	snapshot := &PlayerSnapshot{
		profile: &ProfileSnapshot{
			username: srcProfile.Username(),
			password: srcProfile.Password(),
			rights:   srcProfile.Rights(),
			pos:      srcProfile.Position(),
		},
	}

	session := &SessionSnapshot{
		flags:  srcSession.Flags(),
		region: srcSession.Region(),
	}
	session.currentWalkDir, session.lastWalkDir = srcSession.WalkDirection()
	snapshot.session = session

	skills := &SkillsSnapshot{
		combatLevel: srcProfile.Skills().CombatLevel(),
	}
	snapshot.profile.(*ProfileSnapshot).skills = skills

	srcAppearance := srcProfile.Appearance()
	appearance := &AppearanceSnapshot{
		gender:   srcAppearance.Gender(),
		headIcon: srcAppearance.HeadIcon(),
		models: map[BodyPart]int{
			Torso: srcAppearance.Model(Torso),
			Arms:  srcAppearance.Model(Arms),
			Legs:  srcAppearance.Model(Legs),
			Head:  srcAppearance.Model(Head),
			Hands: srcAppearance.Model(Hands),
			Feet:  srcAppearance.Model(Feet),
			Beard: srcAppearance.Model(Beard),
		},
		colors: map[BodyPart]int{
			Torso: srcAppearance.Color(Torso),
			Hair:  srcAppearance.Color(Hair),
			Legs:  srcAppearance.Color(Legs),
			Feet:  srcAppearance.Color(Feet),
			Skin:  srcAppearance.Color(Skin),
		},
	}
	snapshot.profile.(*ProfileSnapshot).appearance = appearance

	srcAnimations := srcProfile.Animations()
	animations := &AnimationsSnapshot{
		anims: map[Anim]int{
			AnimIdle:       srcAnimations.Animation(AnimIdle),
			AnimSpotRotate: srcAnimations.Animation(AnimSpotRotate),
			AnimWalk:       srcAnimations.Animation(AnimWalk),
			AnimRotate180:  srcAnimations.Animation(AnimRotate180),
			AnimRotateCCW:  srcAnimations.Animation(AnimRotateCCW),
			AnimRotateCW:   srcAnimations.Animation(AnimRotateCW),
			AnimRun:        srcAnimations.Animation(AnimRun),
		},
	}
	snapshot.profile.(*ProfileSnapshot).animations = animations

	return snapshot
}

type PlayerSnapshot struct {
	profile Profile
	session Session
}

func (p *PlayerSnapshot) Profile() Profile {
	return p.profile
}

func (p *PlayerSnapshot) Session() Session {
	return p.session
}

func (p *PlayerSnapshot) Log() *log.Module {
	panic("not implemented")
}

func (p *PlayerSnapshot) Flags() entity.Flags {
	return p.Session().Flags()
}

func (p *PlayerSnapshot) WalkDirection() (current int, last int) {
	return p.Session().WalkDirection()
}

func (p *PlayerSnapshot) Region() *position.Region {
	return p.Session().Region()
}

func (p *PlayerSnapshot) Position() *position.Absolute {
	return p.Profile().Position()
}

func (p *PlayerSnapshot) SetPosition(*position.Absolute) {
	panic("not implemented")
}

type SessionSnapshot struct {
	region         *position.Region
	flags          entity.Flags
	currentWalkDir int
	lastWalkDir    int
}

func (s *SessionSnapshot) Flags() entity.Flags {
	return s.flags
}

func (s *SessionSnapshot) WalkDirection() (current int, last int) {
	return s.currentWalkDir, s.lastWalkDir
}

func (s *SessionSnapshot) Region() *position.Region {
	return s.region
}

type ProfileSnapshot struct {
	username string
	password string
	rights   Rights
	pos      *position.Absolute

	skills     Skills
	appearance Appearance
	animations Animations
}

func (p *ProfileSnapshot) Username() string {
	return p.username
}

func (p *ProfileSnapshot) Password() string {
	return p.password
}

func (p *ProfileSnapshot) Rights() Rights {
	return p.rights
}

func (p *ProfileSnapshot) Position() *position.Absolute {
	return p.pos
}

func (p *ProfileSnapshot) Skills() Skills {
	return p.skills
}

func (p *ProfileSnapshot) Appearance() Appearance {
	return p.appearance
}

func (p *ProfileSnapshot) Animations() Animations {
	return p.animations
}

type SkillsSnapshot struct {
	combatLevel int
}

func (s *SkillsSnapshot) CombatLevel() int {
	return s.combatLevel
}

type AppearanceSnapshot struct {
	gender   int
	headIcon int

	models map[BodyPart]int
	colors map[BodyPart]int
}

func (a *AppearanceSnapshot) Gender() int {
	return a.gender
}

func (a *AppearanceSnapshot) HeadIcon() int {
	return a.headIcon
}

func (a *AppearanceSnapshot) Model(b BodyPart) int {
	return a.models[b]
}

func (a *AppearanceSnapshot) Color(b BodyPart) int {
	return a.colors[b]
}

type AnimationsSnapshot struct {
	anims map[Anim]int
}

func (a *AnimationsSnapshot) Animation(anim Anim) int {
	return a.anims[anim]
}
