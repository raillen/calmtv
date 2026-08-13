# How-to — Add an IPTV Playlist

1. Open **Live TV**.
2. Choose **Add source**.
3. Select a local M3U/M3U8 file or enter an authorized playlist URL.
4. Optionally add an XMLTV EPG source.
5. Review detected channel count/categories.
6. Save.
7. If EPG channel IDs do not match automatically, use the mapping screen to correct them.

Playlist URLs, tokens and credentials are private data. They must not be included in diagnostics unless explicitly redacted/approved.
# Development MVP path

The IPTV screen reads the configured playlist from:

```text
~/.config/tv-shell/iptv/playlist.m3u
```

It accepts Extended M3U metadata (`tvg-id`, `tvg-logo`, `tvg-chno` and
`group-title`) and imports streams through the shared parser. Credentials and
tokens must not be committed or included in diagnostics.
