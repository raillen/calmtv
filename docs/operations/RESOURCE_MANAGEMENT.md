# Resource Management

AppManager uses systemd/cgroups to implement product policy.

Use `MemoryHigh` as the main pressure signal and `MemoryMax` as a last boundary after hardware benchmarks define safe values. Use CPU/IO weights to deprioritize downloads/background work during interactive apps/games.

ZRAM is the preferred first line of swap pressure mitigation on 2 GB systems. Disk swap is a last-resort safety net and must not be relied on for acceptable UX.

At idle, torrent engines, mpv, RetroArch, Firefox and llama.cpp should not exist as processes.
