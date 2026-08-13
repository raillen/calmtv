# Tutorial — Navigate With a Remote

The core interface is designed around five actions:

- Up / Down / Left / Right — move focus.
- OK / Enter — activate.
- Back — return one level or close an overlay.
- Home — return to Home.
- Menu — from Home, open the advanced Diagnostics surface.

A mini keyboard/touchpad remote is treated as standard keyboard/mouse HID when the device exposes those interfaces. The touchpad is optional for normal Shell navigation but remains useful for web content or legacy applications.

Media keys map globally through the shared media-control layer. Volume keys change system volume rather than per-app volume unless a screen explicitly says otherwise.

In Diagnostics, **Terminal avançado** opens `qterminal`, `xfce4-terminal`,
`mate-terminal`, `konsole` or `xterm`, whichever is installed. The terminal
runs as the logged-in user and is not required for normal operation.
