# Controls Reference

| Action | Default input |
|---|---|
| Navigate | Arrow keys / D-pad |
| Select | Enter / OK |
| Back | Escape / Back |
| Home | Home key or configured Home action |
| Context/Menu | Menu key or configured shortcut |
| Play/Pause | Media Play/Pause |
| Previous/Next | Media previous/next |
| Volume | System volume keys |
| Mute | Mute key |
| IPTV numeric channel | Number keys when channel-entry mode is active |
| Search text | Physical keyboard, phone remote keyboard or on-screen keyboard |
| Pointer | Optional mouse/touchpad |

All device backends map to semantic InputActions. Individual screens must not implement incompatible global shortcuts.
# Current keyboard/HID map

The Shell maps the following semantic actions centrally:

- arrows: `NAV_UP`, `NAV_DOWN`, `NAV_LEFT`, `NAV_RIGHT`;
- Enter/OK: `ACCEPT`;
- Escape/Backspace: `BACK`;
- Home or `H`: `HOME`;
- `M`: `MENU`;
- Space: `PLAY_PAUSE`;
- Page Up/Page Down: `CHANNEL_UP`/`CHANNEL_DOWN`;
- XF86 audio keys: play/pause, volume and mute.
- number and keypad digits: `CHANNEL_0` through `CHANNEL_9` for IPTV
  selection;

Mouse remains complementary. Screens do not own physical-key bindings.
