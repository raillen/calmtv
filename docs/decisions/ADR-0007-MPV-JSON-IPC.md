# ADR-0007 — mpv JSON IPC backend in V1

**Status:** Accepted
**Date:** 2026-08-13

## Context

Deep libmpv integration adds callback/FFI complexity before the product contract is proven.

## Decision

Start mpv as a process and use JSON IPC; keep a backend interface that can move to libmpv later.

## Consequences

Faster implementation and crash isolation. X11 embedding/window handling must be designed carefully.

## Alternatives considered

Direct libmpv remains a future optimization if evidence requires it.
