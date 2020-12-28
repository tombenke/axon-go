/*
Package actor provides the generic, built-in functions of all kind of axon nodes.

Message Processing Nodes and Network

Withing the axon-go library we can create so called actor nodes.
These actor nodes can communicate with each other via messaging topics,
so the actor nodes form a message processing network.

There are three kind of actor nodes:

* Message Producer Node: can send only messages to other nodes.
* Message Consumer Node: can receive only messages from other nodes.
* Message Processor Node: can both receive, process and send messages.

Messages

Messages are

Ports

Each actor node may have zero, one or more ports, and a processing function.
Producer nodes have only output ports, consumer nodes has only input ports,
and processing nodes have both input and output ports.

The ports are identified by their names, and they have message type assigned to them.
Each port has only one type of message assigned to it.

Topics

The topics through which the nodes communicate with each other also have their names.
The input and output ports needs to be assigned to these messaging topics, but this assignment is optional.
If an output port han no topic assigned, it will not send any message.
If an input port has no topic assigned, then it will not receive any message.
The input ports have a default message value determined either by the message type itself,
or the configuration of that specific port. This default value is used by the processing function in case
there is no incoming message from the pinput port but the computation has to be executed.

Synchronous vs. Asynchronous Mode

The nodes of a message processing network can work in one of two modes:

* Synchronous Mode: There is a central master unit, that controls and harmonizes the cooperation of the nodes.
The nodes synchronizes the receiving, processing and sending of messages with each other, via the contribution of the master.

* Asynchronous Mode: There is synchronized processing.
The nodes immediately processing the incoming messages arrived to the input ports,
then send the results through the output ports as soon as they can.

In theory any node can work both in synchronous or asynchronous mode, but only in one of the mode at a given time.
To change between the modes, the node needs to be restarted. in most of the cases only one mode makes sense to a specific node type.

Processing

All kind of nodes has a processing function.

The producer-only nodes usually has an internal `Next(message)` function that is used to make the node to emit a message.
Before emit, it calls its processor function with the `message` argument provided by the caller.

The implementation steps of an actor node application

1. Define the config structure for the actor node, that includes the `common/config/Node struct`
2. Initialize the actor config with the default values and with the predefined IO ports.
3. Parse the CLI arguments, that returns with the CliConfig struct.
4. Load the config from file, if it exists (Default: ./config.yml), that returns with the FileConfig struct.
5. Merge the config structures into the combined one: Cliconfig -> FileConfig -> Config, if it is enabled.
6. Start the processes according to the config parameters.

*/
package actor
