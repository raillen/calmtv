# Debugging Playbook

## Start from the symptom
Classify: Shell/focus, app lifecycle, system adapter, playback, provider/network, storage, hardware decode or image/update.

## Evidence
Prefer:
- structured application logs;
- `journalctl` scoped to project units;
- systemd/cgroup state;
- adapter command/D-Bus response;
- mpv JSON IPC events;
- `vainfo`/playback stats;
- reproducible fixture.

## Rule
Reproduce before patching. Add a regression test or explicit hardware reproduction script whenever the failure can recur.

Raw logs containing playlist credentials/tokens must be redacted before sharing.
