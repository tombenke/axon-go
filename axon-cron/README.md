axon-cron
=========

The `axon-cron` is a producer-only agent, that emits messages in predefined intervals,
inluding the actual time. The `type` of the message is configurable, 
as well as the time precision (nanosec, microsec, millisec, sec), as well as the interval.

This is an example that the `axon-cron` agent emits:

```JavaScript
    {
        "type":"measure",
        "meta":{
            "timePrecision":"ns"
        },
        "body": {
            "time":1578563641000543512
        }
    }
```

Execute the agent with the `-h` switch to get help:

```bash
    Usage: axon-cron -h

      -c string
            User Credentials
      -config string
            Config file name (default "config.yml")
      -creds string
            User Credentials
      -cron string
            Cron definition (default "@every 20s")
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
            The name of the node (default "axon-cron")
      -name string
            The name of the node (default "axon-cron")
      -out value
            Output. Format: <name>[|<channel>[|<type>|<representation>]]
      -p	Print configuration parameters
      -precision string
            The precision of time value: ns, us, ms, s (default "ns")
      -print-config
            Print configuration parameters
      -u string
            The Messaging server's URLs (separated by comma) (default "localhost:4222")
```

You can control the interval with the `-cron` parameter.
Some examples:

```bash
    -cron "30 * * * *"                      # Every hour on the half hour
    -cron "30 3-6,20-23 * * *"              # .. in the range 3-6am, 8-11pm
    -cron "CRON_TZ=Asia/Tokyo 30 04 * * *"  # Runs at 04:30 Tokyo time every day
    -cron "@hourly"                         # Every hour, starting an hour from now
    -cron "@every 1h30m"                     # Every hour thirty, starting an hour thirty from now
```

The agent uses the [robfig/cron](https://github.com/robfig/cron) package to define the time inteval,
so read [its documentation](https://godoc.org/github.com/robfig/cron)
to learn how to define `-cron` parameter of the agent.

