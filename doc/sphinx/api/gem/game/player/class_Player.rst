Player
------

.. py:currentmodule:: gem.game.player

.. py:class:: Player

   Player is self explanatory

   Player inherits methods and attributes of :py:class:`gem.game.server.Connection` and :py:class:`gem.game.entity.GenericMob`

   .. note:: Should not be constructed from Python

   .. rubric:: Attributes

   .. py:attribute:: username

      The player's username

   .. py:attribute:: appearance

      The :py:class:`Appearance` associated with this player

   .. py:attribute:: skills

      The :py:class:`Skills` associated with this player

   .. py:attribute:: entity_type

   .. rubric:: Methods

   .. py:method:: warp(position)

      Warps the player to a given location

      :param gem.game.position.Absolute position: The position to warp to

   .. py:method:: send_message(message)

      Puts a message to the player's chat box
