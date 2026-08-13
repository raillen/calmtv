# NanoTube Integration

NanoTube remains a separate first-party application/project surface, integrated through TV Shell contracts rather than copied into the Shell process.

## Integration points
- Home/Continue Watching provider.
- Universal search provider.
- MediaCore player/session contract.
- Shared history/favorites where semantics match.
- MPRIS/global media controls.
- AppManager lifecycle/resource policy.
- Common design-system/navigation conventions.

## Boundary

NanoTube-specific YouTube resolution/auth/catalog logic stays inside NanoTube. TV Shell receives normalized content/playback state and does not become a YouTube client implementation itself.

If NanoTube is unavailable/crashes, Home and other media apps remain functional.
