# Smart Organizer Security

The LLM is not an authority.

Executor invariants:
- canonicalize every path;
- enforce allowed roots;
- reject traversal/symlink escape where applicable;
- detect collisions;
- preview mutation batch;
- journal before applying;
- support undo;
- no autonomous deletion;
- no arbitrary executable/tool invocation.

The model sees only the minimum filenames/metadata needed for an ambiguous batch and terminates after inference.
