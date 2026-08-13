# IPTV and EPG Architecture

## Import

Extended M3U is parsed into normalized channels/groups. XMLTV is decompressed/stream-parsed incrementally, normalized and inserted into SQLite; large XML trees are not held in RAM.

`internal/iptv` implements this import contract with a bounded scanner for
Extended M3U and `xml.Decoder` callbacks for XMLTV. Persistence is owned by
the caller, so imports can write channels/programs incrementally to SQLite.

## Playback

Channel selection produces a MediaCore descriptor and mpv performs HLS/HTTP playback. The EPG is advisory metadata: a channel can still play when EPG mapping is absent.

## Compatibility

Adapters may add required headers/user-agent/token refresh semantics without polluting the core channel model. DRM bypass is out of scope.

## Refresh

Playlist/EPG refresh occurs on explicit action or bounded schedule; no high-frequency polling at idle.
