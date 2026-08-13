package media

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

type PlaybackState string

const (
	StateStopped PlaybackState = "stopped"
	StatePlaying PlaybackState = "playing"
	StatePaused  PlaybackState = "paused"
	StateError   PlaybackState = "error"
)

type Track struct {
	ID    int    `json:"id"`
	Type  string `json:"type"`
	Lang  string `json:"lang"`
	Title string `json:"title"`
}

type Transport interface {
	Command(ctx context.Context, command string, args ...any) (json.RawMessage, error)
	Close() error
}

type Client struct {
	mu        sync.Mutex
	transport Transport
	state     PlaybackState
}

func NewClient(transport Transport) *Client {
	return &Client{transport: transport, state: StateStopped}
}

func (c *Client) Open(ctx context.Context, path string) error {
	if path == "" {
		return errors.New("media path is empty")
	}
	if _, err := c.command(ctx, "loadfile", path, "replace"); err != nil {
		return fmt.Errorf("open media: %w", err)
	}
	c.setState(StatePlaying)
	return nil
}

func (c *Client) Play(ctx context.Context) error {
	if _, err := c.command(ctx, "set_property", "pause", false); err != nil {
		return fmt.Errorf("play media: %w", err)
	}
	c.setState(StatePlaying)
	return nil
}

func (c *Client) Pause(ctx context.Context) error {
	if _, err := c.command(ctx, "set_property", "pause", true); err != nil {
		return fmt.Errorf("pause media: %w", err)
	}
	c.setState(StatePaused)
	return nil
}

func (c *Client) Stop(ctx context.Context) error {
	if _, err := c.command(ctx, "stop"); err != nil {
		return fmt.Errorf("stop media: %w", err)
	}
	c.setState(StateStopped)
	return nil
}

func (c *Client) Seek(ctx context.Context, position time.Duration) error {
	seconds := position.Seconds()
	if seconds < 0 {
		seconds = 0
	}
	if _, err := c.command(ctx, "seek", seconds, "absolute"); err != nil {
		return fmt.Errorf("seek media: %w", err)
	}
	return nil
}

func (c *Client) SetVolume(ctx context.Context, percent int) error {
	if percent < 0 || percent > 100 {
		return errors.New("volume must be between 0 and 100")
	}
	if _, err := c.command(ctx, "set_property", "volume", percent); err != nil {
		return fmt.Errorf("set volume: %w", err)
	}
	return nil
}

func (c *Client) Tracks(ctx context.Context, trackType string) ([]Track, error) {
	value, err := c.command(ctx, "get_property", "track-list")
	if err != nil {
		return nil, fmt.Errorf("read tracks: %w", err)
	}
	var tracks []Track
	if err := json.Unmarshal(value, &tracks); err != nil {
		return nil, fmt.Errorf("decode tracks: %w", err)
	}
	filtered := tracks[:0]
	for _, track := range tracks {
		if track.Type == trackType {
			filtered = append(filtered, track)
		}
	}
	return filtered, nil
}

func (c *Client) AudioTracks(ctx context.Context) ([]Track, error)    { return c.Tracks(ctx, "audio") }
func (c *Client) SubtitleTracks(ctx context.Context) ([]Track, error) { return c.Tracks(ctx, "sub") }

func (c *Client) Position(ctx context.Context) (time.Duration, error) {
	return c.propertyDuration(ctx, "time-pos")
}

func (c *Client) Duration(ctx context.Context) (time.Duration, error) {
	return c.propertyDuration(ctx, "duration")
}

func (c *Client) State() PlaybackState {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.state
}

func (c *Client) command(ctx context.Context, name string, args ...any) (json.RawMessage, error) {
	if c.transport == nil {
		return nil, errors.New("mpv transport is not configured")
	}
	return c.transport.Command(ctx, name, args...)
}

func (c *Client) propertyDuration(ctx context.Context, property string) (time.Duration, error) {
	value, err := c.command(ctx, "get_property", property)
	if err != nil {
		return 0, err
	}
	var seconds float64
	if err := json.Unmarshal(value, &seconds); err != nil {
		return 0, fmt.Errorf("decode %s: %w", property, err)
	}
	return time.Duration(seconds * float64(time.Second)), nil
}

func (c *Client) setState(state PlaybackState) {
	c.mu.Lock()
	c.state = state
	c.mu.Unlock()
}

type ProcessStarter interface {
	Start(ctx context.Context, socketPath string) (Process, error)
}
type Process interface{ Stop() error }

type MpvProcessStarter struct{ Binary string }

func (s MpvProcessStarter) Start(ctx context.Context, socketPath string) (Process, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	binary := s.Binary
	if binary == "" {
		binary = "mpv"
	}
	if err := os.RemoveAll(socketPath); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(filepath.Dir(socketPath), 0700); err != nil {
		return nil, err
	}
	// The context governs startup/connect time only. Runtime lifetime is
	// controlled explicitly by Process.Stop so callers can cancel a short
	// Open timeout without killing a successfully started player.
	command := exec.Command(binary, "--idle=yes", "--no-terminal", "--input-ipc-server="+socketPath)
	if err := command.Start(); err != nil {
		return nil, fmt.Errorf("start mpv: %w", err)
	}
	return &mpvProcess{command: command}, nil
}

type mpvProcess struct{ command *exec.Cmd }

func (p *mpvProcess) Stop() error {
	if p.command.Process == nil {
		return nil
	}
	killErr := p.command.Process.Kill()
	waitErr := p.command.Wait()
	if killErr != nil && !errors.Is(killErr, os.ErrProcessDone) {
		return killErr
	}
	if waitErr != nil && !errors.Is(waitErr, os.ErrProcessDone) {
		return waitErr
	}
	return nil
}

type IPCTransport struct {
	connection net.Conn
	reader     *bufio.Reader
	mu         sync.Mutex
	nextID     int64
}

func NewIPCTransport(ctx context.Context, socketPath string) (*IPCTransport, error) {
	dialer := net.Dialer{}
	connection, err := dialer.DialContext(ctx, "unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("connect mpv IPC: %w", err)
	}
	return &IPCTransport{connection: connection, reader: bufio.NewReader(connection)}, nil
}

func (t *IPCTransport) Command(ctx context.Context, command string, args ...any) (json.RawMessage, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.nextID++
	request := map[string]any{"command": append([]any{command}, args...), "request_id": t.nextID}
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	if deadline, ok := ctx.Deadline(); ok {
		_ = t.connection.SetDeadline(deadline)
	} else {
		_ = t.connection.SetDeadline(time.Time{})
	}
	if _, err = t.connection.Write(append(payload, '\n')); err != nil {
		return nil, err
	}
	for {
		line, readErr := t.reader.ReadBytes('\n')
		if readErr != nil {
			return nil, readErr
		}
		var response struct {
			Error     string          `json:"error"`
			Data      json.RawMessage `json:"data"`
			RequestID int64           `json:"request_id"`
		}
		if err := json.Unmarshal(line, &response); err != nil {
			continue
		}
		if response.RequestID != t.nextID {
			continue
		}
		if response.Error != "success" {
			return nil, errors.New(response.Error)
		}
		return response.Data, nil
	}
}

func (t *IPCTransport) Close() error { return t.connection.Close() }
