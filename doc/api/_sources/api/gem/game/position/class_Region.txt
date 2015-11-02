Region
------

.. py:currentmodule:: gem.game.position

.. py:class:: Region(origin)

   A region is a 13x13 sector (104x104 tile) chunk.
   This is mainly used to represent the loaded area of the map around the player

   :param Sector origin: The sector located at the lowest corner of the region

   .. rubric:: Attributes

   .. py:attribute:: origin

      The :py:class:`Sector` located at the lowest corner of the region (ie. NOT the center of the region)

   .. rubric:: Methods

   .. py:method:: rebase(position)

      Rebase adjusts the region such that it;s new center is the sector containing the given position

      :param Absolute position: The coordinate to rebase onto
