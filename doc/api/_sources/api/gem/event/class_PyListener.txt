PyListener
----------

.. py:currentmodule:: gem.event

.. py:class:: PyListener(callback)

   A PyListener is an observer of a :py:class:`Event` which is tied to a python function or method.

   When the listener is notified, callback is called with the list of arguments passed to :py:meth:`~Event.notify_observers`

   :param function callback: The function to call when this listener is notified

   .. rubric:: Attributes

   .. py:attribute:: id

      The unique identifier for this listener
