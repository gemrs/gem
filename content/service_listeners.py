import gem
import gem.runite
import gem.archive as archive
import gem.engine.event
import gem.game as game
import gem.game.server as game_server

import config
import core.event

@core.event.listener
class ServiceListeners(object):
    archive_server_started = False
    game_server_started = False

    @core.event.callback(gem.engine.event.Startup)
    def startup(self):
        try:
            self.archive_server = archive.Server()
            self.archive_server.start(config.archive_server_listen, gem.runite.context)
            self.archive_server_started = True
        except Exception as e:
            raise Exception("Couldn't start archive server: {0}".format(e))

        try:
            game_service = game.GameService(gem.runite.context, config.rsa_key_path, config.auth_provider)

            update_service = game.UpdateService(gem.runite.context)

            self.game_server = game_server.Server(config.game_server_listen)
            self.game_server.set_service(14, game_service)
            self.game_server.set_service(15, update_service)
            self.game_server.start()
            self.game_server_started = True
        except Exception as e:
            raise Exception("Couldn't start game server: {0}".format(e))

    @core.event.callback(gem.engine.event.Shutdown)
    def shutdown(self):
        if self.archive_server_started == True:
            self.archive_server.stop()

        if self.game_server_started == True:
            self.game_server.stop()
