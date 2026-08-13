# ADR-0012 — Firefox ESR as commercial streaming runtime

**Status:** Accepted
**Date:** 2026-08-13

## Context

Commercial websites require a full supported browser/DRM path; embedding a custom web engine adds maintenance.

## Decision

Keep Firefox ESR as an on-demand kiosk runtime for configured streaming tiles.

## Consequences

Web memory exists only during use. Service support/quality remains best-effort and externally controlled.

## Alternatives considered

Building a custom DRM browser was rejected; WPE remains experimental.
