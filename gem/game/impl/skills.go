package impl

import (
	"math"

	"github.com/gemrs/gem/gem/protocol"
)

var experienceTable [99]int

func init() {
	accum := 0
	for level := 0; level < 99; level++ {
		realLevel := level + 1
		exp := int(float64(realLevel+300) * math.Pow(2.0, float64(realLevel)/7.0))
		accum += exp
		experienceTable[level] = accum / 4
	}
}

func skillExpToLevel(exp int) int {
	for i := 0; i < 98; i++ {
		if exp < experienceTable[i] {
			return i + 1
		}
	}
	return 99
}

//glua:bind
type Skills struct {
	player protocol.Player

	combatLevel int
	skills      [21]*Skill
}

//glua:bind constructor Skills
func NewSkills() *Skills {
	skills := &Skills{}
	for i := range skills.skills {
		skills.skills[i] = NewSkill(protocol.SkillId(i), 0) // skill?
	}
	return skills
}

func (s *Skills) setPlayer(p protocol.Player) {
	s.player = p
	for i := range s.skills {
		s.skills[i].player = s.player
		s.skills[i].signalUpdate()
	}
	s.signalUpdate()
}

func (s *Skills) signalUpdate() {
	if s.player != nil {
		// Combat level is technically part of the player's appearance
		s.player.SetAppearanceChanged()
	}
}

//glua:bind accessor
func (s *Skills) CombatLevel() int {
	return s.combatLevel
}

func (s *Skills) SetCombatLevel(combatLevel int) {
	s.combatLevel = combatLevel
	s.signalUpdate()
}

//glua:bind
func (s *Skills) Skill(id protocol.SkillId) *Skill {
	return s.skills[id]
}

//glua:bind
type Skill struct {
	player          protocol.Player
	id              protocol.SkillId
	experience      int
	levelPercentage int
	levelOffset     int
}

//glua:bind constructor Skill
func NewSkill(id protocol.SkillId, experience int) *Skill {
	return &Skill{
		id:              id,
		experience:      experience,
		levelPercentage: 100,
		levelOffset:     0,
	}
}

func (s *Skill) signalUpdate() {
	if s.player != nil {
		s.player.SendSkill(int(s.id), s.EffectiveLevel(), s.experience)
	}
}

//glua:bind accessor
func (s *Skill) Experience() int {
	return s.experience
}

func (s *Skill) SetExperience(exp int) {
	s.experience = exp
	s.signalUpdate()
}

//glua:bind
func (s *Skill) MaximumLevel() int {
	return skillExpToLevel(s.Experience())
}

//glua:bind
func (s *Skill) EffectiveLevel() int {
	return (s.MaximumLevel() + s.levelOffset) * (s.levelPercentage / 100)
}

//glua:bind accessor
func (s *Skill) LevelOffset() int {
	return s.levelOffset
}

func (s *Skill) SetLevelOffset(i int) {
	s.levelOffset = i
}

//glua:bind accessor
func (s *Skill) LevelPercentage() int {
	return s.levelPercentage
}

func (s *Skill) SetLevelPercentage(i int) {
	s.levelPercentage = i
}
