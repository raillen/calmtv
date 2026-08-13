# MediaCore

MediaCore is a shared code/service boundary, not a permanently large daemon.

## Responsibilities

- Player control abstraction.
- Current item/session state.
- History and Continue Watching.
- Favorites.
- Audio/subtitle preferences.
- Download coordination.
- Search registration.
- MPRIS integration.
- Local library metadata.

## V1 playback backend

Run `mpv` as a subprocess and control it through its JSON IPC Unix socket. This isolates crashes and avoids early libmpv/cgo callback complexity.

`internal/media` implements the shared client/runtime boundary. It supports
open, play, pause, stop, seek, volume, tracks, position, duration and explicit
transport errors. The runtime starts mpv only on demand and closes the IPC
transport/process when stopped.

## Process rule

No player process exists at idle. A playback request starts mpv; leaving playback terminates it unless a permitted background audio session remains.

## Shared consumers

Local media, NanoTube, IPTV, music, radio, podcasts and provider-based movie/series playback use the same semantic player contract.
