# ADR 001 — Use gRPC for Control Plane to Node Agent Communication

**Status:** Accepted

## Context

Helixctl needs an internal protocol for commands and status exchange between `helixd` and `helix-agent`.

Operations include:

```text
RegisterNode
Heartbeat
RunContainer
StopContainer
InspectContainer
ListContainers
ApplyRoutes
```

## Decision

Use **gRPC with Protobuf** for the internal Control Plane <-> Node Agent boundary.

We will use two explicit logical gRPC service contracts for v1:

1. `ControlPlaneService` (hosted by `helixd`, called by `helix-agent`)
2. `AgentService` (hosted by `helix-agent`, called by `helixd`)

This approach acknowledges the distinct directionality of node-to-cluster status reporting versus cluster-to-node command execution.

## Why

- RPC semantics match the operations.
- Protobuf gives explicit typed contracts.
- Go support is mature.
- Binary framing is efficient.
- Deadlines/status codes are built into the model.
- Streaming can be introduced later for heartbeat/status flows.
- It provides a clear boundary between domain logic and transport.

## Alternatives Considered

### REST/JSON

Would work and would reduce tooling.

Not selected internally because the project already uses REST externally, while the internal boundary benefits from a typed RPC contract and future streaming.

### Custom TCP Protocol

Would provide lower-level protocol learning but would distract from the project's container/orchestration goals and require solving framing/versioning/error semantics manually.

## Consequences

Positive:

- generated typed client/server code
- clear API versioning
- explicit deadlines
- future streaming path

Negative:

- Protobuf compiler/tooling required
- generated code must be managed
- domain/proto conversion layer required

## Rules

- generated Protobuf types do not become domain models
- every RPC has a deadline where appropriate
- container identity makes lifecycle operations idempotent where practical
