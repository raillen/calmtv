# Planned Development / Diagnostic Commands

Exact wrappers will be added with implementation.

Framework-level:
- `atlas validate`
- `atlas goal list`
- `atlas context plan`
- `atlas report summary`

Project development:
- `go test ./...`
- `golangci-lint run`
- `govulncheck ./...`

System/performance evidence:
- `systemd-analyze`
- `systemd-cgtop`
- `systemd-analyze security <unit>`
- `vainfo`
- `ffprobe`
- `stress-ng`
- project hardware test scripts

Do not make a CLI command canonical until it exists in code; this document distinguishes planned commands from Project Atlas commands.
