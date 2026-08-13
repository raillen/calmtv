# Smart Organizer

## Decision pipeline

1. File extension / MIME.
2. Filename parser.
3. Metadata (ffprobe/tags/etc.).
4. Hash/checksum/database lookup.
5. Domain rules such as Libretro ROM databases.
6. Only if confidence remains insufficient: on-demand local LLM.

## LLM boundary

llama.cpp runs as a separate process and must produce constrained structured output. The LLM proposes; a Go validator/executor decides.

Allowed semantic operations:
- inspect metadata;
- propose folder creation;
- propose move/copy/rename;
- identify/organize ROMs;
- search local indexed metadata.

Deletion is not an autonomous AI operation in the baseline.

## Safety
- allowed directory allowlist;
- no shell;
- no `/etc`, `/boot`, `/root`, SSH keys or arbitrary Home access;
- preview and undo for mutations;
- confidence thresholds;
- clear user confirmation for ambiguous cases.
