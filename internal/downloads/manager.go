package downloads

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Progress struct {
	URL, Destination string
	Received, Total  int64
	Done             bool
	Err              error
}

type Manager struct {
	client *http.Client
	root   string
}

func NewManager(root string) *Manager {
	return &Manager{client: &http.Client{Timeout: 30 * time.Second}, root: root}
}

func (m *Manager) Download(ctx context.Context, rawURL, name string, onProgress func(Progress)) error {
	if !strings.HasPrefix(rawURL, "https://") && !strings.HasPrefix(rawURL, "http://") {
		return fmt.Errorf("download URL must use HTTP(S)")
	}
	if name == "" || filepath.Base(name) != name || strings.ContainsAny(name, `/\\`) {
		return fmt.Errorf("invalid download name")
	}
	if err := os.MkdirAll(m.root, 0750); err != nil {
		return err
	}
	destination := filepath.Join(m.root, name)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return err
	}
	response, err := m.client.Do(request)
	if err != nil {
		return fmt.Errorf("download request: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("download returned HTTP %d", response.StatusCode)
	}
	file, err := os.OpenFile(destination+".part", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	progress := Progress{URL: rawURL, Destination: destination, Total: response.ContentLength}
	buffer := make([]byte, 32*1024)
	for {
		read, readErr := response.Body.Read(buffer)
		if read > 0 {
			written, writeErr := file.Write(buffer[:read])
			if writeErr != nil {
				_ = file.Close()
				return writeErr
			}
			progress.Received += int64(written)
			if onProgress != nil {
				onProgress(progress)
			}
		}
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			_ = file.Close()
			return fmt.Errorf("read download: %w", readErr)
		}
	}
	if err := file.Close(); err != nil {
		return err
	}
	if err := os.Rename(destination+".part", destination); err != nil {
		return err
	}
	progress.Done = true
	if onProgress != nil {
		onProgress(progress)
	}
	return nil
}
