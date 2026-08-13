# ADR-0019 — Provider/torrent isolation and content boundary

**Status:** Accepted
**Date:** 2026-08-13

## Context

Torrent streaming/providers add network/file/security/legal concerns.

## Decision

Use isolated provider/torrent helpers; official distribution only supports lawful/public/authorized sources and user-provided inputs.

## Consequences

Keeps Shell safe and distribution policy explicit. Some source ecosystems remain intentionally unsupported.

## Alternatives considered

Bundling piracy-focused indexes or executing provider code in-process was rejected.
