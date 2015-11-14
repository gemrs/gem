Area
------

.. py:currentmodule:: gem.game.position

.. py:class:: Area(origin)

   An Area is a rectangular grouping of sectors, with an arbitrary width and height.

   :param gem.game.position.Sector origin: The sector located at the lowest corner of the area
   :param int w: The width of the area in sectors
   :param int h: The height of the area in sectors

   .. rubric:: Attributes

   .. py:attribute:: min_sector

      The minimum :py:class:`gem.game.position.Sector` of this area

   .. py:attribute:: max_sector

      The maximum :py:class:`gem.game.position.Sector` of this area

   .. py:attribute:: min

      The :py:class:`Absolute` of the minimum tile in this area

   .. py:attribute:: max

      The :py:class:`Absolute` of the maximum tile in this area

   .. rubric:: Methods

   .. py:method:: contains(position)

      Determines if a given tile is within this area.

      :param Absolute position: The coordinate to check
      :return: True if position is within this area
      :rtype: boolean
