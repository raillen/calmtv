# AI-Assisted Development Workflow

Project Atlas manifests under `.ai/` are the machine-readable workforce selection. Canonical project decisions remain in Goals/docs/ADRs.

## Preferred provider aliases
- `openai/codex`
- `anthropic/claude`
- `google/gemini`

These are project-level provider-family aliases. Pin exact executable model IDs before a Goal enters `EXECUTING` so model choice is explicit and auditable.

## Workflow
- Explorer maps affected evidence.
- Architect owns boundary/ADR work.
- Implementer changes only locked scope.
- Tester derives verification from acceptance.
- Reviewer independently checks correctness.
- Security/Performance/UX reviewers enter when risk requires.
- Documentation Maintainer applies Documentation Delta.
- Release Verifier owns release evidence.

Use the least workforce necessary; do not run every role for trivial work.

See also [`../reference/GENERATED_ADAPTERS.md`](../reference/GENERATED_ADAPTERS.md).
