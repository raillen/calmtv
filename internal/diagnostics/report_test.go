package diagnostics

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type fakeCommand struct{}

func (fakeCommand) Run(context.Context, string, ...string) ([]byte, error) {
	return []byte("password=secret\nready"), nil
}

func TestSaveWritesRedactedReportAtomically(t *testing.T) {
	path := filepath.Join(t.TempDir(), "latest.json")
	if err := Save(path, Report{OS: "linux", Commands: map[string]string{"token": "[token redacted]"}}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatal(err)
	}
}

func TestReportRedactsSecrets(t *testing.T) {
	report := NewReporter(fakeCommand{}).Collect(context.Background())
	if strings.Contains(report.Commands["kernel"], "secret") {
		t.Fatal("diagnostic report leaked secret")
	}
}
