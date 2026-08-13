# How-to — Organize Downloads

1. Open **Files** → **Downloads**.
2. Choose **Organize**.
3. Review the proposed moves/renames.
4. Confirm all or individual actions.
5. Use **Undo** if the result is wrong.

The organizer tries deterministic rules first. The local AI fallback is used only when enabled and when filenames/context remain ambiguous. AI never receives unrestricted shell access and does not delete files automatically.
# MVP behavior

The download service writes to a configured Downloads directory using a
`.part` file and renames only after completion. Names are confined to the
destination directory; arbitrary shell commands and path traversal are
rejected.
