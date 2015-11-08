EntityList
----------

.. py:currentmodule:: gem.game.entity

.. py:class:: EntityList

   List is a list of entities, efficient for searching by entity index.

   .. rubric:: Attributes

   .. py:attribute:: slice

      The equivalent :py:class:`EntitySlice` of this List.

   .. py:attribute:: size

      The number of entities in the list

   .. rubric:: Methods

   .. py:method:: add(entity)

      Adds an entity to the list

      :param Entity entity: The entity to add

   .. py:method:: remove(entity)

      Removes an entity from the list

      :param Entity entity: The entity to remove

   .. py:method:: add_all(slice)

      Adds a slice of entities to the list

      :param EntitySlice slice: The slice of entities to add

   .. py:method:: remove_all(slice)

      Removes a slice of entities from the list

      :param EntitySlice slice: The slice of entities to remove
