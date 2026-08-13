package focus

import (
	"testing"

	"github.com/raillen/calmtv/internal/input"
)

func TestManagerSkipsDisabledNodesAndKeepsEdgesSafe(t *testing.T) {
	manager, err := NewManager([]Node{
		{ID: "media", Row: 0, Column: 0, Enabled: true},
		{ID: "settings", Row: 0, Column: 1, Enabled: false},
		{ID: "games", Row: 0, Column: 2, Enabled: true},
	}, "media")
	if err != nil {
		t.Fatal(err)
	}

	if got := manager.Move(input.NavRight); got != "games" {
		t.Fatalf("Move right = %q, want games", got)
	}
	if got := manager.Move(input.NavRight); got != "games" {
		t.Fatalf("edge move changed focus to %q", got)
	}
}

func TestManagerFallsBackToFirstEnabledTarget(t *testing.T) {
	manager, err := NewManager([]Node{
		{ID: "second", Row: 1, Column: 0, Enabled: true},
		{ID: "first", Row: 0, Column: 0, Enabled: true},
	}, "missing")
	if err != nil {
		t.Fatal(err)
	}
	if manager.Current() != "first" {
		t.Fatalf("initial focus = %q, want first", manager.Current())
	}
}
