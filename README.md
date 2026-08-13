# TV Shell

**TV Shell** is the engineering codename for an ultra-light Linux TV/media appliance designed for low-end x86-64 hardware, especially Atom/Celeron-class machines with 2 GB RAM.

The product is not a conventional desktop environment. It is a remote-first TV shell over a minimal Debian base, with aggressive application lifecycle management, local/network media, IPTV, NanoTube integration, retro emulation, optional commercial streaming runtime and a constrained local smart organizer.

## Status

- Product name: **pending**; `TV Shell` is a codename.
- Project Atlas: **v0.2**.
- Current phase: **P01 — Shell Foundation**.
- Active Goal: **P01-G01 — Boot to navigable TV Shell**.
- Implementation: **development MVP slice in progress; image and hardware gates pending**.
- Canonical documentation: [`docs/ATLAS.md`](docs/ATLAS.md).

## Core target

- Debian 13 minimal, amd64/x86-64.
- Xorg + Matchbox, no compositor in V1.
- Go + GTK3/gotk3 for the shell.
- systemd/D-Bus adapters for system integration.
- SQLite + FTS5 for local runtime state/search.
- mpv via JSON IPC in V1.
- One heavy foreground app at a time by default.
- Processes and models loaded on demand.
- 720p/1080p-first; 4K is not a V1 promise.
- Provisional idle target: ≤250 MiB PSS; hard regression gate proposed at ≤350 MiB until hardware validation.

The current tree includes the Go/GTK3 shell, central keyboard/D-pad focus,
Linux adapters, AppManager policy, mpv/IPTV/state boundaries, local media,
downloads, diagnostics, recovery and on-demand NanoTube/RetroArch/Firefox
runtime launchers. Boot time,
idle PSS, input latency, VA-API and clean-image behavior remain measured
evidence gates.

The Debian package now exposes `TV Shell` as an additional display-manager
session, so it can be tested beside an existing desktop without replacing it.

## Start here

Humans and AI agents should start at [`ENTRYPOINT.md`](ENTRYPOINT.md), then use [`docs/ATLAS.md`](docs/ATLAS.md) as an intent router.

## Content boundary

The project supports user-provided, authorized, public-domain and otherwise lawful media sources. It does not define bypasses for DRM/paywalls or ship piracy-focused indexes.
