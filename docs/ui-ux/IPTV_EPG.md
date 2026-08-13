# IPTV / EPG Interaction

The guide uses a conventional time-grid mental model.

- Up/Down: channel.
- Left/Right: program/time.
- OK: tune or open program details depending on state.
- Number keys: direct channel entry when enabled.
- Channel Up/Down: immediate channel change.
- Back: close details/guide hierarchy.

The EPG must remain usable with partial/missing data. Channel playback cannot depend on a perfect `tvg-id` match.

Large XMLTV payloads are not retained in UI memory; UI queries normalized SQLite ranges.
