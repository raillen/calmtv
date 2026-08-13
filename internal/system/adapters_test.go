package system

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

type fakeRunner struct {
	command string
	args    []string
	output  []byte
	err     error
}

func (f *fakeRunner) Run(_ context.Context, command string, args ...string) ([]byte, error) {
	f.command, f.args = command, args
	return f.output, f.err
}

func TestNetworkAdapterUsesNetworkManagerContract(t *testing.T) {
	runner := &fakeRunner{output: []byte("Home:wlan0:yes\nGuest:wlan0:no\n")}
	networks, err := NewNetworkAdapter(runner).List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if runner.command != "nmcli" || !reflect.DeepEqual(networks, []Network{
		{Name: "Home", Device: "wlan0", Active: true},
		{Name: "Guest", Device: "wlan0", Active: false},
	}) {
		t.Fatalf("unexpected network adapter result: %#v, command %s", networks, runner.command)
	}
}

func TestDisplayAdapterRejectsShellSyntax(t *testing.T) {
	runner := &fakeRunner{}
	err := NewDisplayAdapter(runner).SetMode(context.Background(), "HDMI-1; reboot", "1920x1080")
	var serviceErr *ServiceError
	if !errors.As(err, &serviceErr) || serviceErr.Code != ErrorInvalid {
		t.Fatalf("error = %v, want invalid service error", err)
	}
	if runner.command != "" {
		t.Fatal("invalid display command reached runner")
	}
}

func TestAdapterTranslatesPermissionFailure(t *testing.T) {
	runner := &fakeRunner{err: errors.New("not authorized")}
	err := NewPowerAdapter(runner).PowerOff(context.Background())
	var serviceErr *ServiceError
	if !errors.As(err, &serviceErr) || serviceErr.Code != ErrorPermission {
		t.Fatalf("error = %v, want permission error", err)
	}
}

func TestAudioAdapterControlsMuteThroughWpctl(t *testing.T) {
	runner := &fakeRunner{}
	if err := NewAudioAdapter(runner).SetMute(context.Background(), "@DEFAULT_AUDIO_SINK@", true); err != nil {
		t.Fatal(err)
	}
	if runner.command != "wpctl" || !reflect.DeepEqual(runner.args, []string{"set-mute", "@DEFAULT_AUDIO_SINK@", "1"}) {
		t.Fatalf("command=%s args=%v", runner.command, runner.args)
	}
}
