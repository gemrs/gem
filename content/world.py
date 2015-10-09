import gem.event
import event

logger = gem.syslog.Module(__name__)

@event.listener
class World(object):
    players = {}

    @event.callback(gem.event.PlayerLogin)
    def register_player(self, player):
        self.players[player.Index] = player
        logger.Info("registered player %s" % player)

    @event.callback(gem.event.PlayerLogout)
    def unregister_player(self, player):
        del self.players[player.Index]
        logger.Info("unregistered player %s" % player)

    def get_players(self):
        return self.players.values()

    def player_count(self):
        return len(self.get_players())

global_world = World()

def get_players():
    return global_world.get_players()

def player_count():
    return global_world.player_count()
