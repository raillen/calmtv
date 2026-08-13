# Hardware Lab Setup

Maintain at least three test profiles:

1. **Minimum reference** — Atom/Celeron-class x86-64, 2 GB RAM, older Intel graphics, HDD or low-end SSD.
2. **Mid reference** — newer Intel iGPU, 4+ GB RAM.
3. **Diversity reference** — different GPU/network/Bluetooth combination.

The minimum machine is the performance gate. CI labels should identify hardware capabilities rather than vague names.

Hardware-required checks include VA-API, HDMI/audio routing, suspend/resume, Bluetooth, CEC where available, display hotplug and real remote/HID behavior.
