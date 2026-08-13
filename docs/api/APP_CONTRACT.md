# App Contract

Each launchable application declares:
- ID, title and icon/artwork reference;
- command/entrypoint;
- resource class (`light`, `medium`, `heavy`, `exclusive-heavy`);
- background capability;
- restorable state support;
- permissions/capabilities;
- media role if applicable.

Core operations:
- `Launch(ctx, appID)`
- `Stop(ctx, appID)`
- `SaveState(ctx, appID)`
- `Restore(ctx, appID)`
- `ResourcePolicy(appID)` / `SetResourcePolicy(appID, policy)`

App manifests must not grant arbitrary privilege solely for convenience.
