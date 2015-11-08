Connection
----------

.. py:currentmodule:: gem.game.server

.. py:class:: Connection

   Connection is a network-level representation of a client.
   It handles read/write buffering, and decodes data into game packets or update requests for processing.

   .. note:: Should not be constructed from Python

   .. rubric:: Attributes

   .. py:attribute:: log

      The :py:class:`gem.log.Module` associated with this connection

   .. rubric:: Methods

   .. py:method:: disconnect()

      Closes the connection forcibly. Non-blocking.
