# ADR-0010 — MPRIS for global media controls

**Status:** Accepted
**Date:** 2026-08-13

## Context

Media keys/Quick Settings should not know each player implementation.

## Decision

Use MPRIS as the external/global player control surface.

## Consequences

Consistent media controls across first-party/external compatible apps.

## Alternatives considered

Per-app custom media-control IPC was rejected.
