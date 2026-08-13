# Test Strategy

Test from contracts and Goal acceptance.

Layers:
- Go unit tests for domain logic/parsers/state.
- Adapter contract tests with fakes.
- Integration tests against D-Bus/CLI services where practical.
- UI focus/navigation tests with deterministic screen fixtures.
- Media playback fixtures.
- Security/static dependency checks.
- Image/boot smoke tests.
- Physical hardware tests for capabilities VMs cannot prove.

A green generic CI run cannot substitute for reference-hardware evidence on VA-API, HDMI, Bluetooth, CEC or low-memory behavior.
