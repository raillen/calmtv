package appmanager

import (
	"context"
	"testing"
)

type fakeProcess struct{ stopped bool }

func (p *fakeProcess) Wait() error                { return nil }
func (p *fakeProcess) Stop(context.Context) error { p.stopped = true; return nil }

type fakeStarter struct {
	started   []string
	processes map[string]*fakeProcess
}

func (s *fakeStarter) Start(_ context.Context, manifest Manifest, _ ResourcePolicy) (Process, error) {
	if s.processes == nil {
		s.processes = make(map[string]*fakeProcess)
	}
	process := &fakeProcess{}
	s.started = append(s.started, manifest.ID)
	s.processes[manifest.ID] = process
	return process, nil
}

func TestExclusiveLaunchStopsPreviousProcess(t *testing.T) {
	starter := &fakeStarter{}
	manager := NewManager(starter, nil)
	if err := manager.Register(Manifest{ID: "mpv", Command: "mpv", ResourceClass: Heavy}); err != nil {
		t.Fatal(err)
	}
	if err := manager.Register(Manifest{ID: "retroarch", Command: "retroarch", ResourceClass: ExclusiveHeavy, Exclusive: true}); err != nil {
		t.Fatal(err)
	}
	if err := manager.Launch(context.Background(), "mpv"); err != nil {
		t.Fatal(err)
	}
	if err := manager.Launch(context.Background(), "retroarch"); err != nil {
		t.Fatal(err)
	}
	if !starter.processes["mpv"].stopped {
		t.Fatal("exclusive launch did not stop previous app")
	}
	if manager.Foreground() != "retroarch" {
		t.Fatalf("foreground = %q", manager.Foreground())
	}
}

func TestManagerRequiresExplicitBackgroundPermission(t *testing.T) {
	starter := &fakeStarter{}
	manager := NewManager(starter, nil)
	if err := manager.Register(Manifest{ID: "media", Command: "mpv"}); err != nil {
		t.Fatal(err)
	}
	if err := manager.Launch(context.Background(), "media"); err != nil {
		t.Fatal(err)
	}
	if err := manager.Background("media"); err == nil {
		t.Fatal("background execution was allowed without opt-in")
	}
	if err := manager.Register(Manifest{ID: "music", Command: "mpv", BackgroundAllowed: true, ResourceClass: Light}); err != nil {
		t.Fatal(err)
	}
	if err := manager.Launch(context.Background(), "music"); err != nil {
		t.Fatal(err)
	}
	if err := manager.Background("music"); err != nil {
		t.Fatal(err)
	}
	if manager.Foreground() != "" {
		t.Fatalf("foreground = %q", manager.Foreground())
	}
	policy, err := manager.ResourcePolicy("music")
	if err != nil || policy.MemoryMax != DefaultPolicy(Light).MemoryMax {
		t.Fatalf("policy = %#v, err = %v", policy, err)
	}
	custom := DefaultPolicy(Light)
	custom.CPUWeight = 80
	if err := manager.SetResourcePolicy("music", custom); err != nil {
		t.Fatal(err)
	}
	policy, err = manager.ResourcePolicy("music")
	if err != nil || policy.CPUWeight != 80 {
		t.Fatalf("custom policy = %#v, err = %v", policy, err)
	}
}
