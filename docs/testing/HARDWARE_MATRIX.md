# Hardware Matrix

## Tier 1 — release gate
Minimum 64-bit Atom/Celeron-class reference machine, 2 GB RAM, Intel iGPU and representative HDD/SSD.

The available Q40S Atom notebook is the first Tier 1 candidate. It should be
registered with the exact CPU model, Debian version, kernel, RAM, storage,
display and Wi-Fi/Bluetooth chipsets after running `scripts/target-preflight`.

## Tier 2 — compatibility
Newer Intel x86-64 machine, 4+ GB RAM.

## Tier 3 — diversity/experimental
Additional AMD/Nvidia/network/Bluetooth configurations as hardware becomes available.

## Validate
- boot/session;
- 720p/1080p modes;
- VA-API capabilities and dropped frames;
- HDMI audio routing;
- Wi-Fi reconnect;
- Bluetooth pairing/reconnect;
- USB mount/unmount;
- suspend/resume;
- remote HID;
- memory pressure;
- optional CEC.

The reproducible first pass is:

```bash
./scripts/target-preflight build/target-preflight
```

Then select the `TV Shell` session at login, run the smoke flow, and preserve
the generated environment and boot files with the test evidence.

With TV Shell selected, run:

```bash
./scripts/target-session-check build/target-session-check
cat build/target-session-check/session.txt
```

The check confirms Xorg, Matchbox and the Shell are alive and fails if a known
resident compositor is detected.

If Debian does not support a device, TV Shell does not maintain a private driver fork in V1.
