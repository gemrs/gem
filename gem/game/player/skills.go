package player

type Skills struct {
	player *Player

	combatLevel int
}

func NewSkills() *Skills {
	return &Skills{}
}

func (s *Skills) setPlayer(p *Player) {
	s.player = p
	s.signalUpdate()
}

func (s *Skills) signalUpdate() {
	if s.player != nil {
		// Combat level is technically part of the player's appearance
		s.player.AppearanceUpdated()
	}
}

func (s *Skills) CombatLevel() int {
	return s.combatLevel
}

func (s *Skills) SetCombatLevel(combatLevel int) {
	s.combatLevel = combatLevel
	s.signalUpdate()
}
