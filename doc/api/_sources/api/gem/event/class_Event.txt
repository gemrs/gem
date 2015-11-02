Event
------

.. py:currentmodule:: gem.event

.. py:class:: Event(key)

   Event is an object which :py:class:`Listeners <PyListener>` can register to be notified by when certain events occur.

   This is the basis of Gem's event driven API. Many events are provided and triggered by the kernel, however events can also be created and triggered from within python code.

   :param string key: A unique identifier for this event

   .. rubric:: Attributes

   .. py:attribute:: key

      The unique identifier for this event

   .. rubric:: Methods

   .. py:method:: register(listener)

      Subscribes a :py:class:`PyListener` to this event

      :param PyListener listener: The listener to register

   .. py:method:: unregister(listener)

      Removes a :py:class:`PyListener` from the list of listeners for this event

      :param PyListener listener: The listener to unregister

   .. py:method:: notify_observers(*args)

      Notifies all registered listeners that this event was raised, and passes down an arbitrary set of arguments.

      :param list args: an arbitrary list of arguments to provide to listeners
