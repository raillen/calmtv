# InputAction Contract

Input backends emit semantic actions, not screen-specific key codes.

Required actions:
- navigation: up/down/left/right;
- activation: accept/back/home/menu;
- search/text;
- media: play-pause/next/previous/stop;
- volume: up/down/mute;
- TV: channel up/down and numeric entry;
- numeric entry is normalized as `CHANNEL_0` … `CHANNEL_9`;
- pointer/text events as optional capabilities.

Actions contain source/device metadata for diagnostics but screens should normally ignore the physical source.
