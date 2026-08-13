package library

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanRecognizesLocalMediaOnly(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "movie.mkv"), []byte("fixture"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "notes.txt"), []byte("ignore"), 0600); err != nil {
		t.Fatal(err)
	}
	items, err := Scan(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 || items[0].Kind != "video" {
		t.Fatalf("items = %#v", items)
	}
}
