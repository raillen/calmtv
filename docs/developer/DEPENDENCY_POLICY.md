# Dependency Policy

Before adding project code or a dependency:

1. Does Linux/Debian already provide the capability?
2. Is there a stable Freedesktop/system service or D-Bus contract?
3. Is there an official small CLI suitable behind an adapter?
4. Is there a mature library with acceptable license/maintenance?
5. Only then implement the protocol/feature ourselves.

Prefer protocol/CLI boundaries over fragile deep bindings in V1. Add FFI only when measurements or required capability justify it.

Every dependency change reviews:
- license;
- Debian availability/packaging;
- maintenance status;
- binary/RAM footprint;
- security surface;
- replacement strategy.
