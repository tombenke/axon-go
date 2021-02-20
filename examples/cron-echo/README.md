The `cron-echo` flow
====================

## Prerequisites

In order to run this demo, you need a messaging middleware running.
You can start it via the following command:

```bash
    docker run -it --rm --network=host -p 4222:4222 nats-streaming:latest
```

## The `cron-echo` flow

The [`cron-echo/`](./) folder contains the configuration and start script,
that demonstrates the most primitive working example of an axon flow.

This flow is made of a producer agent, called `axon-cron`, that produces a message every 1 seconds
and it sends the message into the `$CRON_CHANNEL` subject.

Another agent, called `axon-debug` subscribes to the same topic,
and receives the messages sent by the `axon-cron`, then prints out the message to the console.

The figure below shows the flow diagram:

![The `cron-echo` flow diagram](../../docs/cron-echo-flow-diagram.png)

You can execute the flow with the following two commands in two separate terminal windows,
first start the `axon-debug`, then the `axon-cron`. You will see something like these:

```bash
    $ axon-cron -config axon-cron-config.yml
```

```bash
    axon-debug -config axon-debug-config.yml

```

You also can start the same flow with only one command,
using the [`forego`](https://github.com/ddollar/forego) process manager:

```bash
    cd examples/cron-echo
    forego start -f Procfile

```
