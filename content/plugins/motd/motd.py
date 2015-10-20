import plugins.gem_plugin as plugins

import gem.game

import world
import event

@event.listener
class MOTD(plugins.GemPlugin):

    @event.callback(gem.game.PlayerLogin)
    def player_login(self, player):
        profile = player.Profile
        player.SendMessage("Welcome to Gielinor")
        player.SendMessage("There are currently {0} players online".format(world.player_count()))

    @event.callback(gem.game.PlayerLogout)
    def player_logout(self, player):
        pass
