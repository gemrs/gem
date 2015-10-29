gem.engine.event module
=======================

Events related to the :py:class:`gem.engine.Engine`

.. py:data:: Startup

             Raised on engine startup, immediately before the engine starts ticking

.. py:data:: Shutdown

             Raised on engine shutdown, immediately before the engine stops ticking

.. py:data:: PreTick

             Raised during each cycle

.. py:data:: Tick

             Raised during each cycle

.. py:data:: PostTick

             Raised during each cycle
