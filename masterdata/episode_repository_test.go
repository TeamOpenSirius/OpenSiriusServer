package masterdata

import (
	"context"
	"database/sql"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestEpisodeRepositoryFindEpisodeDetail(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "master.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`CREATE TABLE episode_master (id INTEGER PRIMARY KEY, title TEXT, story_type INTEGER, episode_order INTEGER, episode_detail_asset_source TEXT)`); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec(`INSERT INTO episode_master (id, title, story_type, episode_order, episode_detail_asset_source) VALUES (140271, 'Title A', 2, 3, 'asset/path')`); err != nil {
		t.Fatal(err)
	}
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}

	repo := Open(dbPath, "")
	defer repo.Close()
	ep, err := repo.Episodes().FindEpisodeDetail(context.Background(), 140271)
	if err != nil {
		t.Fatal(err)
	}
	if ep.Title != "Title A" || ep.StoryType != 2 || ep.EpisodeOrder != 3 || ep.EpisodeDetailAssetSource != "asset/path" {
		t.Fatalf("unexpected episode: %#v", ep)
	}
}
