import gem.auth

class DummyProvider(gem.auth.ProviderImpl):
    def LookupProfile(self, username, password):
        profile = gem.auth.Profile()
        profile.Username = username
        profile.Password = password
        return profile, gem.auth.AuthOkay
