import plugins.gem_plugin as plugins

import gem.game.event

import world
import core.event

@core.event.listener
class MOTD(plugins.GemPlugin):

    @core.event.callback(gem.game.event.PlayerLogin)
    def player_login(self, player):
        player.send_message("Welcome to Gielinor")
        player.send_message("There are currently {0} players online".format(world.player_count()))

    @core.event.callback(gem.game.event.PlayerLogout)
    def player_logout(self, player):
        pass
