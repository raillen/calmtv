# Development Setup

## Baseline host

Use a modern Linux development machine; the target image is Debian 13 amd64.

Required development tools are expected to include:
- Git;
- Go toolchain;
- GTK3/gotk3 build dependencies;
- systemd/D-Bus development utilities;
- mpv/FFmpeg;
- SQLite;
- Debian packaging tools (`debootstrap`, `live-build`, `debhelper`, `dh-golang`);
- GitHub Actions-compatible CI scripts.

## Workflow

1. Read `ENTRYPOINT.md`, `PROJECT_STATE.md` and the active Goal.
2. Build/test modules on the development host.
3. Build the Debian package/image from scripts, not manual machine state.
4. Run VM/synthetic tests where possible.
5. Run hardware-required gates on labeled self-hosted machines.
6. Record evidence before marking a Goal complete.

Do not install Docker in the target runtime solely for development reproducibility; containers are a build/CI tool here.
