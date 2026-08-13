# Performance Budgets

These values are **provisional targets**, not measured facts, until P01 hardware evidence exists.

| Metric | Proposed target |
|---|---:|
| Cold boot → Home | <20 s |
| Shell ready after session start | <2 s |
| Idle PSS | ≤250 MiB |
| Idle hard regression gate | ≤350 MiB |
| Home idle CPU | near system idle |
| D-pad response p95 | <100 ms |
| Heavy foreground apps | 1 by default |

Benchmark resource classes separately for Firefox, mpv, RetroArch and torrent workloads before assigning hard `MemoryHigh/MemoryMax` values.

Performance evidence records hardware, OS/image version, workload, warmup and variance.
