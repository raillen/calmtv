# ADR-0020 — Package-based MVP update model

**Status:** Accepted
**Date:** 2026-08-13

## Context

A custom transactional updater would increase early product scope.

## Decision

Use signed `.deb` packages/repository for MVP; evaluate transactional/A-B or systemd-sysupdate later.

The MVP update experience will have two layers:

1. automatic, lightweight version checking exposed in the Calm TV Settings;
2. controlled package installation through APT/PolicyKit, with user
   confirmation and an explicit restart during the first MVP implementation.

Fully unattended installation is not the default. It requires signed
repository metadata, a tested rollback path and recovery independent of the
normal Shell.

## Current implementation status

- **Implemented:** GitHub Release `.deb` assets, SHA-256 validation, SSH
  bootstrap and manual remote updater.
- **Partial:** package-based helper and PolicyKit boundary exist; the helper
  can install the latest release after the remote bootstrap.
- **Planned:** Settings UI, periodic check, progress state, confirmation,
  health check and automatic rollback path.

## Consequences

Fastest path to a maintainable MVP. Rollback/recovery needs explicit release policy.

## Alternatives considered

Building a custom OTA system in P01 was rejected.
