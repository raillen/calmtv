package recovery

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestResetUIRemovesOnlyUIConfiguration(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "ui"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "state.db"), []byte("keep"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := NewManager(root, func(context.Context, string, ...string) error { return nil }).Execute(context.Background(), ResetUI); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "ui")); !os.IsNotExist(err) {
		t.Fatal("UI configuration was not removed")
	}
	if _, err := os.Stat(filepath.Join(root, "state.db")); err != nil {
		t.Fatal("runtime state was removed")
	}
}
