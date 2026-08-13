# Product and Engineering Principles

1. **Appliance before desktop.** Design for television use, not mouse-first windows.
2. **Light by architecture.** The first optimization is not starting a process.
3. **Deterministic before AI.** Rules, MIME, metadata, hashes and databases resolve tasks before invoking an LLM.
4. **One interaction model.** HID remote, keyboard, gamepad, CEC and phone control map to shared InputActions.
5. **One media contract.** First-party media apps share MediaCore/MPRIS concepts instead of inventing controls per app.
6. **State over background residency.** Save/recreate state when practical; do not keep applications alive merely to simulate multitasking.
7. **Adapters around mature services.** NetworkManager, BlueZ, PipeWire/WirePlumber, UDisks2, logind and XRandR remain replaceable behind project contracts.
8. **Errors in user language.** Raw D-Bus/CLI/service failures are diagnostic details, not UX.
9. **Evidence before optimization claims.** Performance budgets are provisional until measured on reference hardware.
10. **Content boundary is explicit.** No DRM bypass or piracy-focused source distribution.
11. **Canonical before generated.** Markdown/JSON in Git are truth; docs site/adapters/caches are derived.
12. **Least context/workforce.** Agentic development follows Project Atlas LPC/PCA.
