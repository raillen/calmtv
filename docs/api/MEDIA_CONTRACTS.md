# Media Contracts

## Player
- open media descriptor;
- play/pause/stop;
- seek absolute/relative;
- query duration/position/state;
- enumerate/select audio tracks;
- enumerate/select subtitle tracks;
- volume/mute;
- observe errors/end-of-file.

## Media item
Stable local ID, title, type, source, artwork refs, duration, metadata refs and optional external IDs.

## Playback progress
Item ID, position, duration, completion state and updated time.

## Download
Source descriptor, destination policy, progress, speed/state and retention mode.

MPRIS is the external/global control surface; it does not replace internal domain contracts.
