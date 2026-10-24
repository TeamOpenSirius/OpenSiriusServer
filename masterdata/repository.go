package masterdata

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var ErrMasterDataUnavailable = errors.New("master data unavailable")
var ErrEpisodeNotFound = errors.New("episode not found")

type EpisodeMaster struct {
	EpisodeMasterID          int64
	Title                    string
	EpisodeTitle             string
	StoryType                int64
	EpisodeOrder             int64
	EpisodeDetailAssetSource string
}

type MasterDataRepository struct {
	db   *sql.DB
	root string
}

func Open(dbPath string, root string) *MasterDataRepository {
	repo := &MasterDataRepository{root: root}
	if strings.TrimSpace(dbPath) != "" {
		if _, err := os.Stat(dbPath); err == nil {
			if db, err := sql.Open("sqlite3", "file:"+filepath.ToSlash(dbPath)+"?mode=ro&_busy_timeout=5000"); err == nil {
				repo.db = db
			}
		}
	}
	return repo
}

func (r *MasterDataRepository) Close() error {
	if r != nil && r.db != nil {
		return r.db.Close()
	}
	return nil
}
func (r *MasterDataRepository) Episodes() *EpisodeMasterRepository {
	return &EpisodeMasterRepository{repo: r}
}

type EpisodeMasterRepository struct{ repo *MasterDataRepository }

func (r *EpisodeMasterRepository) FindEpisodeDetail(ctx context.Context, episodeMasterID int64) (*EpisodeMaster, error) {
	if r == nil || r.repo == nil {
		return nil, ErrMasterDataUnavailable
	}
	if r.repo.db != nil {
		queries := []string{
			`SELECT id, title, story_type, episode_order, episode_detail_asset_source FROM episode_master WHERE id = ?`,
			`SELECT episode_master_id, title, story_type, episode_order, episode_detail_asset_source FROM episode_master WHERE episode_master_id = ?`,
			`SELECT id, episode_title, story_type, episode_order, episode_detail_asset_source FROM episode_master WHERE id = ?`,
			`SELECT id, title, story_type, "order", episode_detail_asset_source FROM episode_masters WHERE id = ?`,
		}
		for _, q := range queries {
			ep := EpisodeMaster{}
			err := r.repo.db.QueryRowContext(ctx, q, episodeMasterID).Scan(&ep.EpisodeMasterID, &ep.Title, &ep.StoryType, &ep.EpisodeOrder, &ep.EpisodeDetailAssetSource)
			if err == nil {
				if ep.Title == "" {
					ep.Title = ep.EpisodeTitle
				}
				return &ep, nil
			}
			if err != sql.ErrNoRows && !strings.Contains(strings.ToLower(err.Error()), "no such table") && !strings.Contains(strings.ToLower(err.Error()), "no such column") {
				return nil, fmt.Errorf("query episode master: %w", err)
			}
		}
	}
	if ep, err := r.findEpisodeJSON(episodeMasterID); err == nil {
		return ep, nil
	}
	return nil, ErrEpisodeNotFound
}

func (r *EpisodeMasterRepository) findEpisodeJSON(episodeMasterID int64) (*EpisodeMaster, error) {
	if strings.TrimSpace(r.repo.root) == "" {
		return nil, ErrMasterDataUnavailable
	}
	candidates := []string{
		filepath.Join(r.repo.root, "json", "episode_master.json"),
		filepath.Join(r.repo.root, "json", "episode_masters.json"),
		filepath.Join(r.repo.root, "episode_master.json"),
		filepath.Join(r.repo.root, "episode_masters.json"),
	}
	for _, path := range candidates {
		body, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var rows []EpisodeMaster
		if err := json.Unmarshal(body, &rows); err == nil {
			for _, row := range rows {
				if row.EpisodeMasterID == 0 {
					continue
				}
				if row.EpisodeMasterID == episodeMasterID {
					if row.Title == "" {
						row.Title = row.EpisodeTitle
					}
					return &row, nil
				}
			}
		}
		var keyed map[string]EpisodeMaster
		if err := json.Unmarshal(body, &keyed); err == nil {
			key := fmt.Sprintf("%d", episodeMasterID)
			if row, ok := keyed[key]; ok {
				if row.EpisodeMasterID == 0 {
					row.EpisodeMasterID = episodeMasterID
				}
				if row.Title == "" {
					row.Title = row.EpisodeTitle
				}
				return &row, nil
			}
		}
	}
	return nil, ErrEpisodeNotFound
}
