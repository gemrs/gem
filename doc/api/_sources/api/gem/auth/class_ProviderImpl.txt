ProviderImpl
------------

.. py:currentmodule:: gem.auth

.. py:class:: ProviderImpl

   Base class for authentication providers

   .. warning:: Should not be used directly. This is a base class whose methods should be implemented by your own python class.

   .. seealso:: Your auth provider implementation should be passed as a parameter to :py:class:`gem.game.GameService`

   .. rubric:: Methods

   .. py:method:: load_profile(username, password)

      Authenticates and loads the player's profile.

      :param string username: The player's username
      :param string password: The encrypted form of the player's password
      :return: Tuple of profile and auth response
      :rtype: (:py:class:`gem.game.player.Profile`, :ref:`auth response <auth-constants>`)
