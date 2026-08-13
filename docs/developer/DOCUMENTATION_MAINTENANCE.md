# Documentation Maintenance

Before editing docs, compute Documentation Delta:

```text
User:
Developer:
Architecture/API:
Testing:
Operations/support:
Migration/release:
ATLAS links:
```

Patch before proliferation. Create a new document only for a stable concept, distinct audience/task or independent specification.

Rules:
- Markdown for human knowledge.
- JSON for machine contracts.
- No new YAML canonical state.
- Generated docs site is never source of truth.
- Do not duplicate the same explanation in user/dev/agent surfaces.
- Broken local links are a CI failure.
