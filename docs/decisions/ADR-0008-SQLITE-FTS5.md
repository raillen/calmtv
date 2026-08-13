# ADR-0008 — SQLite + FTS5 for runtime state/search

**Status:** Accepted
**Date:** 2026-08-13

## Context

The product needs local state/search without heavy resident indexers.

## Decision

Use SQLite/FTS5 for user/runtime data and local search; do not use it as canonical engineering truth.

## Consequences

Small footprint, transactional state and rebuildable search. Requires migrations.

## Alternatives considered

Tracker/Baloo-style resident indexing and external search servers were rejected.
