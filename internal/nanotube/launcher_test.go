package nanotube

import (
	"context"
	"errors"
	"testing"
)

type fakeProcess struct {
	started bool
	stopped bool
}

func (p *fakeProcess) Start() error { p.started = true; return nil }
func (p *fakeProcess) Stop() error  { p.stopped = true; return nil }

type fakeFactory struct{ process *fakeProcess }

func (f fakeFactory) New(context.Context, string, ...string) Process { return f.process }

func TestLauncherStartsOnlyOnDemand(t *testing.T) {
	process := &fakeProcess{}
	started, err := NewLauncher("nanotube-tv", fakeFactory{process: process}).Start(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if started != process || !process.started {
		t.Fatalf("process = %#v, started = %v", started, process.started)
	}
}

type nonStartableFactory struct{}

func (nonStartableFactory) New(context.Context, string, ...string) Process {
	return unavailableProcess{}
}

type unavailableProcess struct{}

func (unavailableProcess) Stop() error { return errors.New("not started") }

func TestLauncherRejectsNonStartableProcess(t *testing.T) {
	if _, err := NewLauncher("nanotube", nonStartableFactory{}).Start(context.Background()); err == nil {
		t.Fatal("non-startable process was accepted")
	}
}
