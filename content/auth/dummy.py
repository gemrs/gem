import gem
import gem.auth
import gem.game.player as player

logger = gem.syslog.Module(__name__)

class DummyProvider(gem.auth.ProviderImpl):
    def LookupProfile(self, username, password):
        if username == "x" and password == "x":
            logger.Info("invalid login credentials")
            return player.Profile(), gem.auth.AuthInvalidCredentials

        profile = player.Profile()
        profile.Username = username
        profile.Password = password
        return profile, gem.auth.AuthOkay
