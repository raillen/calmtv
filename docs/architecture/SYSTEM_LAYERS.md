# System Layers

## OS layer
Debian, stock kernel/firmware, systemd, udev and package management.

## Display/session layer
Xorg and Matchbox. No compositor in V1. Session startup is deterministic and dedicated to the appliance user.

## Shell layer
GTK3 presentation, focus/navigation, application launch/restore, system settings and global overlays.

## Service adapter layer
Small Go interfaces wrapping D-Bus or official CLIs first:
- NetworkManager;
- BlueZ;
- WirePlumber/wpctl;
- XRandR/xrandr;
- UDisks2;
- logind;
- MPRIS/playerctl.

## Media layer
MediaCore plus mpv/FFmpeg, SQLite runtime state and optional source/provider helpers.

## Application layer
NanoTube, Live TV, local media, music/radio/podcasts, files/downloads, games and optional streaming tiles.

## Derived/runtime layer
`.atlas/runtime`, media caches, thumbnails, EPG cache and local databases. None replace canonical Git documentation/configuration.
