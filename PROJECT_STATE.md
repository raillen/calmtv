# Current Project State

- Project: **Calm TV** (technical package/runtime identifier: `tv-shell`)
- Framework: **Project Atlas 0.2.0**
- Current phase: **P01 — Shell Foundation**
- Current goal: **P01-G01 — Boot to navigable TV Shell (EXECUTING)**
- Context methodology: **Lean Progressive Context (LPC)**
- Last updated: `2026-08-13T17:46:18-03:00`

## Completed foundation

P00-G01 established the documentation and architecture baseline, aligned to Project Atlas v0.2. Canonical project knowledge now uses Markdown; machine-maintained contracts use JSON; no new canonical YAML is permitted.

## Current execution

P01 implementation now includes the GTK3 shell, central input/focus, session
packaging, recovery launcher, AppManager policy and the first shared system/media
boundaries. Host build, FTS5 tests, a real mpv IPC smoke and a windowed GTK smoke
pass. The remaining P01 gates are a clean image, Matchbox runtime inventory and
measured reference hardware/VM evidence. Do not mark P01 complete from host
compilation alone.

The Q4OS Atom notebook has now launched the navigable development shell. A
dual-monitor fullscreen defect was reproduced there and fixed by targeting one
GDK monitor at a time; complete MVP flows and performance measurements remain
pending.

P02, P03 and P04 have been reconciled to `docs/product/SCOPE.md` and are executing
compatible host work while the external image/hardware gates remain pending. Their
partial evidence is recorded under `.atlas/history/evidence/`; no Goal is marked
DONE without its missing target evidence.

## Recovery order

1. `ENTRYPOINT.md` or a generated platform adapter.
2. `atlas.json`.
3. `PROJECT_STATE.md`.
4. `docs/ATLAS.md`.
5. Active Goal under `.ai/goals/`.
6. Only relevant canonical docs/symbols/tests selected by the context strategy.

Do not load the entire repository by default.
