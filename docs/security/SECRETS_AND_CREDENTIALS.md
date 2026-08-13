# Secrets and Credentials

Potential secrets:
- Wi-Fi credentials managed by NetworkManager/system policy;
- SMB credentials;
- IPTV URL tokens;
- provider API keys;
- streaming cookies/session data;
- optional remote-pairing tokens.

Rules:
- never store secrets in Git/Atlas docs;
- never include them in Project Intelligence;
- redact diagnostics;
- use OS/profile storage appropriate to the service;
- define deletion/reset behavior;
- separate browser profiles by purpose.
