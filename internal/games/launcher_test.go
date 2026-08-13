package games

import (
	"context"
	"strings"
	"testing"
)

type fakeProcess struct{}

func (fakeProcess) Start() error { return nil }
func (fakeProcess) Wait() error  { return nil }
func (fakeProcess) Stop() error  { return nil }

type fakeFactory struct {
	name string
	args []string
}

func (f *fakeFactory) New(_ context.Context, name string, args ...string) Process {
	f.name, f.args = name, args
	return fakeProcess{}
}

func TestLauncherPassesROMDirectlyToRetroArch(t *testing.T) {
	factory := &fakeFactory{}
	launcher := NewLauncher("retroarch", "/state/saves", factory)
	_, err := launcher.Launch(context.Background(), ROM{Path: "/games/test.nes", System: NES, Core: "mesen"})
	if err != nil {
		t.Fatal(err)
	}
	if factory.name != "retroarch" || !strings.Contains(strings.Join(factory.args, " "), "/games/test.nes") {
		t.Fatalf("command=%s args=%v", factory.name, factory.args)
	}
}
