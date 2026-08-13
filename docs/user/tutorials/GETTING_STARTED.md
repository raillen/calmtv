# Tutorial — Getting Started

> This describes the intended V1 product flow. The current checkout provides the development shell; installer labels remain provisional until image validation.

On an existing Debian amd64 installation, TV Shell can be installed as an
additional login session. It does not replace the current desktop. Install
the `.deb`, choose `TV Shell` in the display manager's session menu, and
choose the original desktop there to return.

1. Boot the installation image and install TV Shell to the target amd64 machine.
2. First boot enters the setup flow rather than a conventional Linux desktop.
3. Select display/resolution if automatic selection is unsuitable.
4. Connect network if needed.
5. Pair Bluetooth devices only when used; a USB HID remote/keyboard works without Bluetooth pairing.
6. Choose preferred audio/subtitle languages.
7. Select Simple or Advanced Mode.
8. Add optional media sources: local folders/USB, IPTV playlist, network share and game library.
9. Finish setup and enter Home.

The system should never require a terminal for the normal setup path.
# Status

The current development build boots a GTK3 Home with remote navigation and
screen surfaces for Quick Settings, local media, NanoTube, IPTV, games,
files/downloads, diagnostics and power. Service-backed actions report a
recoverable user-facing status when the corresponding Debian service is not
installed or authorized. An optional Streaming tile launches Firefox ESR only
when a service URL and the packaged runtime are configured.
