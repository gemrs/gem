EntitySlice
-----------

.. py:currentmodule:: gem.game.entity

.. py:class:: EntitySlice

   EntitySlice is a slice of entities.

   EntitySlices can be added to and emptied, but not removed from. They are intended for buffering entities for addition to a :py:class:`EntityList`.

   .. rubric:: Attributes

   .. py:attribute:: size

      The number of entities in the slice

   .. py:attribute:: list

      A python list containing the entities in the slice.

   .. rubric:: Methods

   .. py:method:: add(entity)

      Adds an entity to the slice

      :param Entity entity: The entity to add

   .. py:method:: empty()

      Empties the slice

   .. py:method:: filter(type)

      Creates a new slice, which is the subset of entities with the given type

      :param type: The entity type to select
      :return: The new filtered slice
      :rtype: EntitySlice
