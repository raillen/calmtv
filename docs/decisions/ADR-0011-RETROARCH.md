# ADR-0011 — RetroArch/Libretro on demand

**Status:** Accepted
**Date:** 2026-08-13

## Context

Implementing multiple emulators/frontends would be expensive and duplicate mature tooling.

## Decision

TV Shell is the frontend; launch RetroArch/core/ROM on demand and terminate on return.

## Consequences

Near-zero idle emulation cost and broad 8/16-bit support. Core/hardware settings need profiles.

## Alternatives considered

A separate always-running EmulationStation/Pegasus-style frontend was rejected.
