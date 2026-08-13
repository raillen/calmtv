# Product

## Vision

Create a Linux-based TV/media appliance that makes very old x86-64 hardware feel intentional rather than merely tolerated. The system should boot directly into a calm, remote-first interface, keep background resource use extremely low, reuse mature Linux services instead of reimplementing them, and make common media tasks accessible without exposing desktop/Linux internals.

## Product shape

Calm TV is a **dedicated appliance shell**, not a general-purpose desktop environment:

- Home and universal search.
- NanoTube integration.
- Local/network media.
- IPTV with EPG.
- Music, radio and podcasts.
- Retro console emulation.
- File/download management.
- Optional commercial streaming web runtime.
- Optional phone/CEC control.
- Smart Organizer with constrained, on-demand local AI fallback.
- System settings for network, Bluetooth, audio, display, storage, power and updates.

## Primary differentiators

1. **Low idle cost:** no compositor, no browser engine in the Shell, no resident media indexer or LLM.
2. **Remote-first:** every essential path works with D-pad + OK + Back + Home.
3. **Lifecycle-aware:** one heavy foreground application by default; previous apps restore state rather than remain resident.
4. **Linux hidden behind product language:** users select “Audio output” rather than “PipeWire node”.
5. **Reuse over reinvention:** kernel/Debian/services handle hardware and protocols; project code owns orchestration and experience.
6. **Local-first privacy:** no telemetry by default and local AI only when explicitly needed.
