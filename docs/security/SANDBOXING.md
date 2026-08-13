# Sandboxing

Baseline layers:
1. separate process;
2. dedicated systemd scope/service;
3. `NoNewPrivileges`;
4. filesystem restrictions;
5. address-family/network restrictions when possible;
6. cgroup resource limits;
7. AppArmor profile;
8. optional seccomp hardening where justified.

Priority profiles:
- Firefox: strong.
- Provider host: strong.
- Torrent engine: strong.
- Smart Organizer: strong.
- llama.cpp: strong/minimal filesystem.
- mpv: medium.
- RetroArch: medium.
- Shell: conservative profile to avoid breaking system integration.

Do not treat seccomp alone as a complete sandbox.
