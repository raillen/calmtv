# Operational Runbooks

## Shell will not start
Check service/session logs → validate display/session config → enter recovery → reset latest user UI setting if implicated.

## Wi-Fi unavailable
Confirm kernel device → firmware → NetworkManager device state → adapter translation. Do not debug Shell UI first if the OS has no device.

## HDMI has picture but no sound
Inspect PipeWire/WirePlumber nodes/profile → select HDMI output → persist project preference only after confirmed.

## Playback stutters
Check hardware decode support, selected codec/resolution, dropped frames, CPU and IO pressure before changing player code.

## Memory pressure
Inspect cgroups → identify foreground/background processes → verify lifecycle policy → confirm zram → reproduce with workload.
