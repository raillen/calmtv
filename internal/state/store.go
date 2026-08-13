package state

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct{ database *sql.DB }

func Open(path string) (*Store, error) {
	database, err := sql.Open("sqlite3", path+"?_busy_timeout=5000&_foreign_keys=on")
	if err != nil {
		return nil, fmt.Errorf("open state database: %w", err)
	}
	store := &Store{database: database}
	if err := store.Migrate(context.Background()); err != nil {
		_ = database.Close()
		return nil, err
	}
	return store, nil
}

func (s *Store) Close() error { return s.database.Close() }

func (s *Store) Migrate(ctx context.Context) error {
	_, err := s.database.ExecContext(ctx, `
CREATE TABLE IF NOT EXISTS schema_migrations (version INTEGER PRIMARY KEY, applied_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS settings (key TEXT PRIMARY KEY, value TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS app_state (app_id TEXT PRIMARY KEY, route TEXT NOT NULL DEFAULT '', selected_id TEXT NOT NULL DEFAULT '', position_sec INTEGER NOT NULL DEFAULT 0, updated_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS media_items (id TEXT PRIMARY KEY, path TEXT NOT NULL UNIQUE, title TEXT NOT NULL, kind TEXT NOT NULL, duration_sec INTEGER NOT NULL DEFAULT 0, updated_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS playback_progress (media_id TEXT PRIMARY KEY REFERENCES media_items(id) ON DELETE CASCADE, position_sec INTEGER NOT NULL DEFAULT 0, duration_sec INTEGER NOT NULL DEFAULT 0, updated_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS favorites (item_id TEXT PRIMARY KEY, item_kind TEXT NOT NULL, created_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS downloads (id TEXT PRIMARY KEY, url TEXT NOT NULL, destination TEXT NOT NULL, received_bytes INTEGER NOT NULL DEFAULT 0, total_bytes INTEGER NOT NULL DEFAULT 0, state TEXT NOT NULL, updated_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS library_sources (path TEXT PRIMARY KEY, kind TEXT NOT NULL, enabled INTEGER NOT NULL DEFAULT 1);
CREATE TABLE IF NOT EXISTS channels (id TEXT PRIMARY KEY, name TEXT NOT NULL, group_name TEXT NOT NULL DEFAULT '', logo TEXT NOT NULL DEFAULT '', stream_url TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS epg_programs (channel_id TEXT NOT NULL, start_at TEXT NOT NULL, end_at TEXT NOT NULL, title TEXT NOT NULL, description TEXT NOT NULL DEFAULT '', PRIMARY KEY(channel_id, start_at));
CREATE TABLE IF NOT EXISTS games (id TEXT PRIMARY KEY, path TEXT NOT NULL UNIQUE, hash TEXT NOT NULL, title TEXT NOT NULL, system TEXT NOT NULL, core TEXT NOT NULL, updated_at TEXT NOT NULL);
CREATE TABLE IF NOT EXISTS game_sessions (game_id TEXT PRIMARY KEY REFERENCES games(id) ON DELETE CASCADE, save_path TEXT NOT NULL DEFAULT '', updated_at TEXT NOT NULL);
`)
	if err != nil {
		return fmt.Errorf("migrate state database: %w", err)
	}
	if _, err := s.database.ExecContext(ctx, `CREATE VIRTUAL TABLE IF NOT EXISTS search_fts USING fts5(item_id UNINDEXED, title, aliases)`); err != nil {
		// Debian builds enable SQLite FTS5 explicitly. Keep development hosts
		// without that compile option usable with a bounded fallback table.
		if _, fallbackErr := s.database.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS search_fts (item_id TEXT PRIMARY KEY, title TEXT NOT NULL, aliases TEXT NOT NULL DEFAULT '')`); fallbackErr != nil {
			return fmt.Errorf("create search index: %w", fallbackErr)
		}
	}
	_, err = s.database.ExecContext(ctx, `INSERT OR IGNORE INTO schema_migrations(version, applied_at) VALUES (1, ?)`, time.Now().UTC().Format(time.RFC3339))
	return err
}

type AppState struct {
	Route, SelectedID string
	PositionSec       int64
}

type MediaItem struct {
	ID, Path, Title, Kind string
	Duration              time.Duration
}

type Download struct {
	ID, URL, Destination, State string
	Received, Total             int64
}

func (s *Store) SaveDownload(ctx context.Context, download Download) error {
	_, err := s.database.ExecContext(ctx, `INSERT INTO downloads(id, url, destination, received_bytes, total_bytes, state, updated_at) VALUES(?,?,?,?,?,?,?) ON CONFLICT(id) DO UPDATE SET url=excluded.url, destination=excluded.destination, received_bytes=excluded.received_bytes, total_bytes=excluded.total_bytes, state=excluded.state, updated_at=excluded.updated_at`, download.ID, download.URL, download.Destination, download.Received, download.Total, download.State, time.Now().UTC().Format(time.RFC3339))
	return err
}

type Game struct {
	ID, Path, Hash, Title, System, Core string
}

func (s *Store) SaveGame(ctx context.Context, game Game) error {
	_, err := s.database.ExecContext(ctx, `INSERT INTO games(id, path, hash, title, system, core, updated_at) VALUES(?,?,?,?,?,?,?) ON CONFLICT(id) DO UPDATE SET path=excluded.path, hash=excluded.hash, title=excluded.title, system=excluded.system, core=excluded.core, updated_at=excluded.updated_at`, game.ID, game.Path, game.Hash, game.Title, game.System, game.Core, time.Now().UTC().Format(time.RFC3339))
	return err
}

func (s *Store) SaveMediaItem(ctx context.Context, item MediaItem) error {
	_, err := s.database.ExecContext(ctx, `INSERT INTO media_items(id, path, title, kind, duration_sec, updated_at) VALUES(?,?,?,?,?,?) ON CONFLICT(id) DO UPDATE SET path=excluded.path, title=excluded.title, kind=excluded.kind, duration_sec=excluded.duration_sec, updated_at=excluded.updated_at`, item.ID, item.Path, item.Title, item.Kind, int64(item.Duration.Seconds()), time.Now().UTC().Format(time.RFC3339))
	return err
}

func (s *Store) SaveFavorite(ctx context.Context, itemID, kind string, favorite bool) error {
	if favorite {
		_, err := s.database.ExecContext(ctx, `INSERT OR REPLACE INTO favorites(item_id, item_kind, created_at) VALUES(?,?,?)`, itemID, kind, time.Now().UTC().Format(time.RFC3339))
		return err
	}
	_, err := s.database.ExecContext(ctx, `DELETE FROM favorites WHERE item_id=?`, itemID)
	return err
}

func (s *Store) SaveAppState(ctx context.Context, appID string, value AppState) error {
	_, err := s.database.ExecContext(ctx, `INSERT INTO app_state(app_id, route, selected_id, position_sec, updated_at) VALUES(?,?,?,?,?) ON CONFLICT(app_id) DO UPDATE SET route=excluded.route, selected_id=excluded.selected_id, position_sec=excluded.position_sec, updated_at=excluded.updated_at`, appID, value.Route, value.SelectedID, value.PositionSec, time.Now().UTC().Format(time.RFC3339))
	return err
}

func (s *Store) RestoreAppState(ctx context.Context, appID string) (AppState, error) {
	var value AppState
	err := s.database.QueryRowContext(ctx, `SELECT route, selected_id, position_sec FROM app_state WHERE app_id=?`, appID).Scan(&value.Route, &value.SelectedID, &value.PositionSec)
	if err == sql.ErrNoRows {
		return AppState{}, nil
	}
	return value, err
}

func (s *Store) SaveProgress(ctx context.Context, mediaID string, position, duration time.Duration) error {
	_, err := s.database.ExecContext(ctx, `INSERT INTO playback_progress(media_id, position_sec, duration_sec, updated_at) VALUES(?,?,?,?) ON CONFLICT(media_id) DO UPDATE SET position_sec=excluded.position_sec, duration_sec=excluded.duration_sec, updated_at=excluded.updated_at`, mediaID, int64(position.Seconds()), int64(duration.Seconds()), time.Now().UTC().Format(time.RFC3339))
	return err
}

func (s *Store) RestoreProgress(ctx context.Context, mediaID string) (time.Duration, time.Duration, error) {
	var position, duration int64
	err := s.database.QueryRowContext(ctx, `SELECT position_sec, duration_sec FROM playback_progress WHERE media_id=?`, mediaID).Scan(&position, &duration)
	if err == sql.ErrNoRows {
		return 0, 0, nil
	}
	if err != nil {
		return 0, 0, err
	}
	return time.Duration(position) * time.Second, time.Duration(duration) * time.Second, nil
}

func (s *Store) SaveChannel(ctx context.Context, id, name, group, logo, streamURL string) error {
	_, err := s.database.ExecContext(ctx, `INSERT INTO channels(id, name, group_name, logo, stream_url) VALUES(?,?,?,?,?) ON CONFLICT(id) DO UPDATE SET name=excluded.name, group_name=excluded.group_name, logo=excluded.logo, stream_url=excluded.stream_url`, id, name, group, logo, streamURL)
	return err
}

func (s *Store) SaveEPGProgram(ctx context.Context, channelID string, start, end time.Time, title, description string) error {
	_, err := s.database.ExecContext(ctx, `INSERT INTO epg_programs(channel_id, start_at, end_at, title, description) VALUES(?,?,?,?,?) ON CONFLICT(channel_id, start_at) DO UPDATE SET end_at=excluded.end_at, title=excluded.title, description=excluded.description`, channelID, start.UTC().Format(time.RFC3339), end.UTC().Format(time.RFC3339), title, description)
	return err
}
