Player
------

.. py:currentmodule:: gem.game.player

.. py:class:: Player

   Player is self explanatory

   Player inherits methods and attributes of :py:class:`gem.game.server.Connection`

   .. note:: Should not be constructed from Python

   .. rubric:: Attributes

   .. py:attribute:: appearance

      The :py:class:`Appearance` associated with this player

   .. py:attribute:: session

      The :py:class:`Session` associated with this player

   .. py:attribute:: profile

      The :py:class:`Profile` associated with this player

   .. py:attribute:: entity_type

   .. rubric:: Methods

   .. py:method:: warp(position)

      Warps the player to a given location

      :param gem.game.position.Absolute position: The position to warp to

   .. py:method:: send_message(message)

      Puts a message to the player's chat box
