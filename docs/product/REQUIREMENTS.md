# Requirements

## Functional

### Shell and navigation
- Boot directly into a TV-oriented Home.
- Complete daily tasks with D-pad + OK + Back + Home.
- Expose search, recent/continue content, content areas, apps/tools and quick system controls.
- Restore prior application state where practical after resource-driven termination.

### System
- Manage Wi-Fi/Ethernet, Bluetooth, volume/output, display, removable storage, suspend/reboot/poweroff.
- Use Debian/kernel/services for hardware support; project owns UX/policy.
- Detect and surface hardware decode capabilities.
- Check for Calm TV updates automatically and expose controlled installation,
  confirmation and restart from Settings.

### Media
- Local audio/video playback.
- Shared play/pause/seek/tracks/subtitle/language contract.
- IPTV Extended M3U import and XMLTV EPG.
- NanoTube as an integrated first-party app.
- Music/radio/podcasts through the shared media layer.

### Games
- Launch RetroArch cores on demand.
- Organize ROMs deterministically by format/hash/database.
- Return to the Shell cleanly after emulator exit.

### Streaming web
- Commercial streaming tiles may launch isolated Firefox ESR kiosk sessions.
- Generic browser is not required in the default image.
- Optional browser package may use uBlock Origin, a low tab cap and inactive-tab discard.

### Smart Organizer
- Rules-first classification.
- LLM only for ambiguous language/name interpretation.
- Structured plan, preview, allowed directories and undo.
- No arbitrary shell execution.

## Non-functional

- amd64/x86-64 baseline.
- 2 GB RAM minimum target.
- Remote-first and keyboard-accessible.
- Idle work should be event-driven.
- No telemetry by default.
- High-risk processes sandboxed.
- Reproducible package/image pipeline.
- Clean recovery path from shell/update failures.
- Update integrity, interrupted-update recovery and a user-visible update
  status path.
- Canonical docs and JSON contracts validated in CI.
