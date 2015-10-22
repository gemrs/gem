import plugins.gem_plugin as plugins

import gem.game.event

import world
import event

@event.listener
class MOTD(plugins.GemPlugin):

    @event.callback(gem.game.event.PlayerLogin)
    def player_login(self, player):
        profile = player.Profile
        player.SendMessage("Welcome to Gielinor")
        player.SendMessage("There are currently {0} players online".format(world.player_count()))

    @event.callback(gem.game.event.PlayerLogout)
    def player_logout(self, player):
        pass
