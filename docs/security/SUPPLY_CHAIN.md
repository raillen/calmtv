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

The Q4OS bootstrap updater currently uses GitHub HTTPS plus the release
checksum and package metadata validation. This is suitable for hardware
testing, but signed APT repository metadata or signed release artifacts remain
required before unattended production updates.
