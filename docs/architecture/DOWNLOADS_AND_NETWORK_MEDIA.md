# Downloads and Network Media

## DownloadManager

One shared runtime model represents HTTP/provider, podcast, torrent and future recording downloads. UI shows source, progress, speed, destination, retention and errors.

Background downloads are allowed only within AppManager resource policy and may be deprioritized during games/video.

## Network media

SMB is the first network-filesystem target. Discovery may use mDNS/Avahi; direct server configuration must remain available. DLNA/UPnP is optional/future and should use a mature helper/library rather than a new protocol implementation.

Network scans are event/manual/TTL driven rather than permanent aggressive polling.
