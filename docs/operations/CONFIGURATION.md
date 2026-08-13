# Configuration

Configuration has three classes:

1. **System-managed:** NetworkManager/BlueZ/PipeWire/systemd state.
2. **Project settings:** TV Shell user settings stored in the project runtime DB/config path.
3. **Canonical engineering config:** `atlas.json` and repository docs/ADRs.

User settings include language, subtitle/audio preference, UI mode, library sources, resource profile and optional services.

Do not place user secrets in `atlas.json`.
