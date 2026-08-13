# Codebase Tour

The implementation repository should converge on this layout:

```text
cmd/
  tv-shell/
  tv-diagnostics/
  tv-organizer/
  tv-provider-host/

internal/
  appmanager/
  input/
  media/
  database/
  search/
  organizer/
  adapters/
    networkmanager/
    bluez/
    wpctl/
    xrandr/
    udisks/
    logind/
    mpv/
    mpris/
    retroarch/
  ui/

ui/
  *.ui
  theme.css
  assets/

packaging/debian/
image/live-build/
security/systemd/
security/apparmor/
tests/
  unit/
  integration/
  ui/
  hardware/
  fixtures/
```

The first implementation slice is now present:

- `cmd/tv-shell`: GTK entrypoint;
- `internal/shell`: GtkBuilder Home, GTK CSS, Stack navigation and screen wiring;
- `internal/input` and `internal/focus`: semantic actions and central focus;
- `internal/system`: Linux service adapters and Quick Settings facade;
- `internal/appmanager`: lifecycle/resource policy boundary;
- `internal/media`: mpv JSON IPC client/runtime;
- `internal/iptv`: bounded M3U/XMLTV streaming parsers;
- `internal/state`: SQLite runtime state and migrations;
- `internal/library`, `internal/downloads`, `internal/games`: local media, downloads and deterministic ROM classification;
- `internal/diagnostics`, `internal/recovery`, `internal/web`, `internal/nanotube`: MVP support boundaries.

The Debian image and hardware behavior remain build/evidence work, not claims
that this checkout has already booted on the reference machine.

The highest-value boundaries are `AppManager`, `InputManager`, `MediaCore` and adapter interfaces. Agents should prefer implementing behind these boundaries rather than adding cross-cutting helpers directly to UI code.
