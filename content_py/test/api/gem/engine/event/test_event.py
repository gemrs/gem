import pytest

import gem.engine.event as ev

def test_game_event_api():
    """Detects changes to the game events"""
    for e in ["Startup", "Shutdown", "PreTick", "Tick", "PostTick"]:
        assert hasattr(ev, e)
