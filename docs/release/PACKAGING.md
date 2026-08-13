# Packaging

Baseline distribution:
- Debian 13 amd64 base;
- project `.deb` packages;
- `debootstrap`/`live-build` image pipeline;
- Debian stock kernel/firmware;
- reproducible package lists/hooks;
- no conventional DE package set.

Third-party components remain packages/dependencies where possible rather than copied binaries.

The package accepts both Debian's modern `pkexec`/`polkitd` packages and the
older `policykit-1` transitional package, so installation remains compatible
with Q4OS releases based on different Debian versions.

The package also installs an optional display-manager session at
`/usr/share/xsessions/tv-shell.desktop`. Selecting it starts
`/usr/bin/tv-shell-session`, which owns the Matchbox process and returns to the
display manager when Calm TV exits. The existing desktop session remains the
default and is not replaced. The user service is installed but intentionally
not enabled automatically; the selected display-manager session owns startup
to prevent duplicate Shell processes.

The image definition must be version-controlled; manual post-install tweaks are release defects.
