SysLog
------

.. py:currentmodule:: gem.log

.. py:class:: SysLog

   SysLog is a log object which can contain contextual information

   .. note:: Should not be constructed directly. Instead, use the kernel-provided :py:attr:`gem.syslog`

   .. rubric:: Methods

   .. py:method:: module(prefix)

      Constructs a submodule of this log object, using prefix as a prefix for all log messages

      :param string prefix: The prefix to prepend all log messages with

   .. py:method:: begin_redirect

      Starts redirecting all log output to an in-memory buffer. Used by the interactive console to avoid log messages printing over the prompt.

   .. py:method:: end_redirect

      Ends the log redirect and resumes logging to stdout. The in-memory buffer created by :py:meth:`~SysLog.begin_redirect` is flushed to stdout and emptied.
