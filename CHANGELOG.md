# Changelog

All notable project changes should be recorded here at release boundaries. Task-level implementation detail belongs in Git history and Project Intelligence rather than this file.

## Unreleased

### Changed
- Renamed the user-facing product to Calm TV while retaining `tv-shell` technical identifiers for package upgrades.
- Fullscreen now targets one monitor, defaulting to the primary output, instead of spanning the X11 virtual desktop.

### Added
- Project Atlas v0.2 project scaffold and canonical documentation baseline.
- Architecture, UI/UX, security, operations, testing, support and release specifications.
- Goal plan from foundation through release hardening.

### Decisions
- Debian 13 minimal / amd64 baseline.
- Xorg + Matchbox V1 shell.
- Go + GTK3/gotk3.
- On-demand subprocess architecture and systemd/cgroup lifecycle control.
- Remote-first Calm TV UI.
