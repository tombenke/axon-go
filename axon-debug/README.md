axon-debug
==========

This is a very simple subscribe-only agent, that accepts any message,
and writes that onto the terminal as that is.
It provides debugging purposes mostly.

Execute the agent with the `-h` switch to get help:


Execute the agent with the `-h` switch to get help:

```bash
    Usage: axon-debug -h

      -c string
            User Credentials
      -config string
            Config file name (default "config.yml")
      -creds string
            User Credentials
      -f string
            The log format: json | text (default "json")
      -h	Show help message
      -help
            Show help message
      -in value
            Input. Format: <name>[|<channel>[|<type>|<representation>|<default>]]
      -l string
            The log level: panic | fatal | error | warning | info | debug | trace (default "debug")
      -log-format string
            The log format: json | text (default "json")
      -log-level string
            The log level: panic | fatal | error | warning | info | debug | trace (default "debug")
      -messaging-urls string
            The Messaging server's URLs (separated by comma) (default "localhost:4222")
      -n string
            The name of the node (default "axon-debug")
      -name string
            The name of the node (default "axon-debug")
      -out value
            Output. Format: <name>[|<channel>[|<type>|<representation>]]
      -p	Print configuration parameters
      -print-config
            Print configuration parameters
      -u string
            The Messaging server's URLs (separated by comma) (default "localhost:4222")
```
