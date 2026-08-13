# Games and Emulation Architecture

TV Shell owns library/navigation; RetroArch owns emulation.

## Library pipeline
1. Scan configured folders on demand.
2. Detect known ROM/archive types.
3. Compute checksum/serial when useful.
4. Match Libretro-compatible databases.
5. Normalize title/system.
6. Store lightweight runtime metadata in SQLite.

## Launch
AppManager starts `retroarch -L <core> <rom>` inside a resource scope. Shell is not a second emulation frontend process.

## Baseline cores/systems
Select lightweight, maintained cores during hardware benchmarking. Initial product targets NES, SNES, GB/GBC, GBA, Master System, Mega Drive/Genesis and Game Gear.

Shaders, run-ahead and rewind are disabled by the low-end default profile unless benchmarks prove headroom.
