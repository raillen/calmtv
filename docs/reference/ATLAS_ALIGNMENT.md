# Project Atlas v0.2 Alignment

Framework source used for this reconstruction:
- repository: `raillen/project-atlas-framework`
- framework version: `0.2.0`
- source commit observed during reconstruction: `819a4c26a8af024c381b4abd0d96129616b0fcbf`
- reconstruction date: `2026-08-13`

## Required v0.2 project files

- `docs/ATLAS.md`
- `PROJECT_STATE.md`
- `atlas.json`
- `.ai/orchestration/model-policy.json`
- `.ai/agents/manifest.json`
- `.ai/skills/manifest.json`
- `.ai/recipes/manifest.json`
- `.atlas/history/project-intelligence.json`

All are present.

## Format policy

Canonical human knowledge: Markdown.
Canonical machine contracts: JSON.
Runtime/derived indexing/context: SQLite allowed but gitignored.
New canonical YAML: prohibited. The root `PROJECT_MANIFEST.yaml`, orchestration
YAML files and Goal YAML projections are generated read-only compatibility views
for older Atlas CLI builds; JSON remains the only maintained machine source.

## Documentation taxonomy

This project uses semantic Project Atlas surfaces: product, user, onboarding, architecture, data, api, ui-ux, developer, testing, security, operations, support, release, decisions and reference.

No numbered legacy documentation taxonomy is used.

## Agent adapters

`AGENTS.md` and `CLAUDE.md` are replaceable platform adapters. They point to canonical repository sources rather than duplicating project truth.

The installed legacy CLI's Goal evidence mismatch is recorded in
`.atlas/history/evidence/ATLAS-CLI-compatibility.md`; it does not override the
canonical JSON Goal state.
