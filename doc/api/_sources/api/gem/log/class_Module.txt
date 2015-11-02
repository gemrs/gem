Module
------

.. py:currentmodule:: gem.log

.. py:class:: Module

   Module is a log object which can contain contextual information

   .. note:: Should not be constructed directly. Instead, one should derive a Module from an existing logger using either :py:meth:`Module.submodule` or :py:meth:`SysLog.module`

   .. rubric:: Methods

   .. py:method:: submodule(prefix)

      Constructs a submodule of this log object, using prefix as a prefix for all log messages

      :param string prefix: The prefix to prepend all log messages with

   .. py:method:: critical(message)

      :param sting message: The log message

   .. py:method:: debug(message)

      :param sting message: The log message

   .. py:method:: error(message)

      :param sting message: The log message

   .. py:method:: info(message)

      :param sting message: The log message

   .. py:method:: notice(message)

      :param sting message: The log message

   .. py:method:: warning(message)

      :param sting message: The log message
