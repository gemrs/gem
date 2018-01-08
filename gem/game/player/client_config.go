package player

//glua:bind
const (
	TabAttack int = iota
	TabSkills
	TabQuests
	TabInventory
	TabEquipment
	TabPrayer
	TabMagic
	TabUnused
	TabFriends
	TabIgnore
	TabLogout
	TabSettings
	TabRun
	TabMusic
)

//glua:bind
type ClientConfig struct {
	player        *Player
	TabInterfaces map[int]int
}

//glua:bind constructor ClientConfig
func NewClientConfig(player *Player) *ClientConfig {
	return &ClientConfig{
		player:        player,
		TabInterfaces: make(map[int]int),
	}
}

//glua:bind
func (c *ClientConfig) TabInterface(id int) int {
	return c.TabInterfaces[id]
}

//glua:bind
func (c *ClientConfig) SetTabInterface(tab int, id int) {
	c.TabInterfaces[tab] = id
	c.player.sendTabInterface(tab, id)
}
