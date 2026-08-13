# Recovery

Required recovery scenarios:
- Shell crashes repeatedly on boot.
- Bad configuration prevents Home.
- Update/package is broken.
- Display mode becomes unusable.
- Storage/database migration fails.

Baseline behavior:
1. detect repeated Shell startup failure;
2. enter a minimal recovery menu/session;
3. allow reset of Shell settings/display profile;
4. expose diagnostics;
5. support package/update rollback strategy defined for the release;
6. provide terminal access only as an advanced recovery tool.

Recovery must not depend on the normal Shell being healthy.

The packaged launcher counts consecutive startup failures under the user
runtime directory. After three failures it enters a recovery-marked launch
instead of silently creating an endless blank restart loop. The recovery
package can reset only the UI configuration, preserving runtime state.

When Calm TV is installed beside another desktop, recovery is scoped to the
Calm TV session. Selecting the previous desktop session at login remains the
fallback if the Calm TV session cannot start.

From Home, the remote's `Menu` action opens Diagnostics. The **Terminal
avançado** action there starts an installed graphical terminal as the current
user, never as root. It is an optional maintenance escape hatch; normal users
do not need it for ordinary media flows.
