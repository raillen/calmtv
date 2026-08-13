# TV Shell — Project ATLAS

This file is an **intent router**, not a request to read the documentation tree. Start with the task you are trying to accomplish and load only the linked material you need.

## Current state / active work

- [Current project state](../PROJECT_STATE.md)
- `atlas.json` — canonical machine configuration
- `.ai/goals/P01/P01-G01.goal.json` — active locked Goal

## I want to understand the product

- [Product vision and principles](product/PRODUCT.md)
- [Scope and non-goals](product/SCOPE.md)
- [Requirements](product/REQUIREMENTS.md)
- [Roadmap](product/ROADMAP.md)
- [Content/legal boundary](product/CONTENT_POLICY.md)

## I want to use the product

- [User documentation map](user/ATLAS.md)
- [Getting started tutorial](user/tutorials/GETTING_STARTED.md)
- [Controls reference](user/reference/CONTROLS.md)
- [Troubleshooting](support/TROUBLESHOOTING.md)

## I want to set up development

- [Development setup](onboarding/DEVELOPMENT_SETUP.md)
- [First contribution](onboarding/FIRST_CONTRIBUTION.md)
- [Codebase tour](developer/CODEBASE_TOUR.md)
- [Build and run](developer/BUILD_AND_RUN.md)
- [Tooling/reuse matrix](reference/TOOLING_AND_REUSE.md)

## I want to understand or change the architecture

- [Architecture map](architecture/ATLAS.md)
- [System architecture](architecture/ARCHITECTURE.md)
- [App lifecycle and resource control](architecture/APP_MANAGER.md)
- [Media architecture](architecture/MEDIA_CORE.md)
- [API/contracts map](api/ATLAS.md)
- [Accepted decisions](decisions/ATLAS.md)

## I want to work on UI/UX

- [UI/UX map](ui-ux/ATLAS.md)
- [Design system](ui-ux/DESIGN_SYSTEM.md)
- [Focus and remote navigation](ui-ux/FOCUS_AND_NAVIGATION.md)
- [Screen specifications](ui-ux/SCREEN_SPECS.md)
- [Wireframes](ui-ux/WIREFRAMES.md)

## I want to add or test a feature

- [How to add an app](developer/HOW_TO_ADD_APP.md)
- [How to add a system adapter](developer/HOW_TO_ADD_SYSTEM_ADAPTER.md)
- [How to add a media provider](developer/HOW_TO_ADD_MEDIA_PROVIDER.md)
- [Test strategy](testing/TEST_STRATEGY.md)
- [Performance budgets](testing/PERFORMANCE_BUDGETS.md)

## I want to operate/support/release it

- [Operations map](operations/ATLAS.md)
- [Configuration](operations/CONFIGURATION.md)
- [Recovery](operations/RECOVERY.md)
- [Diagnostics/support](support/DIAGNOSTICS.md)
- [Release process](release/RELEASE_PROCESS.md)

## I am an AI agent

1. Read the active Goal.
2. Apply Lean Progressive Context and load the smallest sufficient document/symbol/test slices.
3. Respect ADRs and Goal acceptance criteria.
4. Keep delegation bounded.
5. Compute a Documentation Delta before editing docs.
6. Record evidence/Project Intelligence; do not persist task-specific context summaries.

Project-specific workforce and routing are described in [AI workflow](developer/AI_WORKFLOW.md). Platform adapters such as `AGENTS.md` and `CLAUDE.md` are generated/replaceable views, not canonical knowledge.
