package web

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Launcher struct {
	binary      string
	profileRoot string
}

func NewLauncher(binary, profileRoot string) Launcher {
	return Launcher{binary: binary, profileRoot: profileRoot}
}

func (l Launcher) Start(ctx context.Context, serviceURL string) (*exec.Cmd, error) {
	if !strings.HasPrefix(serviceURL, "https://") {
		return nil, fmt.Errorf("streaming URL must use HTTPS")
	}
	if l.binary == "" {
		l.binary = "firefox-esr"
	}
	profile := filepath.Join(l.profileRoot, "streaming")
	if err := os.MkdirAll(profile, 0700); err != nil {
		return nil, fmt.Errorf("create streaming profile: %w", err)
	}
	command := exec.CommandContext(ctx, l.binary, "--kiosk", "--no-remote", "-profile", profile, serviceURL)
	if err := command.Start(); err != nil {
		return nil, fmt.Errorf("start streaming runtime: %w", err)
	}
	return command, nil
}
