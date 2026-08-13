package recovery

import (
	"context"
	"os"
	"path/filepath"
)

type Action string

const (
	RestartShell Action = "restart-shell"
	ResetUI      Action = "reset-ui"
	Diagnostics  Action = "diagnostics"
	Reboot       Action = "reboot"
	PowerOff     Action = "power-off"
)

type Manager struct {
	configRoot string
	run        func(context.Context, string, ...string) error
}

func NewManager(configRoot string, run func(context.Context, string, ...string) error) Manager {
	return Manager{configRoot: configRoot, run: run}
}

func (m Manager) Execute(ctx context.Context, action Action) error {
	switch action {
	case ResetUI:
		return os.RemoveAll(filepath.Join(m.configRoot, "ui"))
	case Reboot:
		return m.run(ctx, "loginctl", "reboot")
	case PowerOff:
		return m.run(ctx, "loginctl", "poweroff")
	default:
		return nil
	}
}
