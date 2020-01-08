axon
====

## About

This project holds event-driven agents implemented in go, that are communicating via NATS channels with each other.

Axon is a set of independent components, that can be written in any programming languages,
which has a library to access https://nats.io/.

The components are event driven agents that either consume and/or produce messages through NATS.
These agents use NATS subjects for communicating with each others.

The structure of the messages a given kind of agent is able to consume,
or produces depents on the given agent, as well as its behavior.

The axon package contains a `common` module, that provides generic functions for the agents,
e.g. connecting to the NATS server, etc.

The package also contain a set of predefined agents, such as `axon-cron`, `axon-debug`,
that are compiled and can be executed as standalone applications.

From a given perspective, axon is similar to the Node-RED (https://nodered.org/)
in the meaning that its agents work similarly like the Node-RED components.
There are three fundamental differences:

1. the axon agents' inputs and outputs are NATS subjects, or channels,

2. the agents can be written in any language,

3. the agents can run on different machines and in any number of instances.

This project currently provides only a handful of agents:

- [`axon-cron`](axon-cron/README.md)
- [`axon-debug`](axon-debug/README.md)
- [`axon-influxdb-writer`](axon-influxdb-writer/README.md)

## Install

In case you only want to use the agents, then you can use the pre-compiled binaries
that you find under the `dist/` folder.

If you want to compile from source, then clone the repository, or install it via the `go get` command.

```bash
    go get https://github.com/tombenke/axon-go
```

Then you need to build the common module, and install the agents one-by-one.

To build the agents to several platforms, run the `build.sh` script from the root of the project folder.

## Examples

The [`examples`](examples/) demonstrates how to run complete networks of agents, using foreman.

### The `cron-echo` flow

The [`cron-echo/`](cron-echo/) folder contains the configuration and start script,
that demonstrates the most primitive working example of an axon flow,
that is made of a producer agent, called `axon-cron`, which produces a message every 1 seconds and sends it
into the `axon.test42.watch` subject. Another agent, called `axon-debug` subscribes to the same topic,
and receives the messages sent by the `axon-cron`, then prints out the message to the console.

The figure below shows the flow diagram:

![The `cron-echo` flow diagram](docs/cron-echo-flow-diagram.png)

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

### The `th-sensor` flow

The following figure shows a more advanced flow, that collects thermometer and humidity sensor data measured periodically, then store into a time-series database, that can be visualized.

![The `th-sensor` flow diagram](docs/th_sensor-flow-diagram.png)

This flow is also triggered by the `axon-cron` agent, but there are sensor agents, that do measurements, when the got the trigger message from the `axon-cron`, then forwards the measured results towards the `influxdb-writer` agent, that will store the data into a time-series database, called [`InfluxDB`](https://docs.influxdata.com/influxdb/v1.7/).

This [`axon-sensor-th` component](https://github.com/tombenke/axon-sensor-th) is implemented in C++ using an Arduino library, and runs on an ESP8266 device. 

The collected data is visualized by the [Chronograf](https://docs.influxdata.com/chronograf/v1.7/) dashboard, as you can see on the following figure:

![Chronograf dashboard](docs/chronograf-dashboard.png)

