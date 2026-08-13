# Technology Stack Reference

| Layer | Baseline |
|---|---|
| OS | Debian 13 amd64 minimal |
| Init/service/resource | systemd + cgroups |
| Display | Xorg |
| WM | Matchbox |
| Compositor | none V1 |
| Core language | Go |
| UI | GTK3 + gotk3 + GtkBuilder + GTK CSS |
| IPC | D-Bus/godbus; Unix sockets where appropriate |
| Network | NetworkManager |
| Bluetooth | BlueZ |
| Audio | PipeWire + WirePlumber/wpctl |
| Display control | XRandR/xrandr |
| Storage | UDisks2/udisksctl |
| Power | logind |
| DB/search | SQLite + FTS5 |
| Media | mpv JSON IPC + FFmpeg |
| Global media control | MPRIS/playerctl |
| IPTV | Extended M3U + XMLTV |
| Torrent | anacrolix/torrent |
| Games | RetroArch + Libretro |
| Network files | SMB/libsmb2 planned |
| Web streaming | Firefox ESR kiosk on demand |
| Local AI | llama.cpp on demand |
| Packaging | debhelper/dh-golang + debootstrap/live-build |
| CI | GitHub Actions + physical self-hosted runners |
