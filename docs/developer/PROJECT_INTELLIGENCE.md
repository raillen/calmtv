# Project Intelligence

Durable intelligence lives at `.atlas/history/project-intelligence.json`.

Record compact task facts:
- task ID/status/type/component;
- observed or estimated token/cost data with provenance;
- effort evidence where available;
- tests/gates;
- documentation impact;
- risk/debt;
- evidence pointers;
- model/provider IDs used.

Do not persist hidden reasoning, raw chat history or task-specific context dumps.

Project-level dashboards are derived from this JSON and may be regenerated or replaced.
