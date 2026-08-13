package appmanager

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
)

type SystemdStarter struct{}

func (SystemdStarter) Start(ctx context.Context, manifest Manifest, policy ResourcePolicy) (Process, error) {
	if err := policy.Validate(); err != nil {
		return nil, err
	}
	unit := "tv-shell-app-" + manifest.ID
	args := []string{"--user", "--scope", "--unit", unit,
		"-p", "MemoryHigh=" + policy.MemoryHigh,
		"-p", "MemoryMax=" + policy.MemoryMax,
		"-p", "CPUWeight=" + strconv.Itoa(policy.CPUWeight),
		"-p", "IOWeight=" + strconv.Itoa(policy.IOWeight),
		"--", manifest.Command}
	args = append(args, manifest.Args...)
	command := exec.CommandContext(ctx, "systemd-run", args...)
	if err := command.Start(); err != nil {
		return nil, fmt.Errorf("systemd-run: %w", err)
	}
	return &systemdProcess{command: command, unit: unit}, nil
}

type systemdProcess struct {
	command *exec.Cmd
	unit    string
}

func (p *systemdProcess) Wait() error { return p.command.Wait() }

func (p *systemdProcess) Stop(ctx context.Context) error {
	stop := exec.CommandContext(ctx, "systemctl", "--user", "stop", p.unit)
	if err := stop.Run(); err != nil {
		_ = p.command.Process.Kill()
		return err
	}
	return nil
}
