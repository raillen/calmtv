# Data Model

SQLite is runtime/user state, not canonical engineering truth.

Primary logical entities:
- `AppState`
- `MediaItem`
- `PlaybackProgress`
- `Favorite`
- `Download`
- `LibrarySource`
- `Channel`
- `EpgProgram`
- `Game`
- `Provider`
- `Setting`
- `OrganizerTransaction`

Stable IDs should survive path moves where practical. File paths are attributes, not universal identities.

External metadata IDs may be stored alongside local IDs but must not become mandatory for local playback.
