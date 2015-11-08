Module
------

.. py:currentmodule:: gem.log

.. py:class:: Module(tag, context)

   Module is a log object which can contain contextual information

   :param string tag: The string with which to tag all log records created by this logger
   :param LogContext context: A context provider for all log records created by this logger. May be None.

   .. rubric:: Attributes

   .. py:attribute:: tag

      The log tag

   .. py:attribute:: context

      The context provider

   .. rubric:: Methods

   .. py:method:: child(prefix)

      Constructs a submodule of this log object, using prefix as a prefix for all log messages

      :param string prefix: The prefix to prepend all log messages with

   .. py:method:: debug(message)

      :param sting message: The log message

   .. py:method:: error(message)

      :param sting message: The log message

   .. py:method:: info(message)

      :param sting message: The log message

   .. py:method:: notice(message)

      :param sting message: The log message
