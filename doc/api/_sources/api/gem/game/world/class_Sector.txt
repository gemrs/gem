Sector
------

.. py:currentmodule:: gem.game.world

.. py:class:: Sector

   Sector is an instance of an 8x8 chunk of the world. It's primary function is to track which entities are in the sector.

   .. note:: Should not be constructed from Python

   .. rubric:: Attributes

   .. py:attribute:: active

      Whether this sector is active. Inactive sectors are garbage collected when :py:meth:`WorldInstance.gc` is called.
