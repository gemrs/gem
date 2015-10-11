import plugins.gem_plugin as plugins

import gem.event

import world
import event

@event.listener
class MOTD(plugins.GemPlugin):

    @event.callback(gem.event.PlayerLogin)
    def player_login(self, player):
        profile = player.Profile
        player.Session().SendMessage("Welcome to Gielinor")
        player.Session().SendMessage("There are currently {0} players online".format(world.player_count()))

    @event.callback(gem.event.PlayerLogout)
    def player_logout(self, player):
        pass
