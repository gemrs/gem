import gem.log
import gem.game.event
import gem.game.entity
import core.event
import core.task

logger = gem.log.Module(__name__, None)

class CollectionUpdateTask(core.task.Task):
    def __init__(self, collection):
        super(CollectionUpdateTask, self).__init__(1, gem.task.PostTick)
        self.collection = collection

    def execute(self):
        self.collection.update()
        return True

@core.event.listener
class World(object):
    entities = gem.game.entity.EntityCollection()
    entity_update_task = CollectionUpdateTask(entities)

    def __init__(self):
        self.entity_update_task.submit()
        self.instance = None

    @core.event.callback(gem.game.event.PlayerLoadProfile)
    def load_profile(self, player, profile):
        pass

    @core.event.callback(gem.game.event.PlayerLogin)
    def register_player(self, player):
        self.entities.add(player)
        logger.info("registered player [%s]" % player)

    @core.event.callback(gem.game.event.PlayerLogout)
    def unregister_player(self, player):
        self.entities.remove(player)
        logger.info("unregistered player [%s]" % player)

    def set_instance(self, world_instance):
        self.instance = world_instance

    def get_players(self):
        entity_list = self.entities.entities
        player_list = entity_list.filter(gem.game.entity.PlayerType)
        return player_list.list

    def player_count(self):
        return len(self.get_players())

global_world = World()

def get_players():
    return global_world.get_players()

def player_count():
    return global_world.player_count()
