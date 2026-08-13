# Diagnostics

The diagnostics tool should report:
- OS/image/build version;
- CPU/RAM/storage;
- GPU/VA-API codecs;
- active display/modes;
- audio outputs;
- network/Bluetooth status;
- cgroup/app memory;
- failed project units;
- mpv/RetroArch/browser availability;
- disk free space;
- recent redacted project errors.

Export format should be human-readable text/JSON with secrets redacted. Never export playlist tokens, passwords or browser cookies.

`internal/diagnostics` collects required host/service surfaces through a
command boundary, including RAM/zram and libinput, and redacts lines
containing secrets, tokens, passwords or cookies. `scripts/measure-shell`
records environment metadata; measured PSS, boot and input values still
require a real target or VM run.

The Shell writes the latest redacted report atomically to
`~/.local/share/tv-shell/diagnostics/latest.json`.
