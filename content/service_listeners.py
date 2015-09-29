import gem
import gem.runite
import gem.service.archive as archive
import gem.service.game as game

import config
from event_listener import EventListener

class ServiceListeners(EventListener):
    archive_server_started = False
    game_server_started = False

    def __init__(self):
        pass

    def startup(self, event):
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


    def shutdown(self, event):
        if self.archive_server_started == True:
            self.archive_server.Stop()

        if self.game_server_started == True:
            self.game_server.Stop()

    def player_login(self, event, player):
        profile = player.Profile
        print "LOGIN! {0}".format(profile.Username)

    def player_logout(self, event, player):
        pass
