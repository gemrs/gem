Profile
-------

.. py:currentmodule:: gem.game.player

.. py:class:: Profile

   Profile represents the persistent state of a player

   .. note:: Should not be constructed from Python

   .. rubric:: Methods

   .. py:method:: serialize()

      Serializes the profile to JSON

      :return: The serialized object
      :rtype: string

   .. py:method:: deserialize(object)

      Deserializes the player's profile from JSON

      :param string object: The serialized object

   .. rubric:: Attributes

   .. py:attribute:: appearance

      The :py:class:`Appearance` associated with this player

   .. py:attribute:: skills

      The :py:class:`Skills` associated with this player

   .. py:attribute:: position

      The position of the player
