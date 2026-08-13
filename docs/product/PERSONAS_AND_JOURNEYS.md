# Personas and Journeys

## Living-room user

Wants a predictable appliance, not Linux. Primary journeys: boot → continue watching, open live TV, play a local movie, start a retro game, change Wi-Fi/audio.

Success: common journeys require no terminal and no mouse.

## Power user

Uses SMB, downloads, optional browser, custom playlists, external storage and game libraries. Wants advanced settings without cluttering normal Home.

Success: Advanced Mode exposes capability without changing the Simple Mode mental model.

## Maintainer / developer

Needs reproducible image builds, hardware diagnostics, logs, cgroup visibility, safe adapters, stable contracts and regression benchmarks.

Success: failures are reproducible and module boundaries let agents implement adapters/tests independently.

## Low-end hardware owner

Needs the interface to remain responsive under memory/CPU pressure.

Success: foreground interactions remain responsive; heavy processes are limited or terminated rather than letting the whole system swap/thrash.
