axon-cron
=========

The `axon-cron` is a publish-only agent, that emits messages in predefined intervals, inluding the actual time. The `type` of the message is configurable, as well as the time precision (nanosec, microsec, millisec, sec), as well as the interval.


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

Execute the agent with the `-h` swith to get help:

```bash
$ axon-cron -h

Usage: axon-cron [-u nats-server] [-creds file] [-s] <subject> [-t] <message-type> [-e] <period> [-p] <precision>
  -creds string
    	User Credentials File
  -cron string
    	Cron definition (default "@every 10s")
  -h	Show help message
  -p string
    	The precision of time value: ns, us, ms, s (default "ns")
  -s string
    	The subject to send the outbound messages (default "axon.cron")
  -t string
    	The type of the message
  -u string
    	The nats server URLs (separated by comma) (default "nats://127.0.0.1:4222")
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

The agent uses the [robfig/cron](https://github.com/robfig/cron) package to define the time inteval, so read [its documentation](https://godoc.org/github.com/robfig/cron) to learn how to define `-cron` parameter of the agent.

