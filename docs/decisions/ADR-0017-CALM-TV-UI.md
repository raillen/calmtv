# ADR-0017 — Calm TV UI and central focus

**Status:** Accepted
**Date:** 2026-08-13

## Context

A 10-foot remote UI needs much stronger focus/hierarchy than a conventional desktop.

## Decision

Use a restrained dark content-first design and a single central FocusManager for D-pad navigation.

## Consequences

Consistent navigation and easier automated testing; screens cannot invent special focus behavior.

## Alternatives considered

Desktop-style pointer UI and per-screen focus graphs were rejected.
