package impl

import "github.com/gemrs/gem/gem/protocol"

//glua:bind
type ClientConfig struct {
	player        *Player
	TabInterfaces map[protocol.InterfaceTab]int
}

//glua:bind constructor ClientConfig
func NewClientConfig(player *Player) *ClientConfig {
	return &ClientConfig{
		player:        player,
		TabInterfaces: make(map[protocol.InterfaceTab]int),
	}
}

//glua:bind
func (c *ClientConfig) TabInterface(id protocol.InterfaceTab) int {
	return c.TabInterfaces[id]
}

//glua:bind
func (c *ClientConfig) SetTabInterface(tab protocol.InterfaceTab, id int) {
	c.TabInterfaces[tab] = id
	c.player.sendTabInterface(tab, id)
}
