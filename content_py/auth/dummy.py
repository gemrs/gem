import gem
import gem.auth
import gem.game
import gem.game.player
import gem.game.position
import gem.log

logger = gem.log.Module(__name__, None)

profile_template = """{
  "username": "",
  "password": "",
  "rights": 0,
  "position": {
    "x": 3250,
    "y": 3200,
    "z": 0
  },
  "skills": {
    "combat_level": 0
  },
  "appearance": {
    "gender": 0,
    "head_icon": 0,
    "model_torso": 19,
    "model_arms": 29,
    "model_legs": 39,
    "model_head": 3,
    "model_hands": 35,
    "model_feet": 44,
    "model_beard": 10,
    "color_hair": 7,
    "color_torso": 8,
    "color_legs": 9,
    "color_feet": 5,
    "color_skin": 0
  }
}
"""

class DummyProvider(gem.auth.ProviderImpl):
    def lookup_profile(self, username, password):
        if username == "x" and password == "x":
            logger.info("invalid login credentials")
            return gem.game.player.profile, gem.auth.AuthInvalidCredentials

        profile = gem.game.player.Profile(username, password)
        profile.deserialize(profile_template)
        return profile, gem.auth.AuthOkay
