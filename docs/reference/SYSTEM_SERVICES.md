# System Services Reference

Resident/near-resident baseline should remain small:
- systemd/journald/udev;
- D-Bus;
- Xorg;
- Matchbox;
- TV Shell;
- NetworkManager;
- PipeWire/WirePlumber;
- BlueZ only as required by configured Bluetooth use.

On-demand:
- mpv;
- RetroArch;
- provider/torrent helpers;
- Firefox;
- llama.cpp;
- scans/importers;
- diagnostics-heavy tools.

Actual residency is validated by measurement rather than assumed from this table.
