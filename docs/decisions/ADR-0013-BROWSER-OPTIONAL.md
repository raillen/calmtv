# ADR-0013 — Generic browser is optional

**Status:** Accepted
**Date:** 2026-08-13

## Context

A general browser is likely the heaviest and least predictable workload on 2 GB.

## Decision

Do not include a generic browser destination in the default product; package an optional controlled profile if desired.

## Consequences

Smaller default UX/footprint. Users needing general web access can install the optional component.

## Alternatives considered

Permanent general browser integration was rejected.
