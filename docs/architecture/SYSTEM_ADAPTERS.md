# System Adapters

Project code should prefer official protocol/CLI boundaries before deep FFI.

| Capability | V1 backend |
|---|---|
| Network | NetworkManager D-Bus via godbus |
| Bluetooth | BlueZ D-Bus |
| Audio | `wpctl` / WirePlumber; native API only if needed |
| Display | `xrandr`; native RandR only if needed |
| Storage | `udisksctl` then UDisks2 D-Bus |
| Power | systemd-logind D-Bus |
| Battery | UPower when applicable |
| Media keys | MPRIS/playerctl |
| Hardware decode diagnostics | `vainfo` + mpv/FFmpeg evidence |

Each adapter converts backend-specific failures to stable domain errors and exposes a mock/fake interface for tests.
