# ADR-0018 — Extended M3U + XMLTV

**Status:** Accepted
**Date:** 2026-08-13

## Context

IPTV requires a simple broadly compatible playlist and EPG path without a heavy framework.

## Decision

Parse Extended M3U and stream-parse XMLTV into SQLite; playback goes through mpv/MediaCore.

## Consequences

Low memory, easy fixtures, tolerant partial EPG. Provider-specific auth may need adapters.

## Alternatives considered

Large in-memory EPG models and unnecessary IPTV libraries were rejected.
