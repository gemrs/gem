Server
------

.. py:currentmodule:: gem.archive

.. py:class:: Server

   Server controls the archive service

   .. rubric:: Methods

   .. py:method:: start(laddr, context)

      Starts the archive service

      :param string laddr: The listen address:port
      :param gem.runite.Context context: The runite context to serve content from

   .. py:method:: stop()

      Stops the archive service
