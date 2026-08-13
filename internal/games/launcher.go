package games

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
)

type Process interface {
	Start() error
	Wait() error
	Stop() error
}
type CommandFactory interface {
	New(ctx context.Context, name string, args ...string) Process
}

type ExecFactory struct{}

func (ExecFactory) New(ctx context.Context, name string, args ...string) Process {
	return &execProcess{command: exec.CommandContext(ctx, name, args...)}
}

type execProcess struct{ command *exec.Cmd }

func (p *execProcess) Start() error { return p.command.Start() }
func (p *execProcess) Wait() error  { return p.command.Wait() }
func (p *execProcess) Stop() error {
	if p.command.Process == nil {
		return nil
	}
	return p.command.Process.Kill()
}

type Launcher struct {
	binary        string
	factory       CommandFactory
	saveDirectory string
}

func NewLauncher(binary, saveDirectory string, factory CommandFactory) Launcher {
	return Launcher{binary: binary, saveDirectory: saveDirectory, factory: factory}
}

func (l Launcher) Launch(ctx context.Context, rom ROM) (Process, error) {
	if rom.Path == "" || !filepath.IsAbs(rom.Path) {
		return nil, fmt.Errorf("ROM path must be absolute")
	}
	if rom.Core == "" || rom.System == "" {
		return nil, fmt.Errorf("ROM is not classified")
	}
	binary := l.binary
	if binary == "" {
		binary = "retroarch"
	}
	args := []string{"-L", rom.Core, rom.Path}
	if l.saveDirectory != "" {
		args = append(args, "--save-dir", l.saveDirectory)
	}
	process := l.factory.New(ctx, binary, args...)
	if err := process.Start(); err != nil {
		return nil, fmt.Errorf("start RetroArch: %w", err)
	}
	return process, nil
}
