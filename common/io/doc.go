/*
Package io provides the container structures and handler functions for input and output ports of the actor nodes.

An IO port has the following features:

* Uniquely identifies the port by its `Name`.

* Determines the `Type` of the messages it receives or emits.

* Determines the default value of the message if it is an input port.

* Determines the `Representation` format the input ports decodes from and the output ports encodes to the internal representation of the messages.

* Determines the name of the `Channel` the input port receives, and the output ports sends the messages.

Defining the ports

The actor node has to define its nodes before start using them.

The definition process has the following steps:

1. The actor node implementation defines an initial array of port descriptors.

2. The generic actor configuration module reads the configuration file if it is defined,
and overwrites and/or extends the initial port descriptors with the parameters.

3. The generic actor configuration module also parses the CLI parameters,
and overwrite and/or extends the port descriptors with these additional CLI parameters.

4. The actor uses the IO descriptor configuration to build the `Inputs` and `Outputs` structures.

The configuration of the output ports

An output port has the following configuration parameters:

* `name`: A string value. Mandatory.

* `type`: A valid message-type string, Optional. Default value: `base/Any`.

* `representation`: A valid `msgs.Representation` format. Optional. Default value `application/json`.

* `channel`: A string value. Optional. Default value: "". The default empty string value means, it is a `nil-channel`, so the messages must not be forwarded.

Examples for outputs port configuration:

    outputs:
      - name: water-output
        type: base/Float64
        representation: application/json
        channel: "" # nil, there is no consumer of the output
      - name: water-input-need
        type: base/Float64
        representation: application/json
        channel: high-pressure-wss-input-need

The configuration of the input ports

An input port has the following configuration parameters:

* `name`: A string value. Mandatory.

* `type`: A valid message-type string, Optional. Default value: `base/Any`.

* `representation`: A valid `msgs.Representation` format. Optional. Default value `application/json`.

* `channel`: A string value. Optional. Default value: "". The default empty string value means, it is a `nil-channel`, so there is no messages will be received from any channels.
In case of the input channel the "" empty channel name means that it is a parameter input with a default value.

* `default`: A JSON-format string value. Optional. Default value: "". If the `default` parameter contains a valid JSON string, it will be the default value of the input port.
If it is the "" empty string, then the port will use the default message object that belongs to the message type identified by the `type` parameter.
The default value is used instead of the channel value when either there is no channel defined, or the orchestrator commands the input receiver to forward the inputs to the processor, but there was no input message received yet via the channel.

Examples for inputs port configuration:

    inputs:
      - name: buffer-volume
        type: base/Float64
        representation: application/json
        channel: "" # No channel, this is a parameter
        default: '{"Body": { "Data": 0.05 }}' # m3
      - name: max-pressure
        type: base/Float64
        representation: application/json
        channel: "" # No channel, this is a parameter
        default: '{"Body": { "Data": 2. }}' # bar
      - name: min-pressure
        type: base/Float64
        representation: application/json
        channel: "" # No channel, this is a parameter
        default: '{"Body": { "Data": 1. }}' # bar
      - name: water-output-need
        type: base/Float64
        representation: application/json
        channel: "" # nil, there is no consumer, so no consumption need
        default: '{"Body": { "Data": 0. }}' # m3, There is no high-pressure water consumption by default
      - name: water input
        type: base/Float64
        representation: application/json
        channel: well-water-buffer-tank-water-output
        default: "" # Use the default value defined to the message-type
      - name: water-buffer-tank-level
        type: base/Float64
        representation: application/json
        channel: well-water-buffer-tank-level
        default: "" # Use the default value defined to the message-type

*/
package io
