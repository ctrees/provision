drpcli reservations
===================

Access CLI commands relating to reservations

Synopsis
--------

Access CLI commands relating to reservations

Options
-------

::

      -h, --help   help for reservations

Options inherited from parent commands
--------------------------------------

::

      -d, --debug             Whether the CLI should run in debug mode
      -E, --endpoint string   The Digital Rebar Provision API endpoint to talk to (default "https://127.0.0.1:8092")
      -f, --force             When needed, attempt to force the operation - used on some update/patch calls
      -F, --format string     The serialzation we expect for output.  Can be "json" or "yaml" (default "json")
      -P, --password string   password of the Digital Rebar Provision user (default "r0cketsk8ts")
      -r, --ref string        A reference object for update commands that can be a file name, yaml, or json blob
      -T, --token string      token of the Digital Rebar Provision access
      -U, --username string   Name of the Digital Rebar Provision user to talk to (default "rocketskates")

SEE ALSO
--------

-  `drpcli <drpcli.html>`__ - A CLI application for interacting with the
   DigitalRebar Provision API
-  `drpcli reservations create <drpcli_reservations_create.html>`__ -
   Create a new reservation with the passed-in JSON or string key
-  `drpcli reservations destroy <drpcli_reservations_destroy.html>`__ -
   Destroy reservation by id
-  `drpcli reservations exists <drpcli_reservations_exists.html>`__ -
   See if a reservations exists by id
-  `drpcli reservations indexes <drpcli_reservations_indexes.html>`__ -
   Get indexes for reservations
-  `drpcli reservations list <drpcli_reservations_list.html>`__ - List
   all reservations
-  `drpcli reservations show <drpcli_reservations_show.html>`__ - Show a
   single reservations by id
-  `drpcli reservations update <drpcli_reservations_update.html>`__ -
   Unsafely update reservation by id with the passed-in JSON
-  `drpcli reservations wait <drpcli_reservations_wait.html>`__ - Wait
   for a reservation's field to become a value within a number of
   seconds
