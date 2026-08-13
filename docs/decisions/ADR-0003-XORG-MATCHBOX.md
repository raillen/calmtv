# ADR-0003 — Xorg + Matchbox, no compositor in V1

**Status:** Accepted
**Date:** 2026-08-13

## Context

The product targets old hardware and needs a TV shell rather than modern desktop effects.

## Decision

Use Xorg and Matchbox; do not run a compositor in V1.

## Consequences

Low idle cost and straightforward X11 embedding/input. Tearing/display behavior must be tested and handled with driver/player settings rather than assuming a compositor.

## Alternatives considered

Wayland and X11+Picom were deferred.
