# ADR-0015 — AppArmor + systemd hardening

**Status:** Accepted
**Date:** 2026-08-13

## Context

The system processes hostile web/network/media input and an organizer that can mutate files.

## Decision

Use systemd hardening/cgroups plus AppArmor profiles for high-risk processes.

## Consequences

Defense in depth with Debian-native mechanisms. Profiles need maintenance/testing.

## Alternatives considered

SELinux as an additional mandatory layer was deferred; seccomp-alone was rejected.
