import gem.log
import gem.game.event
import core.event

logger = gem.log.Module(__name__, None)

@core.event.listener
class World(object):
    players = {}

    @core.event.callback(gem.game.event.PlayerLoadProfile)
    def load_profile(self, player, profile):
        pass

    @core.event.callback(gem.game.event.PlayerLogin)
    def register_player(self, player):
        self.players[player.index] = player
        logger.info("registered player [%s]" % player)

    @core.event.callback(gem.game.event.PlayerLogout)
    def unregister_player(self, player):
        del self.players[player.index]
        logger.info("unregistered player [%s]" % player)

    def get_players(self):
        return self.players.values()

    def player_count(self):
        return len(self.get_players())

global_world = World()

def get_players():
    return global_world.get_players()

def player_count():
    return global_world.player_count()
