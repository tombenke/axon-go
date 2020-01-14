axon-debug
==========

This is a very simple subscribe-only agent, that accepts any message, and writes that onto the terminal as that is. It provides debugging purposes mostly.

Execute the agent with the `-h` switch to get help:

```bash
$ axon-debug -h

Usage: axon-debug [-u server] [-creds file] [-s] <subject>
  -creds string
    	User Credentials File
  -h	Show help message
  -s string
    	The subject to subscribe for inbound messages (default "axon.log")
  -t	Display timestamps
  -u string
    	The nats server URLs (separated by comma) (default "nats://127.0.0.1:4222")
```
