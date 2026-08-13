# Data Migrations

Runtime DB migrations follow these rules:

1. Version every schema.
2. Back up or make rollback semantics explicit before destructive migration.
3. Test upgrade from the oldest supported release.
4. Do not silently discard user libraries/history/settings.
5. Keep canonical project configuration in `atlas.json`/Markdown, not SQLite.
6. Release notes identify migrations that affect rollback compatibility.

The MVP may use simple forward migrations; release hardening must define downgrade/restore behavior.
