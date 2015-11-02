Server
------

.. py:currentmodule:: gem.game.server

.. py:class:: Server(laddr)

   Server handles setting up the tcp server and serving services.

   :param string laddr: The listen address:port

   .. rubric:: Methods

   .. py:method:: set_service(selector, service)

      Registers a service with it's selector id.

      See :py:class:`gem.game.GameService` and :py:class:`gem.game.UpdateService`

      :param int selector: The selector id for this service
      :param Service service: The service which handles this selector id

   .. py:method:: start

      Starts the server listener

   .. py:method:: stop

      Stops the server
