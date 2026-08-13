# ADR-0006 — AppManager + systemd/cgroups

**Status:** Accepted
**Date:** 2026-08-13

## Context

2 GB hardware cannot tolerate unrestricted desktop-style multitasking.

## Decision

AppManager owns process lifecycle and uses systemd transient units/scopes/cgroups for resource policy.

## Consequences

Resource behavior is explicit/testable. Policies need hardware benchmark tuning.

## Alternatives considered

Ad-hoc process spawning and relying solely on kernel OOM were rejected.
