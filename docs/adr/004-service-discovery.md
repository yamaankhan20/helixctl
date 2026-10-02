# ADR 004 — Use Health-Aware Internal DNS for Service Discovery

**Status:** Accepted

## Context

Container IPs are ephemeral.

A workload can:

- restart
- move to another node
- receive a new IP
- have multiple replicas

Applications should not hardcode container addresses.

## Decision

Implement Helix Service Discovery as:

```text
Service Registry
    ↓
Endpoint Store
    ↓
Health Filter
    ↓
Internal DNS
```

Stable naming format:

```text
<service>.service.cluster
```

Example:

```text
auth.service.cluster
```

## Why

A simple key/value mapping to one IP is not enough because:

- a service can have multiple replicas
- unhealthy endpoints must be excluded
- rescheduling changes addresses
- applications need a stable lookup interface

DNS is a natural application-facing discovery mechanism.

## Endpoint Rule

A container endpoint becomes discoverable only after it is considered healthy.

When it fails:

```text
endpoint unhealthy/removed
    ↓
DNS stops returning it
```

When replacement starts elsewhere:

```text
new endpoint healthy
    ↓
DNS can return new IP
```

## Consequences

Positive:

- service identity is independent of container identity
- multiple replicas are supported
- rescheduling does not require application reconfiguration

Negative:

- DNS becomes another cluster component
- caching/TTL behavior must be considered
- first single DNS instance is not HA

DNS replication/node-local caching can be future work.
