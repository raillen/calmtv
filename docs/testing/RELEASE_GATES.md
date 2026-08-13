# Release Gates

A release candidate requires:
- all locked Goal acceptance evidence;
- clean package/image build;
- JSON schema/Atlas validation;
- Markdown link validation;
- Go tests/lint/vulnerability scan;
- hardware matrix pass for promised capabilities;
- performance budgets pass or documented approved amendment;
- security sandbox checks;
- migration/update/recovery exercise, including automatic check, confirmation,
  interrupted download, invalid signature/checksum, restart and rollback;
- user/support/release docs complete;
- checksums/artifact provenance recorded.

A failing gate is not waived by changing acceptance criteria after the Goal was locked without an explicit amendment.
