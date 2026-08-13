# Troubleshooting by Symptom

## Home is slow
Check free memory/cgroups, unexpected background processes, artwork/cache behavior and idle redraw rate.

## Video stutters
Run hardware capability diagnostics, lower codec/quality if hardware decode is unavailable, check IO/network and inspect mpv dropped frames.

## Remote arrows do nothing
Verify Linux sees the HID keyboard/gamepad, then InputManager event mapping, then FocusManager state.

## IPTV channel does not open
Check playlist URL/expiry/headers, network reachability and mpv error. EPG mapping is not required for playback.

## Bluetooth device will not pair
Verify BlueZ sees adapter/device; pairing agent/UI errors should expose a user-readable reason.

## Game will not launch
Verify ROM/system identification, selected core availability and RetroArch logs.
