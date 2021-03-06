drpcli machines tasks
=====================

Access task manipulation for machines

Synopsis
--------

Access task manipulation for machines

Options
-------

::

      -h, --help   help for tasks

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

-  `drpcli machines <drpcli_machines.html>`__ - Access CLI commands
   relating to machines
-  `drpcli machines tasks add <drpcli_machines_tasks_add.html>`__ - Add
   tasks to the task list for [id]
-  `drpcli machines tasks del <drpcli_machines_tasks_del.html>`__ -
   Remove tasks from the mutable part of the task list
