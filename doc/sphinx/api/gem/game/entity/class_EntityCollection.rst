EntityCollection
----------------

.. py:currentmodule:: gem.game.entity

.. py:class:: EntityCollection

   EntityCollection is an efficient, cycle based entity collection.

   The underlying collection is transactional and is updated at a fixed interval. The update method should be called to commit added/removed entities.

   .. rubric:: Attributes

   .. py:attribute:: adding

      The slice of entities being added in the next cycle

   .. py:attribute:: removing

      The slice of entities being removed in the next cycle

   .. py:attribute:: entities

      The slice of entities currently in the collection

   .. rubric:: Methods

   .. py:method:: add(entity)

      Requests an entity be added to the collection in the next cycle

      :param Entity entity: The entity to add

   .. py:method:: remove(entity)

      Requests an entity be removed from the collection in the next cycle

      :param Entity entity: The entity to remove

   .. py:method:: update()

      Cycles the collection.

      - Entities in the adding list are added to the main entity list
      - Entities in the removing list are removed from the main entity list
      - Both adding and removing lists are emptied
