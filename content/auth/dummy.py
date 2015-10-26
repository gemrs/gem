import gem
import gem.auth
import gem.game
import gem.game.player
import gem.game.position

logger = gem.syslog.Module(__name__)

class DummyProvider(gem.auth.ProviderImpl):
    def LookupProfile(self, username, password):
        if username == "x" and password == "x":
            logger.Info("invalid login credentials")
            return gem.game.player.Profile(), gem.auth.AuthInvalidCredentials

        profile = gem.game.player.Profile(username, password)
        profile.SetPosition(gem.game.position.Absolute(3200, 3200, 0))
        profile.SetAppearance(gem.game.player.Appearance())
        return profile, gem.auth.AuthOkay
