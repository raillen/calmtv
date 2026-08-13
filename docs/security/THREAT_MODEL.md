# Threat Model

Primary untrusted inputs:
- websites;
- external provider responses;
- torrents/magnets;
- IPTV playlists/XMLTV;
- remote media;
- archives/ROMs/media files;
- filenames/metadata sent to the organizer;
- local network remote-control clients.
- SSH clients and downloaded release metadata/packages.

Primary assets:
- user files;
- service credentials/cookies;
- network credentials;
- update integrity;
- Shell availability;
- media history/preferences.

High-priority threats:
- arbitrary code execution through untrusted data;
- path traversal/file overwrite;
- credential leakage;
- malicious provider/browser content escaping boundaries;
- local network remote abuse;
- resource exhaustion on 2 GB hardware;
- supply-chain compromise.

Remote maintenance is limited to a non-root SSH account plus a single
allowlisted sudo command. The updater pins the GitHub repository, validates the
release version/package identity and checks the published SHA-256 before APT
installation. Password SSH is only a bootstrap mode; public-key auth is the
recommended steady state.
