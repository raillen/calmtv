# ADR-0004 — Go + GTK3/gotk3 for Shell

**Status:** Accepted
**Date:** 2026-08-13

## Context

The Shell must remain light while being highly amenable to modular agentic development.

## Decision

Implement core Shell/services in Go; use GTK3/gotk3, GtkBuilder and GTK CSS.

## Consequences

Simple concurrency/process/network tooling and native UI without a web engine. CGO/system dependencies remain part of the build.

## Alternatives considered

Rust/GTK, Python/PyGObject and TypeScript/Tauri were considered; Tauri would place WebKit in the core shell.
