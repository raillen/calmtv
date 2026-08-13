# Release Process

1. Lock the release Goal.
2. Freeze dependency/package changes except fixes.
3. Build packages/image from clean CI.
4. Run unit/integration/static/security checks.
5. Run physical hardware matrix.
6. Exercise update/migration/recovery.
7. Validate canonical docs, local links and Atlas JSON.
8. Produce checksums and release notes.
9. Release Verifier independently checks evidence.
10. Publish signed artifacts/repository metadata.

Release is evidence-driven; build success alone is insufficient.
