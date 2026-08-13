package downloads

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadUsesPartFileAndReportsCompletion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) { _, _ = writer.Write([]byte("data")) }))
	defer server.Close()
	root := t.TempDir()
	manager := NewManager(root)
	var done bool
	if err := manager.Download(context.Background(), server.URL, "clip.bin", func(progress Progress) { done = progress.Done }); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(filepath.Join(root, "clip.bin"))
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "data" || !done {
		t.Fatalf("content=%q done=%v", content, done)
	}
	if _, err := os.Stat(filepath.Join(root, "clip.bin.part")); !os.IsNotExist(err) {
		t.Fatal("partial file remains")
	}
}
