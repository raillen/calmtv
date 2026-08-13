# ADR-0009 — System-service adapters

**Status:** Accepted
**Date:** 2026-08-13

## Context

Linux already has mature services for network, Bluetooth, audio, display, storage and power.

## Decision

Wrap official D-Bus APIs or small CLIs behind project interfaces.

## Consequences

Reduces code and lets backends evolve. Backend-specific errors must be translated.

## Alternatives considered

Reimplementing system services or adopting deep third-party wrappers as project contracts was rejected.
