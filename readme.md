
# Ironclad

Ironclad is a command line tool for creating and managing encrypted password databases.

    $ ironclad --help

    Usage: ironclad [FLAGS] [COMMAND]

      A utility for creating and managing encrypted password databases.

    Flags:
      --help            Print the application's help text and exit.
      --version         Print the application's version number and exit.

    Commands:
      add               Add a new entry to a database.
      config            Set or print a configuration option.
      decrypt           Decrypt a file.
      delete            Delete entries from a database.
      dump              Dump a database's internal JSON data store.
      edit              Edit an existing database entry.
      encrypt           Encrypt a file.
      export            Export entries from a database.
      gen               Generate a random password.
      import            Import entries into a database.
      list              List database entries.
      new               Create a new database.
      pass              Copy a password to the clipboard.
      purge             Purge deleted entries from a database.
      tags              List database tags.
      user              Copy a username to the clipboard.

    Command Help:
      help <command>    Print the specified command's help text and exit.

See the application's [homepage][] for details and binary downloads.

[homepage]: http://mulholland.xyz/dev/ironclad/
