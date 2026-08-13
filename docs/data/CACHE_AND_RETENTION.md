# Cache and Retention

Caches are bounded and reconstructable.

- Thumbnails/artwork: on-demand, visible-first, disk LRU.
- EPG source XML: stream/decompress into SQLite, then discard source payload unless explicitly cached.
- Provider responses: TTL based on source semantics.
- Search index: rebuildable from local state.
- LLM working context: ephemeral and garbage-collected.
- Browser profiles: persistent credentials/site state only when required; general cache bounded.
- Torrent streaming cache: temporary unless user selects “keep/download”.

Never let artwork or logs grow without a configured size/age limit.
