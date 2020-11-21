axon-file-reader
================

The `axon-file-reader` is a producer-only agent, that emits one message, that is the content of a file selected by the argument.
The `type` of the message is configurable.

This is an example that the `axon-file-reader` agent emits:

```JavaScript
    {
        "type":"file-content",
        "meta":{
            "timePrecision":"ns"
        },
        "body": {
            // The byte stream of the file
        }
    }
```

Execute the agent with the `-h` switch to get help:

```bash
$ axon-file-reader -h

Usage: axon-file-reader [-u nats-server] [-creds file] [-s] <subject> [-t] <message-type> [-f] <file-path> [-p] <precision>
  -creds string
    	User Credentials File
  -h	Show help message
  -s string
    	The subject to send the outbound messages (default "axon.file-reader")
  -t string
    	The type of the message
  -f string
        The path to the file to read, and send
  -u string
    	The nats server URLs (separated by comma) (default "nats://127.0.0.1:4222")
```

