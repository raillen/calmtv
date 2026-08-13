# P02-G01 host validation

Status: PARTIAL — host contracts and tests pass; target-service gates remain
pending.

## Passed

- AppManager tests cover exclusive foreground lifecycle, explicit background
  permission, persisted-state hooks and resource policy lookup.
- Adapter tests cover NetworkManager, Bluetooth, wpctl mute/volume, xrandr,
  UDisks2 and logind error translation.
- Quick Settings screens use the facade and adapter contracts rather than
  invoking service CLIs from GTK callbacks.
- zram-generator configuration and a systemd zram health unit are packaged.

## Pending

- cgroups v2/systemd-run behavior on a clean Debian target;
- active zram and service integration on the reference machine;
- remote configuration smoke with NetworkManager, BlueZ, PipeWire and
  display/storage services;
- final Debian package/image build.
