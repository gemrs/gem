Context
-------

.. py:currentmodule:: gem.runite

.. py:class:: Context

   Context is a handle referring to an instance of the runite game data access library.

   .. rubric:: Methods

   .. py:method:: unpack(data_file, index_files)

      Loads the given game data files into this runite context.

      :param string data_file: The path to the game data file (.dat)
      :param list index_files: A list of index files for locating files within the data file (.idx*)
