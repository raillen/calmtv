# Security Tests

- Validate systemd unit hardening and AppArmor loading.
- Provider malformed input/path traversal/URL scheme tests.
- Smart Organizer allowed-root/path traversal/collision tests.
- No shell interpolation from provider/LLM/user filenames.
- Secrets absent from logs/exported diagnostics.
- Browser/profile separation tests.
- Dependency vulnerability scan.
- Package/repository signature verification in release flow.
- Network listener defaults: bind locally only unless explicitly intended.
