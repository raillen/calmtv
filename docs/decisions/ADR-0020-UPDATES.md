# ADR-0020 — Package-based MVP update model

**Status:** Accepted
**Date:** 2026-08-13

## Context

A custom transactional updater would increase early product scope.

## Decision

Use signed `.deb` packages/repository for MVP; evaluate transactional/A-B or systemd-sysupdate later.

## Consequences

Fastest path to a maintainable MVP. Rollback/recovery needs explicit release policy.

## Alternatives considered

Building a custom OTA system in P01 was rejected.
