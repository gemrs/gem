import gem
import gem.auth
import gem.game
import gem.game.player
import gem.game.position

logger = gem.syslog.module(__name__)

class DummyProvider(gem.auth.ProviderImpl):
    def lookup_profile(self, username, password):
        if username == "x" and password == "x":
            logger.info("invalid login credentials")
            return gem.game.player.profile, gem.auth.AuthInvalidCredentials

        profile = gem.game.player.Profile(username, password)
        profile.position = gem.game.position.Absolute(3200, 3200, 0)
        profile.appearance = gem.game.player.Appearance()
        return profile, gem.auth.AuthOkay
