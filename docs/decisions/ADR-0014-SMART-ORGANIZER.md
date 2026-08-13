# ADR-0014 — Deterministic-first Smart Organizer

**Status:** Accepted
**Date:** 2026-08-13

## Context

File/ROM organization is largely solvable by rules and databases; a resident AI agent is wasteful and riskier.

## Decision

Run deterministic classification first; use a small llama.cpp model only for ambiguity and only to propose structured actions.

## Consequences

AI is 0 RAM when unused and safer. Ambiguous cases may require confirmation.

## Alternatives considered

Always-on Ollama/chatbot and unrestricted shell agent were rejected.
