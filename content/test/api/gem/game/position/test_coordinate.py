import pytest

import gem.game.position as pos

def test_coord_api():
    """Detects changes to the coordinate api"""
    abs = pos.Absolute(3200, 3250, 1)
    assert abs.x == 3200
    assert abs.y == 3250
    assert abs.z == 1
    sector = abs.sector
    assert sector.x == 400
    assert sector.y == 406
    assert sector.z == 1
    region = abs.region
    origin = region.origin
    assert origin.x == 394
    assert origin.y == 400
    assert origin.z == 1
    local = abs.local_to(region)
    assert local.x == 48
    assert local.y == 50
    assert local.z == 1
    region = local.region
    sector = region.origin
    region.rebase(pos.Absolute(abs.x+20, abs.y+20, abs.z))
