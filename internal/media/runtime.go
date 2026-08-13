package media

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type TransportConnector func(context.Context, string) (Transport, error)

type Runtime struct {
	starter ProcessStarter
	connect TransportConnector
	process Process
	client  *Client
	socket  string
}

func NewRuntime(starter ProcessStarter, connect TransportConnector, socket string) *Runtime {
	return &Runtime{starter: starter, connect: connect, socket: socket}
}

func DefaultRuntime(socket string) *Runtime {
	return NewRuntime(MpvProcessStarter{}, func(ctx context.Context, path string) (Transport, error) { return NewIPCTransport(ctx, path) }, socket)
}

func (r *Runtime) Open(ctx context.Context, path string) error {
	if path == "" || filepath.IsAbs(path) == false {
		return fmt.Errorf("media path must be absolute")
	}
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("media file: %w", err)
	}
	return r.openTarget(ctx, path)
}

// OpenURL opens a network media source without making the Shell inspect or
// download it. mpv owns the protocol handling and any required buffering.
func (r *Runtime) OpenURL(ctx context.Context, url string) error {
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return fmt.Errorf("media URL must use http or https")
	}
	return r.openTarget(ctx, url)
}

func (r *Runtime) openTarget(ctx context.Context, target string) error {
	if err := r.Stop(); err != nil {
		return err
	}
	process, err := r.starter.Start(ctx, r.socket)
	if err != nil {
		return err
	}
	transport, err := waitForTransport(ctx, r.connect, r.socket)
	if err != nil {
		_ = process.Stop()
		return err
	}
	r.process, r.client = process, NewClient(transport)
	return r.client.Open(ctx, target)
}

func (r *Runtime) Client() *Client { return r.client }

func (r *Runtime) Stop() error {
	if r.client != nil && r.client.transport != nil {
		_ = r.client.transport.Close()
	}
	r.client = nil
	if r.process == nil {
		return nil
	}
	err := r.process.Stop()
	r.process = nil
	return err
}

func waitForTransport(ctx context.Context, connect TransportConnector, socket string) (Transport, error) {
	deadline := time.NewTicker(25 * time.Millisecond)
	defer deadline.Stop()
	for {
		transport, err := connect(ctx, socket)
		if err == nil {
			return transport, nil
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-deadline.C:
		}
	}
}
