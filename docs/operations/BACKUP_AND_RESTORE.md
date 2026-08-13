# Backup and Restore

Back up user-controlled configuration/state, not reconstructable caches.

Backup candidates:
- TV Shell settings database;
- source definitions excluding secrets when separately managed;
- favorites/history if the user opts in;
- game saves/save states;
- organizer transaction history needed for undo;
- custom input mappings.

Do not back up:
- thumbnail cache;
- EPG source payload cache;
- temporary torrent cache;
- Atlas runtime context/cache.

Restore must be version-aware and tested before release.
