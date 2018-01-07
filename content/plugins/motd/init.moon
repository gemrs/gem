event = require "gem.event"
game_event = require "gem.game.event"

game_event.player_finish_login\register event.Func (event, player) ->
  player\send_message("Welcome to Gielinor!")
