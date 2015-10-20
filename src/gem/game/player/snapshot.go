package player

import (
	"gem/game/entity"
	"gem/game/position"
	"gem/log"
)

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

	torsoModel int
	armsModel  int
	legsModel  int
	headModel  int
	handsModel int
	feetModel  int
	beardModel int

	hairColor  int
	torsoColor int
	legsColor  int
	feetColor  int
	skinColor  int
}

func (a *AppearanceSnapshot) Gender() int {
	return a.gender
}

func (a *AppearanceSnapshot) HeadIcon() int {
	return a.headIcon
}

func (a *AppearanceSnapshot) Model(b BodyPart) int {
	switch b {
	case Torso:
		return a.torsoModel
	case Arms:
		return a.armsModel
	case Legs:
		return a.legsModel
	case Head:
		return a.headModel
	case Hands:
		return a.handsModel
	case Feet:
		return a.feetModel
	case Beard:
		return a.beardModel
	}
	panic("not reached")
}

func (a *AppearanceSnapshot) Color(b BodyPart) int {
	switch b {
	case Hair:
		return a.hairColor
	case Torso:
		return a.torsoColor
	case Legs:
		return a.legsColor
	case Feet:
		return a.feetColor
	case Skin:
		return a.skinColor
	}
	panic("not reached")
}

type AnimationsSnapshot struct {
	idle       int
	spotRotate int
	walk       int
	rotate180  int
	rotateCCW  int
	rotateCW   int
	run        int
}

func (a *AnimationsSnapshot) Animation(anim Anim) int {
	switch anim {
	case AnimIdle:
		return a.idle
	case AnimSpotRotate:
		return a.spotRotate
	case AnimWalk:
		return a.walk
	case AnimRotate180:
		return a.rotate180
	case AnimRotateCCW:
		return a.rotateCCW
	case AnimRotateCW:
		return a.rotateCW
	case AnimRun:
		return a.run
	}
	panic("not reached")
}
