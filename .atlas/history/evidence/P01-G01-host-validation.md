# P01-G01 host validation

Status: PARTIAL — host checks only; Goal remains `EXECUTING`.

## Passed on this host

- `make test` passes with SQLite FTS5 for input, focus, adapters, AppManager,
  media, IPTV, state, library, downloads, games, diagnostics, recovery, web
  and NanoTube contracts.
- `make build` produces an amd64 GTK3 shell binary; `go vet` also passes.
- A windowed GTK smoke remains alive for three seconds on the host display.
- The Debian package metadata installs an additional `TV Shell` display-manager
  session without replacing the existing desktop session.
- `scripts/target-preflight` and `scripts/target-session-check` now provide the
  reproducible evidence collection path for the Q40S Atom target.
- Host shell PSS measured at 23,801 KiB while idle in the final windowed smoke
  (RSS observed at 45,412 KiB); this
  is not a reference Atom/Celeron baseline.
- GTK3 development headers are available and the shell source uses GtkBuilder,
  GTK CSS, central semantic input and central focus.
- Atlas documentation and JSON/compatibility projections validate.

## Pending external/runtime evidence

- clean Debian amd64 image boot through systemd/Xorg/Matchbox;
- GTK shell smoke at 720p and 1080p;
- no compositor/process inventory;
- cold boot, idle PSS and D-pad p95 on reference Atom/Celeron hardware;
- recovery after repeated shell failure.

The current host does not provide `live-build`, `dpkg-buildpackage`, Matchbox,
RetroArch, `vainfo` or a headless X server. These are recorded as unavailable
tools, not as successful gates.
