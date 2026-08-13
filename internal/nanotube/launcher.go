package nanotube

import (
	"context"
	"fmt"
	"os/exec"
)

// Process is intentionally small so NanoTube can be started on demand and
// stopped when the user returns to Home without coupling the Shell to its
// implementation.
type Process interface {
	Stop() error
}

type CommandFactory interface {
	New(context.Context, string, ...string) Process
}

type ExecFactory struct{}

func (ExecFactory) New(ctx context.Context, name string, args ...string) Process {
	return &execProcess{command: exec.CommandContext(ctx, name, args...)}
}

type execProcess struct{ command *exec.Cmd }

func (p *execProcess) Start() error { return p.command.Start() }

func (p *execProcess) Stop() error {
	if p.command.Process == nil {
		return nil
	}
	return p.command.Process.Kill()
}

type Launcher struct {
	binary  string
	factory CommandFactory
}

func NewLauncher(binary string, factory CommandFactory) Launcher {
	return Launcher{binary: binary, factory: factory}
}

func (l Launcher) Start(ctx context.Context) (Process, error) {
	binary := l.binary
	if binary == "" {
		binary = "nanotube"
	}
	if l.factory == nil {
		return nil, fmt.Errorf("NanoTube command factory is not configured")
	}
	process := l.factory.New(ctx, binary)
	startable, ok := process.(interface{ Start() error })
	if !ok {
		return nil, fmt.Errorf("NanoTube process is not startable")
	}
	if err := startable.Start(); err != nil {
		return nil, fmt.Errorf("start NanoTube: %w", err)
	}
	return process, nil
}
