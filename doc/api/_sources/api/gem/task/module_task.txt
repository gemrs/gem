gem.task module
==================

.. py:module:: gem.task

.. rubric:: Functions

.. py:function:: submit(callback, when, interval, userdata)

   Submits a method or function as a task to the engine.

   Callback should be a function of type ``callback(task, userdata)``, which returns ``True`` or ``False`` indicating whether this task should expire or retrigger.

   :param function callback: The function to call when the task ticks
   :param string when: One of the :ref:`Task Constants <task-constants>`. Specifies when in the engine cycle callback should be called.
   :param int interval: The interval in cycles of 600ms which this task should be triggered.
   :param userdata: An arbitrary object which is passed as an argument to the callback.

.. _task-constants:
.. rubric:: Constants

.. py:data:: PreTick
.. py:data:: Tick
.. py:data:: PostTick
