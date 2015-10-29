gem.game.event module
=====================

Events related to game service

.. py:data:: PlayerLogin

             Raised on player login, before :py:attr:`PlayerLoadProfile`

.. py:data:: PlayerLoadProfile

             Raised on player login. Listeners should use this opportunity to apply the player's loaded profile to their current session, eg. warp to the correct position, set appearance


.. py:data:: PlayerFinishLogin

             Raised once the player's profile has loaded; at this point the player is ready to participate in general game logic.

.. py:data:: PlayerLogout

             Raised on player logout after the connection has been terminated.

.. py:data:: EntitySectorChange

             Raised when the player crosses a sector boundary

.. py:data:: EntityRegionChange

             Raised when the player's region changes

.. py:data:: PlayerAppearanceUpdate

             Raised when the player's appearance is updated
