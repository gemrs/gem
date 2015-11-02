GameService
-----------

.. py:currentmodule:: gem.game

.. py:class:: GameService(context, rsa_key, auth_provider)

   GameService is the service handling game clients

   :param gem.runite.Context context: The runite context to serve content from
   :param string rsa_key: The path to the RSA keypair to use for secure login
   :param gem.auth.ProviderImpl auth_provider: The object implementing :py:class:`gem.auth.ProviderImpl` to use for player authentication
