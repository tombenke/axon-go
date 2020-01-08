axon-cron
=========

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

