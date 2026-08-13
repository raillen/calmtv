package media

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"
)

type fakeTransport struct {
	commands  []string
	responses map[string]json.RawMessage
	err       error
}

func (t *fakeTransport) Command(_ context.Context, command string, _ ...any) (json.RawMessage, error) {
	t.commands = append(t.commands, command)
	if t.err != nil {
		return nil, t.err
	}
	return t.responses[command], nil
}
func (t *fakeTransport) Close() error { return nil }

func TestClientUsesSharedMpvCommands(t *testing.T) {
	transport := &fakeTransport{responses: map[string]json.RawMessage{
		"get_property": json.RawMessage(`[ {"id":1,"type":"audio","lang":"pt"},{"id":2,"type":"sub","lang":"en"} ]`),
	}}
	client := NewClient(transport)
	if err := client.Open(context.Background(), "/media/movie.mkv"); err != nil {
		t.Fatal(err)
	}
	if err := client.Pause(context.Background()); err != nil {
		t.Fatal(err)
	}
	tracks, err := client.AudioTracks(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(tracks) != 1 || tracks[0].Lang != "pt" {
		t.Fatalf("tracks = %#v", tracks)
	}
	if client.State() != StatePaused {
		t.Fatalf("state = %q", client.State())
	}
}

func TestClientRejectsInvalidVolumeAndPropagatesTransportError(t *testing.T) {
	client := NewClient(&fakeTransport{err: errors.New("socket closed")})
	if err := client.SetVolume(context.Background(), 101); err == nil {
		t.Fatal("invalid volume was accepted")
	}
	if err := client.Play(context.Background()); err == nil {
		t.Fatal("transport error was swallowed")
	}
	_ = time.Second
}

type fakeProcess struct{ stopped bool }

func (p *fakeProcess) Stop() error { p.stopped = true; return nil }

type fakeStarter struct{ process *fakeProcess }

func (s fakeStarter) Start(context.Context, string) (Process, error) { return s.process, nil }

type recordingTransport struct{ opened string }

func (t *recordingTransport) Command(_ context.Context, command string, args ...any) (json.RawMessage, error) {
	if command == "loadfile" {
		t.opened = args[0].(string)
	}
	return nil, nil
}
func (t *recordingTransport) Close() error { return nil }

func TestRuntimeOpensNetworkURLThroughMpv(t *testing.T) {
	process := &fakeProcess{}
	transport := &recordingTransport{}
	runtime := NewRuntime(fakeStarter{process: process}, func(context.Context, string) (Transport, error) {
		return transport, nil
	}, "/tmp/tv-shell-test.sock")
	if err := runtime.OpenURL(context.Background(), "https://example.invalid/live.m3u8"); err != nil {
		t.Fatal(err)
	}
	if transport.opened != "https://example.invalid/live.m3u8" {
		t.Fatalf("opened = %q", transport.opened)
	}
	if err := runtime.OpenURL(context.Background(), "file:///tmp/movie"); err == nil {
		t.Fatal("unsupported URL was accepted")
	}
	if err := runtime.Stop(); err != nil {
		t.Fatal(err)
	}
	if !process.stopped {
		t.Fatal("mpv process was not stopped")
	}
}
