import gem
import gem.runite
import gem.archive as archive
import gem.game as game

import config
import event

@event.listener
class ServiceListeners(object):
    archive_server_started = False
    game_server_started = False

    @event.callback(gem.event.Startup)
    def startup(self):
        try:
            self.archive_server = archive.Server()
            self.archive_server.Start(config.archive_server_listen, gem.runite.context)
            self.archive_server_started = True
        except Exception as e:
            raise Exception("Couldn't start archive server: {0}".format(e))

        try:
            self.game_server = game.Server()
            self.game_server.Start(config.game_server_listen, gem.runite.context, config.rsa_key_path, config.auth_provider)
            self.game_server_started = True
        except Exception as e:
            raise Exception("Couldn't start game server: {0}".format(e))

    @event.callback(gem.event.Shutdown)
    def shutdown(self):
        if self.archive_server_started == True:
            self.archive_server.Stop()

        if self.game_server_started == True:
            self.game_server.Stop()
