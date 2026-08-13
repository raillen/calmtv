package state

import (
	"context"
	"path/filepath"
	"testing"
	"time"
)

func TestStoreMigratesAndRestoresAppState(t *testing.T) {
	store, err := Open(filepath.Join(t.TempDir(), "state.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	if err := store.SaveAppState(context.Background(), "local-media", AppState{Route: "library", SelectedID: "movie-1", PositionSec: 42}); err != nil {
		t.Fatal(err)
	}
	value, err := store.RestoreAppState(context.Background(), "local-media")
	if err != nil {
		t.Fatal(err)
	}
	if value.Route != "library" || value.PositionSec != 42 {
		t.Fatalf("state = %#v", value)
	}
}

func TestStorePersistsMediaAndFavorites(t *testing.T) {
	store, err := Open(filepath.Join(t.TempDir(), "state.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	ctx := context.Background()
	if err := store.SaveMediaItem(ctx, MediaItem{ID: "movie-1", Path: "/Videos/movie.mkv", Title: "Movie", Kind: "video", Duration: 90 * time.Second}); err != nil {
		t.Fatal(err)
	}
	if err := store.SaveFavorite(ctx, "movie-1", "media", true); err != nil {
		t.Fatal(err)
	}
	if err := store.SaveFavorite(ctx, "movie-1", "media", false); err != nil {
		t.Fatal(err)
	}
	if err := store.SaveProgress(ctx, "movie-1", 30*time.Second, 90*time.Second); err != nil {
		t.Fatal(err)
	}
	position, duration, err := store.RestoreProgress(ctx, "movie-1")
	if err != nil || position != 30*time.Second || duration != 90*time.Second {
		t.Fatalf("position=%s duration=%s err=%v", position, duration, err)
	}
	if err := store.SaveDownload(ctx, Download{ID: "download-1", URL: "https://example.invalid/file", Destination: "/Downloads/file", State: "complete"}); err != nil {
		t.Fatal(err)
	}
	if err := store.SaveGame(ctx, Game{ID: "game-1", Path: "/Games/test.nes", Hash: "hash", Title: "Test", System: "nes", Core: "mesen"}); err != nil {
		t.Fatal(err)
	}
}
