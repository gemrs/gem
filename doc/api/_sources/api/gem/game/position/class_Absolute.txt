Absolute
--------

.. py:currentmodule:: gem.game.position

.. py:class:: Absolute(x, y, z)

   Represents an absolute coordinate within the world

   .. rubric:: Attributes

   .. py:attribute:: x

   .. py:attribute:: y

   .. py:attribute:: z

   .. py:attribute:: sector

      The :py:class:`Sector` containing this coordinate

   .. py:attribute:: region

      The :py:class:`Region` centered at this position

   .. rubric:: Methods

   .. py:method:: local_to(region)

      Calculates the local coordinates relative to a :py:class:`Region`

      :param Region region: The region which the resulting local coordinates should be relative to
      :return: The local coordinates
      :rtype: Local
