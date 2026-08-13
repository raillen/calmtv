# AppManager

AppManager is the core lifecycle authority.

## Responsibilities

- Launch/stop first-party and external apps.
- Decide foreground/background eligibility.
- Save/restore app state.
- Create systemd transient scopes/cgroups.
- Apply MemoryHigh/MemoryMax, CPUWeight and IOWeight policies.
- Track crash/exit state.
- Request graceful termination before forced termination.
- Coordinate media/background exceptions.

## Default policy

- Shell: always resident.
- Heavy foreground app: maximum 1.
- Background audio: at most 1 source.
- Background downloads: bounded count and low priority.
- Browser/RetroArch: exclusive-heavy by default.
- Cached recent apps: state records, not necessarily live processes.

## Memory pressure policy

1. Discard optional caches.
2. Reduce/pause low-priority background work.
3. Terminate cached/background apps.
4. Ask cooperative foreground app to save state if pressure remains critical.
5. Rely on hard cgroup/OOM limits only as the final boundary.

`SIGSTOP` is not considered a memory-saving strategy because it preserves RAM.

`internal/appmanager` contains the manifest/lifecycle boundary and a
`SystemdStarter` that applies `MemoryHigh`, `MemoryMax`, `CPUWeight` and
`IOWeight` to user scopes. `Background` requires an explicit manifest opt-in;
`SetResourcePolicy` allows a validated per-app override. Unit tests exercise
the one-exclusive-heavy-app rule without starting systemd.
