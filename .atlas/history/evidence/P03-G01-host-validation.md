# P03-G01 host validation

Status: PARTIAL — shared playback, persistence and parsers are implemented;
external providers and target media fixtures remain pending.

## Passed

- Full Go test suite passes with the `sqlite_fts5` build tag, including a
  migration assertion for the FTS5 virtual table.
- `scripts/test-mpv-ipc` completes a real mpv JSON IPC request.
- Extended M3U and streaming XMLTV parser fixtures pass with bounded input
  behavior.
- Local media scanning, download progress, SQLite state, ROM metadata and the
  NanoTube process boundary have unit coverage.

## Pending

- playback of real local/USB fixtures and dropped-frame measurement;
- a real NanoTube installation/provider;
- IPTV stream/EPG smoke on target hardware;
- end-to-end Shell flows and persisted Continue Watching in a clean image.
