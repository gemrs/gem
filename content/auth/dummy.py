import gem.auth
import gem.service.game.player as player

class DummyProvider(gem.auth.ProviderImpl):
    def LookupProfile(self, username, password):
        profile = player.Profile()
        profile.Username = username
        profile.Password = password
        return profile, gem.auth.AuthOkay
