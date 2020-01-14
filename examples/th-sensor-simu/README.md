The `th-sensor-simu` flow
=========================

The following figure shows a flow, that collects thermometer and humidity sensor data measured periodically, then store them into a time-series database, that can be visualized.
This flow is very similar to the `th-sensor` flow, but it does not need a hardware component to be installed. Instead, it uses the [`axon-function-js`](../../axon-function-js/) component to simulate the working of a physical device, and generate random temperature and humidity values in the same format as the real physical device would respond.

![The `th-sensor-simu` flow diagram](../../docs/th_sensor-simu-flow-diagram.png)

This flow is triggered by the `axon-cron` agent, and there is a simulated sensor agent, that do measurements, when they got the trigger message from the `axon-cron`, then forwards the measured results towards the `influxdb-writer` agent, that will store the data into a time-series database, called [`InfluxDB`](https://docs.influxdata.com/influxdb/v1.7/).

The collected data is visualized by the [Chronograf](https://docs.influxdata.com/chronograf/v1.7/) dashboard, as you can see on the following figure:

![Chronograf dashboard](../../docs/chronograf-dashboard.png)

