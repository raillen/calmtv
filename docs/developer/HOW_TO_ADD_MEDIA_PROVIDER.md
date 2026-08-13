# How-to — Add a Media Provider

1. Confirm source legality/distribution policy.
2. Implement only supported provider resources (catalog/search/meta/streams/subtitles).
3. Run provider outside the Shell trust boundary when untrusted/network-heavy.
4. Validate all returned URLs, paths and metadata.
5. Never pass provider strings to a shell.
6. Add deterministic fixtures for normal, empty, malformed, slow and unavailable responses.
7. Document authentication/secret handling without storing real credentials.
