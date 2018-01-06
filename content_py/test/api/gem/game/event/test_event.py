import pytest

import gem.game.event as ev

def test_game_event_api():
    """Detects changes to the game events"""
    for e in ["PlayerLoadProfile", "PlayerLogin", "PlayerLogout",
              "PlayerFinishLogin", "EntitySectorChange", "EntityRegionChange",
              "PlayerAppearanceUpdate"]:
        assert hasattr(ev, e)
