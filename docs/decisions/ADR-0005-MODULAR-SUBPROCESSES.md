# ADR-0005 — Modular shell + adapters + subprocesses

**Status:** Accepted
**Date:** 2026-08-13

## Context

Monolithic FFI increases crash coupling; a microservice mesh increases idle/process overhead.

## Decision

Use a modular Go Shell, stable adapter interfaces and isolated on-demand subprocesses for heavy/high-risk components.

## Consequences

Good failure containment and testability with low idle cost. Requires explicit process lifecycle contracts.

## Alternatives considered

One large process and many resident microservices were rejected.
