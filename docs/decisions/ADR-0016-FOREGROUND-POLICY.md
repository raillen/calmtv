# ADR-0016 — One heavy foreground app by default

**Status:** Accepted
**Date:** 2026-08-13

## Context

Desktop-style simultaneous heavy apps would create memory pressure and HDD swapping.

## Decision

Default to one heavy foreground app, bounded background audio/download work and restorable cached state.

## Consequences

Predictable low-end responsiveness. Some users may perceive app relaunch delay.

## Alternatives considered

Unlimited multitasking was rejected for the default performance profile.
