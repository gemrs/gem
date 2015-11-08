gem.log module
==================

.. py:module:: gem.log

.. rubric:: Functions

.. py:function:: begin_redirect

Starts redirecting all log output to an in-memory buffer. Used by the interactive console to avoid log messages printing over the prompt.

.. py:function:: end_redirect

Ends the log redirect and resumes logging to stdout. The in-memory buffer created by :py:func:`begin_redirect` is flushed and emptied.

.. rubric:: Classes

.. toctree::
   :maxdepth: 3
   :glob:

   class_*
