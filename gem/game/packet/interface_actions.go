package packet

import (
	"github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

func init() {
	registerHandler((*protocol.InboundWidgetAction)(nil), player_widget_action)
}

func player_widget_action(p protocol.Player, message server.Message) {
	action := message.(*protocol.InboundWidgetAction)

	game_event.PlayerWidgetAction.NotifyObservers(p, action.Action, action.InterfaceID, action.WidgetID, action.Param, action.ItemID)
}
