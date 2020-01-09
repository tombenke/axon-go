The `cron-echo` flow
====================

### The `cron-echo` flow

The [`cron-echo/`](./) folder contains the configuration and start script, that demonstrates the most primitive working example of an axon flow, that is made of a producer agent, called `axon-cron`, which produces a message every 1 seconds and sends it into the `axon.test42.watch` subject. Another agent, called `axon-debug` subscribes to the same topic, and receives the messages sent by the `axon-cron`, then prints out the message to the console.

The figure below shows the flow diagram:

![The `cron-echo` flow diagram](../../docs/cron-echo-flow-diagram.png)

You can execute the flow with the following two commands in two separate terminal windows, first start the `axon-debug`, then the `axon-cron`. You will see something like these:

```bash
    axon-debug -u demo.nats.io:4222 -s axon.test42.watch
    
    Listening on [axon.test42.watch]
    {"time":1578428984000373745,"type":"heartbeat","meta":{"timePrecision":"ns"}}
    {"time":1578428985000158661,"type":"heartbeat","meta":{"timePrecision":"ns"}}
    {"time":1578428986000153381,"type":"heartbeat","meta":{"timePrecision":"ns"}}
    {"time":1578428987000168817,"type":"heartbeat","meta":{"timePrecision":"ns"}}
```

```bash
    axon-cron -u demo.nats.io:4222 -s axon.test42.watch -t "heartbeat" -cron "@every 1s"

    Published [axon.test42.watch] : '{"time":1578428984000373745,"type":"heartbeat","meta":{"timePrecision":"ns"}}'
    Published [axon.test42.watch] : '{"time":1578428985000158661,"type":"heartbeat","meta":{"timePrecision":"ns"}}'
    Published [axon.test42.watch] : '{"time":1578428986000153381,"type":"heartbeat","meta":{"timePrecision":"ns"}}'
    Published [axon.test42.watch] : '{"time":1578428987000168817,"type":"heartbeat","meta":{"timePrecision":"ns"}}'
```

You also can start the same flow with only one command, using the [`forego`](https://github.com/ddollar/forego) process manager:

```bash
    cd examples/cron-echo
    forego start -e .env -f Procfile

    forego  | starting debug.1 on port 5000
    forego  | starting cron.1 on port 5100
    debug.1 | Listening on [axon.test42.watch]
    cron.1  | Published [axon.test42.watch] : '{"time":1578429144000234322,"type":"heartbeat","meta":{"timePrecision":"ns"}}'
    debug.1 | {"time":1578429144000234322,"type":"heartbeat","meta":{"timePrecision":"ns"}}
    cron.1  | Published [axon.test42.watch] : '{"time":1578429145000163584,"type":"heartbeat","meta":{"timePrecision":"ns"}}'
    debug.1 | {"time":1578429145000163584,"type":"heartbeat","meta":{"timePrecision":"ns"}}
    cron.1  | Published [axon.test42.watch] : '{"time":1578429146000168716,"type":"heartbeat","meta":{"timePrecision":"ns"}}'
    debug.1 | {"time":1578429146000168716,"type":"heartbeat","meta":{"timePrecision":"ns"}}
    ^C      | ctrl-c detected
    forego  | sending SIGTERM to cron.1
    forego  | sending SIGTERM to debug.1
```
