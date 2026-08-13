# Tooling and Reuse Matrix

The project follows “service/protocol/CLI/library before custom implementation”.

| Area | Preferred reuse | Priority | Notes |
|---|---|---:|---|
| App discovery | GIO / GDesktopAppInfo | P0 | Avoid custom `.desktop` parser |
| D-Bus | godbus/dbus v5 | P0 | Core Go IPC |
| Wi-Fi/Ethernet | NetworkManager D-Bus | P0 | ConnMan only as future benchmark |
| Bluetooth | BlueZ D-Bus | P0 | Discovery only when needed |
| Audio | PipeWire/WirePlumber + `wpctl` | P0 | CLI adapter first |
| Displays | `xrandr`; autorandr optional | P0/P1 | Native RandR only if needed |
| Storage | UDisks2 / `udisksctl` | P0 | Mount/unmount/status |
| Power | systemd-logind | P0 | No sudo from Shell |
| Media | mpv JSON IPC + FFmpeg/ffprobe | P0 | libmpv later if justified |
| Global media | MPRIS / playerctl | P1 | Shared media keys |
| Hardware decode | VA-API / `vainfo` | P0 QA | Gate on real hardware |
| Search/state | SQLite + FTS5 | P0 | No resident search daemon |
| IPTV | small Extended M3U parser + `encoding/xml` | P0/P1 | XMLTV streaming parse |
| Torrent | anacrolix/torrent | P1 | Isolated helper |
| SMB | libsmb2 | P1 | Thin integration/helper |
| Discovery | Avahi/mDNS | P1 | Optional; direct addresses remain |
| DLNA | GUPnP/GUPnP-AV | P2 | Sidecar/helper preferred |
| Games | RetroArch/Libretro | P1 | Shell is frontend |
| Shell gamepad | libmanette | P2 | Only if HID mapping is insufficient |
| HDMI-CEC | libCEC | P2 | Hardware-dependent |
| Podcasts | gofeed | P1 | RSS/Atom parsing only for podcast feeds |
| Commercial web | Firefox ESR kiosk | P1 | On-demand |
| Generic browser | Firefox ESR optional profile | P2 | uBlock, tab cap, discard |
| Local AI | llama.cpp | P2 | On-demand, structured output |
| Image build | debootstrap + live-build | P0 | Debian-native |
| Packaging | debhelper + dh-golang | P0/P1 | Reproducible `.deb` |
| Hardening | systemd + AppArmor | P0/P1 | High-risk helpers first |
| CI | GitHub Actions + self-hosted hardware | P0 | VM cannot prove hardware |
| Go quality | go test, golangci-lint, govulncheck | P0 | Release gates |
| Benchmark | systemd-analyze, systemd-cgtop, stress-ng, Phoronix | P1 | Evidence, not claims |

Avoid adding a dependency merely because it exists. Use it only if it deletes meaningful custom code or maintenance burden.
