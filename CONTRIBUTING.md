# Contributing to Helixctl

Helixctl is a low-level Go systems project. Contributions should preserve the separation between the Control Plane, Node Agent, Runtime, networking, and service-discovery layers.

---

## Development Principles

The most important rule is:

```text
Control Plane owns intent.
Node Agent owns local coordination.
Runtime owns process isolation.
Network Manager owns connectivity.
Service Discovery owns stable service identity.
```

Do not solve a local package problem by breaking these boundaries.

---

## Before Making a Change

Read the relevant design document:

```text
docs/ARCHITECTURE.md
docs/RUNTIME.md
docs/NETWORKING.md
docs/CONTROL_PLANE.md
docs/API.md
docs/SECURITY.md
```

Major architecture changes should update documentation in the same change.

---

## Repository Rules

### `cmd/`

Keep `main.go` focused on configuration and dependency wiring.

Do not put scheduler/runtime/network business logic in command packages.

### `internal/domain/`

Keep domain objects transport-independent.

Do not import generated Protobuf types or HTTP handlers into domain types.

### `internal/runtime/`

Do not import Control Plane packages.

### `internal/controlplane/`

Do not call Linux runtime syscalls directly.

### `internal/network/`

Only manipulate Helix-owned network resources.

---

## Generated Code

Files under:

```text
gen/
```

are generated.

Do not hand-edit generated Protobuf code.

Change the source `.proto` file and regenerate.

---

## Formatting

Run:

```bash
gofmt -w .
go vet ./...
```

before submitting changes.

---

## Tests

At minimum:

```bash
go test ./...
```

For concurrency-sensitive changes:

```bash
go test -race ./...
```

Low-level Runtime/network changes should include Linux integration validation where applicable.

---

## Test Placement

Package-level deterministic tests:

```text
internal/package/foo_test.go
```

Cross-component:

```text
tests/integration/
```

Full cluster:

```text
tests/e2e/
```

---

## Privileged Tests

Tests that require:

- root
- namespaces
- cgroups
- mounts
- veths
- routes

must be clearly separated from ordinary unit tests.

Do not make a simple `go test ./...` unexpectedly rewrite the host network.

---

## Commit Style

Prefer focused commits.

Examples:

```text
runtime: add PID namespace setup
network: add bridge reconciliation
controlplane: add least-load scoring
discovery: add health-aware endpoint filter
docs: document route ownership
```

---

## Pull Request Scope

A pull request should ideally have one coherent purpose.

If a change modifies:

```text
Runtime + networking + scheduler + docs
```

all at once without a reason, it will be difficult to review.

---

## Error Handling

Preserve low-level context.

Prefer errors that answer:

```text
what operation failed?
for which container/node?
what underlying error occurred?
```

Avoid swallowing syscall/netlink errors.

---

## Logging

Use structured fields.

Useful:

```text
node_id
container_id
service
operation
request_id
error
```

Avoid logging secrets or arbitrary workload environment contents.

---

## Concurrency

Shared state must have explicit ownership.

Do not introduce package-level mutable maps without synchronization.

Run the race detector for changes involving:

- Control Plane state
- IPAM
- endpoint registry
- heartbeat processing
- reconciliation

---

## Runtime Safety

Runtime changes should be tested in disposable Linux VMs.

Take snapshots before testing new mount, namespace, cgroup, or route behavior.

---

## Networking Safety

Network cleanup must only remove Helix-owned:

```text
bridge
veth
netns
routes
```

Do not write code that broadly flushes host routing or firewall state.

---

## Documentation

Update the appropriate file:

| Change | Document |
|---|---|
| System boundary | `docs/ARCHITECTURE.md` |
| Runtime | `docs/RUNTIME.md` |
| Networking | `docs/NETWORKING.md` |
| Control Plane | `docs/CONTROL_PLANE.md` |
| REST/gRPC | `docs/API.md` |
| Test strategy | `docs/TESTING.md` |
| Setup | `docs/DEVELOPMENT.md` |
| Security | `docs/SECURITY.md` |
| Milestones | `docs/ROADMAP.md` |

---

## ADRs

Use an ADR when a decision is important enough that a future contributor may ask:

> Why was this designed this way?

Current ADRs cover:

```text
gRPC for internal node control
in-memory Control Plane state
routed multi-node networking
health-aware internal service discovery
```

---

## PR Checklist

- [ ] change respects component boundaries
- [ ] code is formatted
- [ ] `go test ./...` passes
- [ ] race test run where relevant
- [ ] Linux integration test added where relevant
- [ ] no unrelated generated changes
- [ ] docs updated
- [ ] security implications considered
- [ ] cleanup/error paths tested
- [ ] no unmeasured performance claims added
