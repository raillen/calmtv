# Configuration

Configuration has three classes:

1. **System-managed:** NetworkManager/BlueZ/PipeWire/systemd state.
2. **Project settings:** Calm TV user settings stored in the project runtime DB/config path.
3. **Canonical engineering config:** `atlas.json` and repository docs/ADRs.

User settings include language, subtitle/audio preference, UI mode, library sources, resource profile and optional services.

Do not place user secrets in `atlas.json`.

## Múltiplos monitores

Calm TV usa uma única superfície de interface por vez e abre em tela cheia no
monitor primário. Isso evita que a Home seja esticada entre a tela do notebook
e um monitor externo. Para escolher outro monitor durante testes, use o índice
GDK/X11 antes de iniciar a sessão:

```bash
TV_SHELL_MONITOR=1 /usr/bin/tv-shell
```

O índice `0` normalmente representa o primeiro monitor detectado. A seleção
persistente e amigável por nome do conector será integrada às Quick Settings.
