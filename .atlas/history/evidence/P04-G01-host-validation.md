# P04-G01 host validation

Status: PARTIAL — closure components and recovery/package contracts exist;
the target runtime cannot be fully exercised on this host.

## Passed

- ROM classification/hash and RetroArch launch argument tests pass.
- Firefox ESR kiosk/profile validation, diagnostics redaction and recovery
  reset tests pass.
- Debian packaging metadata smoke passes; the package script explicitly stages
  artifacts under `build/apt/` when Debian packaging tools are available.
- The package includes `/usr/share/xsessions/tv-shell.desktop` and an owned
  Matchbox/session launcher, allowing side-by-side desktop testing.
- `desktop-file-validate` passes for the installed session descriptor.

## Pending

- RetroArch cores and save-state return on the target image;
- Firefox ESR/Widevine behavior, if officially available;
- `dpkg-buildpackage` and clean APT smoke;
- clean Debian image boot, recovery and hardware diagnostics.
