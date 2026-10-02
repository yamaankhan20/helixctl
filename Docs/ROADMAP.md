# Helixctl Roadmap

> Implementation order for building the complete Runtime + Orchestrator + Networking system without hiding important Linux mechanics behind existing platforms.

---

## 1. Roadmap Principle

The project should grow vertically.

Each milestone should produce something observable before another major subsystem is stacked on top.

```mermaid
flowchart LR
    F[Foundation]
    P[Process Runner]
    NS[Namespaces]
    CG[cgroups v2]
    RF[RootFS / pivot_root]
    LN[Local Networking]
    AG[Node Agent]
    CP[Control Plane]
    SC[Scheduler]
    E2E[End-to-End Run]
    MR[Multi-Node Routing]
    HM[Health]
    RC[Reconciliation]
    RS[Rescheduling]
    SD[Service Discovery]
    DNS[Internal DNS]
    FT[Failure Testing]
    OBS[Observability / Polish]

    F --> P --> NS --> CG --> RF --> LN --> AG --> CP --> SC --> E2E
    E2E --> MR --> HM --> RC --> RS --> SD --> DNS --> FT --> OBS
```

The sequence is:

---

# 2. Milestone 0 — Repository Foundation

Build:

- Go module
- `cmd/helixctl`
- `cmd/helixd`
- `cmd/helix-agent`
- `internal/domain`
- `internal/config`
- structured logging
- Makefile
- docs structure

Success:

```text
all three binaries build
go test ./... passes
```

---

# 3. Milestone 0A — Architecture Contracts / Boundaries

Finalize core design contracts to prevent major rewrites:

- Workload / Container Instance / Node / Service / Endpoint identity
- Runtime vs Network Manager boundary
- Agent ↔ Control Plane transport topology
- Node registration model
- Container lifecycle states
- Node health states
- Route generation semantics
- Idempotent request identity
- v1 health semantics

Success:

```text
Core gRPC contracts and domain boundaries finalized.
```

---

# 4. Milestone 1 — Local Process Runner

Build a minimal local Runtime that can execute a process through Go.

Learn:

- `os/exec`
- process lifecycle
- PID tracking
- exit status
- signals

Success:

```text
Helix Runtime can start, inspect, stop, and wait for a local process
```

No namespaces yet.

---

# 4. Milestone 2 — Linux Namespaces

Add:

```text
UTS
PID
Mount
Network
```

Start with UTS/PID, then mount/network.

Verify with:

```text
hostname
ps
lsns
/proc/<pid>/ns
```

Success:

```text
container process has observable namespace isolation
```

---

# 5. Milestone 3 — cgroups v2

Add:

- cgroup hierarchy
- memory limit
- CPU limit
- PID attachment
- cleanup

Success:

```text
container resource limits are visible in cgroup files and enforced
```

---

# 6. Milestone 4 — Root Filesystem

Start:

```text
BusyBox rootfs
```

Progress:

```text
chroot proof
    ↓
private mount namespace
    ↓
pivot_root
    ↓
/proc mount
```

Success:

```text
container sees the intended isolated root filesystem
```

---

# 7. Milestone 5 — Single-Node Networking

Build:

- `helix0`
- network namespaces
- veth pairs
- IPAM
- container IP
- default route

Success:

```text
two containers on one node communicate through helix0
```

---

# 8. Milestone 6 — Node Agent

Build:

- gRPC server
- node registration client
- heartbeat loop
- local resource reporting
- Runtime integration
- Network Manager integration
- container monitor

Success:

```text
remote caller can ask Agent to run/inspect/stop a container
```

---

# 9. Milestone 7 — Control Plane Skeleton

Build:

- REST server
- in-memory cluster state
- Node Registry
- AgentClient
- node list API
- container request model

Success:

```text
Agents register and helixctl can list nodes
```

---

# 10. Milestone 8 — Scheduler

Implement:

```text
resource filters
round-robin baseline
least-load scoring
```

Success:

```text
container request is assigned to a healthy node with sufficient capacity
```

---

# 11. Milestone 9 — End-to-End Container Run

Connect:

```text
helixctl
    ↓
REST
    ↓
helixd
    ↓
scheduler
    ↓
gRPC
    ↓
helix-agent
    ↓
Runtime + Network
```

Success:

```text
helixctl container run creates a real isolated process on the selected worker
```

This is the first major vertical slice.

---

# 12. Milestone 10 — Multi-Node Routing

Add:

- management IPs
- per-node container CIDRs
- Route Controller
- `ApplyRoutes()` RPC
- route reconciliation
- IP forwarding checks

Success:

```text
container on Node A can reach container on Node B and back
```

---

# 13. Milestone 11 — Health Manager

Add:

- heartbeat timestamps
- HEALTHY
- SUSPECT
- UNREACHABLE
- network readiness

Success:

```text
Control Plane identifies an unavailable worker using heartbeat policy
```

---

# 14. Milestone 12 — Desired / Actual State

Separate:

```text
Desired State
Actual State
```

Add the Reconciler.

Success:

```text
state mismatch produces a deterministic corrective action
```

---

# 15. Milestone 13 — Workload Rescheduling

When a node becomes unavailable:

```text
endpoint becomes unavailable
workload actual state becomes UNKNOWN
scheduler selects new node
replacement container starts
```

Success:

```text
a RUNNING desired workload is recreated on another healthy node
```

---

# 16. Milestone 14 — Service Registry

Add:

- Service model
- Endpoint model
- service membership
- health-aware endpoints

Success:

```text
one logical service can track multiple container endpoints
```

---

# 17. Milestone 15 — Internal DNS

Add:

```text
<service>.service.cluster
```

DNS should return healthy endpoints only.

Success:

```text
a container can resolve another Helix service without hardcoding its IP
```

---

# 18. Milestone 16 — Rescheduling + Discovery

Combine node failure with service discovery.

Success:

```text
old endpoint disappears
replacement starts elsewhere
new endpoint becomes healthy
same service name resolves successfully
```

This proves that service identity is independent of container IP.

---

# 19. Milestone 17 — Control Plane Restart Reconstruction

Because state is intentionally in memory:

- restart `helixd`
- keep Agents/containers running
- Agents reconnect
- list actual running workloads
- rebuild current actual state

Success:

```text
the system accurately documents what it can reconstruct and what desired intent was lost
```

---

# 20. Milestone 18 — Observability

Add:

- structured logs
- request IDs
- runtime metrics
- scheduler metrics
- network metrics
- service discovery metrics

Success:

```text
one workload can be traced from API request to Agent/Runtime outcome
```

---

# 21. Milestone 19 — Test Hardening

Required:

```text
go test ./...
go test -race ./...
```

Add:

- runtime integration tests
- same-node networking test
- cross-node networking test
- route join/remove
- node failure
- DNS endpoint filtering
- cleanup leak checks

---

# 22. Milestone 20 — Repository Polish

Complete:

- README
- Architecture docs
- API docs
- diagrams
- demo commands
- benchmark/test results where measured
- screenshots/log snippets only where useful

Do not claim performance numbers that were not measured.

---

# 23. Future — OCI Image Support

Possible future work:

- OCI image layout
- layer extraction
- registry pull
- digest verification
- image cache

This is not needed for the first Runtime.

---

# 24. Future — Stronger Runtime Security

Add:

```text
USER namespace
capability dropping
seccomp
AppArmor/SELinux experiments
rootless execution
```

---

# 25. Future — VXLAN

After routed networking works:

```text
bridge
    ↓
VXLAN interface
    ↓
underlay
    ↓
remote VXLAN
    ↓
bridge
```

Compare packet path, MTU, and debugging complexity with direct routing.

---

# 26. Future — BGP

Replace/control route distribution with real route advertisement.

Explore:

- peering
- announcement
- withdrawal
- convergence
- routing policy

---

# 27. Future — eBPF

Explore eBPF for:

- networking dataplane
- observability
- policy
- service load balancing

Only after the traditional Linux path is understood.

---

# 28. Future — HA Control Plane

The current single Control Plane + in-memory state does not provide HA.

A future version can explore:

```text
replicated state machine
Raft
leader election
quorum
durable replicated log
```

This should be treated as a separate distributed-systems milestone rather than silently added to v1.

---

# 29. Future — Persistent Control State

If the project later requires desired state to survive Control Plane restarts, persistence can be added behind state/domain interfaces.

Possible learning path:

```text
in-memory
    ↓
local WAL + snapshot
    ↓
replicated log
    ↓
Raft / HA
```

No database is required for the current first version.

---

# 30. Future — Volumes

Persistent workload storage introduces:

- mount lifecycle
- node affinity
- remote storage
- failure recovery
- data consistency

It is deliberately outside the first container/orchestrator networking scope.

---

# 31. Future — Network Policy

Possible policy:

```text
frontend -> backend ALLOW
frontend -> database DENY
```

Can be explored with:

```text
nftables/iptables
later eBPF
```

---

# 32. v1 Completion Criteria

The first complete Helixctl version is reached when:

- [ ] custom Runtime works
- [ ] namespaces work
- [ ] cgroups v2 limits work
- [ ] `pivot_root` works
- [ ] local container networking works
- [ ] Agents register
- [ ] heartbeats work
- [ ] Control Plane schedules workloads
- [ ] multi-node routing works
- [ ] desired/actual reconciliation works
- [ ] node failure triggers replacement
- [ ] service discovery works
- [ ] internal DNS returns healthy endpoints
- [ ] route/service state changes correctly after failure
- [ ] race tests pass
- [ ] E2E demo is reproducible
