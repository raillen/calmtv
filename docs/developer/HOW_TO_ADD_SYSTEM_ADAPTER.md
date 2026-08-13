# How-to — Add a System Adapter

1. Write/confirm the stable project interface.
2. Prefer official D-Bus/CLI protocol.
3. Keep backend-specific types inside the adapter package.
4. Translate errors to project domain errors.
5. Provide a fake/mock implementation.
6. Add contract tests and one real integration test where possible.
7. Record hardware-specific behavior as evidence, not as global assumptions.
8. Create an ADR only if the adapter changes an architecture boundary or replaces an accepted backend.
