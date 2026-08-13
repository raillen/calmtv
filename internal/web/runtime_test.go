package web

import (
	"context"
	"testing"
)

func TestLauncherRejectsNonHTTPSRuntimeURL(t *testing.T) {
	if _, err := NewLauncher("firefox", t.TempDir()).Start(context.Background(), "http://example.invalid"); err == nil {
		t.Fatal("non-HTTPS streaming URL accepted")
	}
}
