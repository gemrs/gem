EntityCollection
----------------

.. py:currentmodule:: gem.game.entity

.. py:class:: EntityCollection

   EntityCollection is an efficient, cycle based entity collection.

   The underlying collection is transactional and is updated at a fixed interval. The update method should be called to commit added/removed entities.

   .. rubric:: Attributes

   .. py:attribute:: adding

      The :py:class:`EntitySlice` of entities being added in this cycle

   .. py:attribute:: removing

      The :py:class:`EntitySlice` of entities being removed in this cycle

   .. py:attribute:: entities

      The :py:class:`EntitySlice` of entities currently in the collection. Includes those which exist in the adding list, but not those in the removing list.

   .. rubric:: Methods

   .. py:method:: add(entity)

      Requests an entity be added to the collection in the next cycle. The new entity goes into the tracking list, and to the adding list.

      :param Entity entity: The entity to add

   .. py:method:: remove(entity)

      Requests an entity be removed from the collection in the next cycle. The entity is removed from the tracking list, and added to the removing list.

      :param Entity entity: The entity to remove

   .. py:method:: update()

      Cycles the collection. Both adding and removing lists are emptied.
