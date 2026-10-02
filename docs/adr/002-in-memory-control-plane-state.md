# ADR 002 — Use In-Memory Control Plane State in v1

**Status:** Accepted

## Context

The Control Plane needs current state for:

```text
nodes
workloads
desired state
actual state
services
endpoints
routes
```

The first version has one Control Plane.

## Decision

Keep Control Plane state **in memory** with concurrency-safe data structures.

Do not use SQLite, PostgreSQL, etcd, or another database in v1.

## Why

The first version is focused on:

- Linux runtime internals
- scheduling
- Agent coordination
- networking
- reconciliation
- service discovery

Adding durable storage would introduce another subsystem before the core orchestration loop is proven.

In-memory state also makes state transitions explicit and easy to inspect while the design is evolving.

## Consequences

Positive:

- no database dependency
- simple local development
- clear state ownership
- focus remains on orchestration/runtime

Negative:

- desired state is lost when `helixd` restarts
- only actual state can be reconstructed from reconnecting Agents
- no Control Plane HA

## Restart Behavior

After Control Plane restart:

```text
Agents reconnect
    ↓
Nodes register
    ↓
Agents report running containers
    ↓
actual state is reconstructed
```

Historical desired intent is not guaranteed to survive.

## Future Direction

Possible learning path:

```text
in-memory
    ↓
local WAL + snapshots
    ↓
replicated log
    ↓
Raft / HA Control Plane
```

That is intentionally outside v1.
