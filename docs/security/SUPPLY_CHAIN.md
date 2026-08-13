# Supply Chain

For every shipped dependency/package:
- prefer Debian repository packages when suitable;
- pin Go module versions and verify checksums;
- review licenses;
- use signed package repositories;
- record build inputs for release artifacts;
- run `govulncheck` and relevant package vulnerability checks;
- minimize custom binaries downloaded at runtime;
- update browser/DRM/security-sensitive components promptly.

Official provider catalogs and update channels are part of the trust boundary and require integrity controls.
