axon-inject
===========

## About

The `axon-inject` is a producer-only agent, that emits messages into a topic.

## Message sources

The messages may come from several kind of sources, such as:

- A string parameter (plain text, or JSON format).
- A file which holds the message(s) in the following formats:

    - A plain text file, 
    - a single JSON document (single-line, or mult-line),
    - NDJSON (multiple JSON documents, each one is in a single line, closed by the `\n` character),
    - a YAML document (single, or multiple documents, separated by the `---\n` sequence),

The actor tries to determine the input file's content-type from its extension,
if it is not defined by the `--file-type <type>` parameter.

If the messages are coming from a file, it can be a physical file, or the `<stdin>` as well.

The output message-type is `base/Bytes` and the representation format is `text/plain`.

### Synchronous vs. asynchronous mode

The actor can inject the messages both in asynchronous and synchronous mode. The default is the asynchronous mode.

In case it is configured to be asynchronous, then all of the messages will be sent as soon as possible.
The sending period can be controlled by the `--delay <value>[us|ms]` parameter.
If this parameter is given, the actor will wait before sending the next message, if there is any.
The delay value can be defined either in microseconds (`usec`), or milliseconds (`msec`) using the appropriate postfix attached to the value.
The default value of the delay is 0.

If it is in synchronous mode, then the actor will synchronize the sending with the orchestrator.
It sendd the next message, when the orchestrator sends a message to the EPN nodes via the `send-results` channel,
and send the status report of sending via the `sending-completed` channel, 

The actor exits, if it is completed the sending, and there is no more message left in the file.

## Get help

Execute the agent with the `-h` switch to get help:

```bash
    Usage: axon-inject -h
```

