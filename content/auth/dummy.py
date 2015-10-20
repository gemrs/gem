import gem
import gem.auth
import gem.game as game
import gem.game.position as position

logger = gem.syslog.Module(__name__)

class DummyProvider(gem.auth.ProviderImpl):
    def LookupProfile(self, username, password):
        if username == "x" and password == "x":
            logger.Info("invalid login credentials")
            return game.Profile(), gem.auth.AuthInvalidCredentials

        profile = game.Profile(username, password)
        profile.SetPosition(position.Absolute(3200, 3200, 0))
        profile.SetAppearance(game.Appearance())
        return profile, gem.auth.AuthOkay
