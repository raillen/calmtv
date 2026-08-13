# Provider API

Providers are isolated integrations.

Minimal resource families:
- `catalog`
- `search`
- `metadata`
- `streams`
- `subtitles`

A stream descriptor may include URL or lawful user-provided torrent info, file selection hints, language/quality metadata and required non-secret headers.

Provider responses are untrusted input. The consumer validates size, schemes, paths and expected fields before use.

No provider API operation exposes arbitrary shell execution.
