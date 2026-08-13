# System Architecture

## Top-level model

```text
Debian 13 amd64
└── systemd / kernel / firmware
    ├── Xorg + Matchbox
    │   └── TV Shell (Go + GTK3)
    │       ├── InputManager
    │       ├── AppManager
    │       ├── MediaCore
    │       ├── Search / state
    │       └── System Adapters
    ├── NetworkManager / BlueZ
    ├── PipeWire / WirePlumber
    ├── UDisks2 / logind
    └── on-demand processes
        ├── mpv
        ├── RetroArch
        ├── torrent/provider helpers
        ├── Firefox ESR streaming runtime
        └── llama.cpp
```

## Architectural style

**Modular Shell + interfaces/adapters + isolated subprocesses**.

The executable implementation follows this split under `internal/`: GTK stays
in `shell`, semantic input/focus are independent packages, Linux commands are
behind `system` adapters, and media/browser/emulator processes are isolated.

Avoid:
- one monolithic process linked to every C library;
- dozens of resident microservices;
- web runtime in the Shell;
- per-feature input/focus implementations.

## Responsibility split

The Linux kernel/Debian packages own device drivers and firmware availability. Mature services own protocols and low-level policy. TV Shell owns:
- user-facing orchestration;
- lifecycle/resource policy;
- semantic state;
- navigation;
- error translation;
- product-level contracts.

## Failure containment

A provider, media player, browser, emulator or local LLM crash must not bring down the Shell. High-risk helpers use separate processes and systemd/AppArmor boundaries.

## Source of truth

Architecture truth is Markdown/ADRs plus machine contracts in JSON. SQLite is runtime/derived state, not canonical project knowledge.
