# ADR-0001 — Debian 13 minimal base

**Status:** Accepted
**Date:** 2026-08-13

## Context

The target needs mature hardware support, firmware, packages, systemd and a reproducible path that remains manageable for solo agentic development.

## Decision

Use a minimal Debian 13 amd64 base and remove conventional desktop/application bundles. Use Debian stock kernel/firmware.

## Consequences

We inherit Debian hardware/package maintenance and avoid becoming a distro-driver team. The image is not as tiny as a purpose-built Buildroot image, but development/support cost is much lower.

## Alternatives considered

Buildroot/Alpine/custom embedded image were deferred until product behavior and hardware matrix are mature.
