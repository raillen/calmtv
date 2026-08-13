# InputManager

InputManager normalizes all user input into semantic actions.

## Core actions

`NAV_UP`, `NAV_DOWN`, `NAV_LEFT`, `NAV_RIGHT`, `ACCEPT`, `BACK`, `HOME`, `MENU`, `SEARCH`, `PLAY_PAUSE`, `NEXT`, `PREVIOUS`, `VOL_UP`, `VOL_DOWN`, `MUTE`, `CHANNEL_UP`, `CHANNEL_DOWN`, numeric entry and text input.

## Backends

- GTK keyboard events / USB HID remote.
- Mouse/touchpad as optional pointer.
- Bluetooth HID through normal Linux input path.
- Gamepad Shell navigation when enabled.
- HDMI-CEC helper when enabled.
- Phone remote/WebSocket when enabled.

## Rule

Screens consume actions; they do not bind global devices directly. RetroArch handles controller mapping inside games, while Shell navigation remains separate.

## Focus

InputManager requests navigation from the central FocusManager. No screen is allowed to implement a private competing focus graph for standard controls.

The current keyboard/HID-compatible mapping is implemented in
`internal/input/action.go`; GTK events are converted before they reach the
screen stack. The central `internal/focus.Manager` replaces targets when a
screen changes and keeps edge navigation deterministic.
