# SQLite Schema Plan

Suggested databases may be physically combined after benchmark; logical ownership remains explicit.

## Core tables
- `settings`
- `app_state`
- `media_items`
- `playback_progress`
- `favorites`
- `downloads`
- `library_sources`
- `search_fts`

## IPTV
- `channels`
- `channel_groups`
- `epg_programs`
- `epg_channel_map`

## Games
- `games`
- `game_systems`
- `game_sessions`

## Organizer
- `organizer_transactions`
- `organizer_actions`

FTS5 indexes titles, aliases and selected metadata only; it is not a filesystem crawler. Schema migrations are versioned and forward-tested.

`internal/state` creates the core tables and an FTS5 virtual table on Debian
builds compiled with the `sqlite_fts5` Go tag. Development hosts without that
SQLite compile option receive a bounded plain-table fallback and cannot claim
FTS5 performance evidence.

The current migration additionally creates `downloads`, `channels`,
`epg_programs`, `games` and `game_sessions`; the Store exposes bounded upsert
operations for these MVP records.
