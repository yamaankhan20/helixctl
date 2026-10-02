# ADR 003 — Use Routed Per-Node CIDRs Before an Overlay Network

**Status:** Accepted

## Context

Containers on different Helix worker nodes need to communicate.

Possible approaches include:

```text
manual static routes
Control-Plane-managed L3 routes
VXLAN
BGP
existing CNI
```

## Decision

Use **one container CIDR per node** and let the Helix Route Controller distribute normal Linux L3 routes through Node Agents.

Example:

```text
Node A host: 192.168.50.11
Node A CIDR: 10.10.1.0/24

Node B host: 192.168.50.12
Node B CIDR: 10.10.2.0/24
```

Node A receives:

```text
10.10.2.0/24 via 192.168.50.12
```

## Why

- packet path stays visible
- no overlay encapsulation initially
- teaches routing and return paths
- node subnet ownership is simple
- route distribution can still be dynamic
- Linux tooling can inspect every step

## Why Not Manual Routes

Manual routes do not fit dynamic node membership.

Helixctl should program/remove the routes automatically.

## Why Not VXLAN Initially

VXLAN introduces:

- tunnel endpoints
- VNI management
- encapsulation
- MTU complexity
- harder packet debugging

It remains a future extension.

## Why Not BGP Initially

BGP is valuable but introduces:

- peering
- advertisements
- withdrawals
- convergence
- policy

The Control Plane already knows topology, so direct route distribution is enough for v1.

## Consequences

The management/underlay network must support host reachability and forwarding.

Container CIDRs must not overlap.

Linux IP forwarding and firewall behavior matter.
