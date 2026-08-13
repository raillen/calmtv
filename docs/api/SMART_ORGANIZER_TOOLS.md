# Smart Organizer Tool Contract

The LLM can only request validated semantic actions.

Allowed request types:
- `inspect_file`
- `identify_rom`
- `propose_create_directory`
- `propose_move`
- `propose_copy`
- `propose_rename`
- `search_local_metadata`

Each mutation includes source, target, reason and confidence. The executor rejects path traversal, disallowed roots, collisions without policy and unsupported operations.

`delete`, `exec`, `sudo`, package management and system configuration are not baseline AI tools.
