//go:build sqlite_fts5

package state

import (
	"context"
	"path/filepath"
	"testing"
)

func TestDebianBuildEnablesFTS5(t *testing.T) {
	store, err := Open(filepath.Join(t.TempDir(), "state.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	var tableType string
	if err := store.database.QueryRowContext(context.Background(), `SELECT type FROM sqlite_master WHERE name='search_fts'`).Scan(&tableType); err != nil {
		t.Fatal(err)
	}
	if tableType != "table" {
		t.Fatalf("search_fts type = %q, want virtual table", tableType)
	}
}
