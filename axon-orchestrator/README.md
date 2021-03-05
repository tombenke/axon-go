# axon-orchestrator

The `axon-orchestrator` plays a central role in an axon event processing network:
- It acts as a process manager (start, stop, restart, etc. the actors that belong to a specific event processing network.
- It sends heart-beat messages and observes the presence of the actor notes based on their responses.
- It synchronizes the communication of those actors that are working in `orchestration.synchronization` mode.


