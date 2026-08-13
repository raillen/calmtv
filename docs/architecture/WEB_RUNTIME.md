# Web Runtime

Firefox ESR is an **optional on-demand runtime**, not part of the Shell rendering stack.

## Commercial streaming profile
- one service/session;
- kiosk/fullscreen launch;
- Widevine/DRM only through supported official browser paths;
- no generic tab workflow;
- process terminates when leaving;
- restrictive cgroup and AppArmor profile.

## Optional generic browser profile
May be packaged separately:
- maximum 3 tabs by default;
- uBlock Origin managed by policy;
- inactive tabs discarded;
- browser terminates on exit from the application;
- no guarantee that all websites/services support Linux.

WPE WebKit is an experimental future benchmark, not a V1 architectural dependency.
