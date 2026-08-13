# System Adapter Contracts

## Network
List devices/networks, connect/disconnect, active connection, error state.

## Bluetooth
Power, discover, pair, connect/disconnect, device status.

## Audio
List outputs, select default, get/set volume, mute.

## Display
List outputs/modes, primary output, resolution/refresh, mirror/extend/disable.

## Storage
List removable volumes, mount/unmount, free space.

## Power
Suspend, reboot, poweroff and inhibitor semantics.

Backends must be mockable and must translate implementation-specific errors into stable project error categories.
