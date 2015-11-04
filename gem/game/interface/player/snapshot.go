package player

import (
	"github.com/gtank/isaac"

	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/game/position"
	"github.com/sinusoids/gem/gem/game/server"
	"github.com/sinusoids/gem/gem/log"
)

// Snapshot clones a player to a static, read-only implementation of the player interfaces
// A snapshot is generally performed every tick, so that when the update packet is being built
// every player is syncing with the same state, and so that the state doesn't change during synchronization
func Snapshot(player Player) Player {
	srcProfile := player.Profile()
	srcWpq := player.WaypointQueue()

	snapshot := &PlayerSnapshot{
		flags:  player.Flags(),
		region: player.Region(),
		profile: &ProfileSnapshot{
			username: srcProfile.Username(),
			password: srcProfile.Password(),
			rights:   srcProfile.Rights(),
			pos:      srcProfile.Position(),
		},
	}

	currentDirection, lastDirection := srcWpq.WalkDirection()
	snapshot.waypointQueue = &WaypointQueueSnapshot{
		currentDirection: currentDirection,
		lastDirection:    lastDirection,
	}

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

	srcAnimations := player.Animations()
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
	snapshot.animations = animations

	return snapshot
}

type PlayerSnapshot struct {
	animations    Animations
	profile       Profile
	flags         entity.Flags
	region        *position.Region
	waypointQueue entity.WaypointQueue
}

func (p *PlayerSnapshot) SetNextStep(*position.Absolute) {
	panic("not implemented")
}

func (p *PlayerSnapshot) FinishInit() {
	panic("not implemented")
}

func (p *PlayerSnapshot) SetProfile(_ Profile) {
	panic("not implemented")
}

func (p *PlayerSnapshot) Profile() Profile {
	return p.profile
}

func (p *PlayerSnapshot) Log() *log.Module {
	panic("not implemented")
}

func (p *PlayerSnapshot) SetDecodeFunc(d DecodeFunc) {
	panic("not implemented")
}

func (p *PlayerSnapshot) Flags() entity.Flags {
	return p.flags
}

func (p *PlayerSnapshot) SetFlags(f entity.Flags) {
	panic("not implemented")
}

func (p *PlayerSnapshot) ClearFlags() {
	panic("not implemented")
}

func (p *PlayerSnapshot) Region() *position.Region {
	return p.region
}

func (p *PlayerSnapshot) Position() *position.Absolute {
	return p.Profile().Position()
}

func (p *PlayerSnapshot) SetPosition(*position.Absolute) {
	panic("not implemented")
}

func (p *PlayerSnapshot) Warp(*position.Absolute) {
	panic("not implemented")
}

func (p *PlayerSnapshot) Conn() *server.Connection {
	panic("not implemented")
}

func (p *PlayerSnapshot) SectorChange() {
	panic("not implemented")
}

func (p *PlayerSnapshot) RegionChange() {
	panic("not implemented")
}

func (p *PlayerSnapshot) AppearanceChange() {
	panic("not implemented")
}

func (p *PlayerSnapshot) WaypointQueue() entity.WaypointQueue {
	return p.waypointQueue
}

// EntityType identifies what kind of entity this entity is
func (p *PlayerSnapshot) EntityType() entity.EntityType {
	return entity.PlayerType
}

func (p *PlayerSnapshot) LoadProfile() {
	panic("not implemented")
}

func (p *PlayerSnapshot) ServerISAACSeed() []uint32 {
	panic("not implemented")
}

func (p *PlayerSnapshot) ISAACIn() *isaac.ISAAC {
	panic("not implemented")
}

func (p *PlayerSnapshot) ISAACOut() *isaac.ISAAC {
	panic("not implemented")
}

func (p *PlayerSnapshot) InitISAAC(inSeed, outSeed []uint32) {
	panic("not implemented")
}

func (p *PlayerSnapshot) SecureBlockSize() int {
	panic("not implemented")
}

func (p *PlayerSnapshot) SetSecureBlockSize(_ int) {
	panic("not implemented")
}

func (p *PlayerSnapshot) Animations() Animations {
	return p.animations
}

type ProfileSnapshot struct {
	username string
	password string
	rights   Rights
	pos      *position.Absolute

	skills     Skills
	appearance Appearance
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

func (p *ProfileSnapshot) SetPosition(*position.Absolute) {
	panic("not implemented")
}

func (p *ProfileSnapshot) Skills() Skills {
	return p.skills
}

func (p *ProfileSnapshot) Appearance() Appearance {
	return p.appearance
}

func (p *ProfileSnapshot) SetAppearance(Appearance) {
	panic("not implemented")
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

type WaypointQueueSnapshot struct {
	lastDirection, currentDirection int
}

func (q *WaypointQueueSnapshot) Empty() bool {
	panic("not implemented")
}

func (q *WaypointQueueSnapshot) Clear() {
	panic("not implemented")
}

func (q *WaypointQueueSnapshot) Push(point *position.Absolute) {
	panic("not implemented")
}

func (q *WaypointQueueSnapshot) Tick(mob entity.Movable) {
	panic("not implemented")
}

func (q *WaypointQueueSnapshot) WalkDirection() (current int, last int) {
	return q.currentDirection, q.lastDirection
}
