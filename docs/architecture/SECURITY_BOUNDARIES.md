# Security Boundaries

Trust zones:

1. **Shell/core** — highest trust; minimal network parsing.
2. **System services** — trusted OS services reached through narrow adapters.
3. **Media processes** — mpv/RetroArch with media/ROM access only as required.
4. **Network/provider helpers** — untrusted remote input; strong sandbox.
5. **Firefox** — hostile web content; dedicated profile/sandbox/cgroup.
6. **Smart Organizer/LLM** — file mutation authority is constrained by executor policy.
7. **User sources** — playlists, media, archives and network shares are untrusted input.

Systemd hardening and AppArmor are layered; seccomp may supplement but is not treated as a complete sandbox by itself.
