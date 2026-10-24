package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"ymst_server/apiwire"
	"ymst_server/masterdata"

	"gorm.io/gorm"
)

type EpisodeService struct {
	master  *masterdata.MasterDataRepository
	db      *gorm.DB
	baseURL string
}

func NewEpisodeService(master *masterdata.MasterDataRepository, db *gorm.DB, baseURL string) *EpisodeService {
	return &EpisodeService{master: master, db: db, baseURL: strings.TrimRight(strings.TrimSpace(baseURL), "/")}
}
func (s *EpisodeService) GetEpisodeDetail(ctx context.Context, episodeMasterID int64) ([]any, error) {
	if s != nil && s.db != nil {
		if ep, err := s.findImportedEpisodeDetail(ctx, episodeMasterID); err == nil {
			return apiwire.EpisodeResult(ep.Title, ep.StoryType, ep.EpisodeOrder, s.episodeDetailAssetSource(episodeMasterID)), nil
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	if s != nil && s.master != nil {
		if ep, err := s.master.Episodes().FindEpisodeDetail(ctx, episodeMasterID); err == nil {
			source := strings.TrimSpace(ep.EpisodeDetailAssetSource)
			if source == "" {
				source = s.episodeDetailAssetSource(episodeMasterID)
			}
			source = s.normalizeEpisodeDetailAssetSource(source)
			return apiwire.EpisodeResult(ep.Title, ep.StoryType, ep.EpisodeOrder, source), nil
		}
	}
	return apiwire.EpisodeResult("Episode", 0, 0, s.episodeDetailAssetSource(episodeMasterID)), nil
}

func (s *EpisodeService) GetEpisodeDetailAsset(ctx context.Context, episodeMasterID int64) ([]any, error) {
	title := "Episode"
	if s != nil {
		if ep, err := s.findImportedEpisodeDetail(ctx, episodeMasterID); err == nil && strings.TrimSpace(ep.Title) != "" {
			title = ep.Title
		}
	}
	return []any{episodeDetailResultRow(episodeMasterID, title)}, nil
}

type importedEpisodeDetail struct {
	Title        string
	StoryType    int64
	EpisodeOrder int64
}

func (s *EpisodeService) findImportedEpisodeDetail(ctx context.Context, episodeMasterID int64) (importedEpisodeDetail, error) {
	if s == nil || s.db == nil {
		return importedEpisodeDetail{}, gorm.ErrRecordNotFound
	}
	if !s.db.Migrator().HasTable("master_episode_master") && !s.db.Migrator().HasTable("master_character_episode_master") {
		return importedEpisodeDetail{}, gorm.ErrRecordNotFound
	}
	if ep, err := s.findImportedMainEpisode(ctx, episodeMasterID); err == nil {
		return ep, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return importedEpisodeDetail{}, err
	}
	if ep, err := s.findImportedCharacterEpisode(ctx, episodeMasterID); err == nil {
		return ep, nil
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return importedEpisodeDetail{}, err
	}
	return importedEpisodeDetail{}, gorm.ErrRecordNotFound
}

func (s *EpisodeService) findImportedMainEpisode(ctx context.Context, episodeMasterID int64) (importedEpisodeDetail, error) {
	if !s.db.Migrator().HasTable("master_episode_master") {
		return importedEpisodeDetail{}, gorm.ErrRecordNotFound
	}
	var row struct {
		Title     string
		StoryType string
		Order     int64
	}
	query := `
SELECT e.title AS title, '' AS story_type, e.` + "`order`" + ` AS ` + "`order`" + `
FROM master_episode_master e
WHERE e.id = ?
LIMIT 1`
	if s.db.Migrator().HasTable("master_story_master") {
		query = `
SELECT e.title AS title, COALESCE(st.type, '') AS story_type, e.`+"`order`"+` AS `+"`order`"+`
FROM master_episode_master e
LEFT JOIN master_story_master st ON st.id = e.story_master_id
WHERE e.id = ?
LIMIT 1`
	}
	err := s.db.WithContext(ctx).Raw(query, episodeMasterID).Scan(&row).Error
	if err != nil {
		return importedEpisodeDetail{}, err
	}
	storyType := storyTypeValue(row.StoryType)
	if row.Title == "" && storyType == 0 && row.Order == 0 {
		return importedEpisodeDetail{}, gorm.ErrRecordNotFound
	}
	return importedEpisodeDetail{Title: row.Title, StoryType: storyType, EpisodeOrder: row.Order}, nil
}

func (s *EpisodeService) findImportedCharacterEpisode(ctx context.Context, episodeMasterID int64) (importedEpisodeDetail, error) {
	if !s.db.Migrator().HasTable("master_character_episode_master") {
		return importedEpisodeDetail{}, gorm.ErrRecordNotFound
	}
	var row struct {
		CharacterName string
		EpisodeOrder  string
	}
	query := `
SELECT '' AS character_name, ce.episode_order AS episode_order
FROM master_character_episode_master ce
WHERE ce.episode_master_id = ? OR ce.id = ?
LIMIT 1`
	if s.db.Migrator().HasTable("master_character_master") {
		query = `
SELECT COALESCE(c.name, '') AS character_name, ce.episode_order AS episode_order
FROM master_character_episode_master ce
LEFT JOIN master_character_master c ON c.id = ce.character_master_id
WHERE ce.episode_master_id = ? OR ce.id = ?
LIMIT 1`
	}
	err := s.db.WithContext(ctx).Raw(query, episodeMasterID, episodeMasterID).Scan(&row).Error
	if err != nil {
		return importedEpisodeDetail{}, err
	}
	order := characterEpisodeOrderValue(row.EpisodeOrder)
	if row.CharacterName == "" && order == 0 {
		return importedEpisodeDetail{}, gorm.ErrRecordNotFound
	}
	title := strings.TrimSpace(row.CharacterName)
	if title == "" {
		title = "Episode"
	}
	if order > 0 {
		title = fmt.Sprintf("%s Episode %d", title, order)
	}
	return importedEpisodeDetail{Title: title, StoryType: 4, EpisodeOrder: order}, nil
}

func characterEpisodeOrderValue(value string) int64 {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "first", "1":
		return 1
	case "second", "2":
		return 2
	default:
		return 0
	}
}

func storyTypeValue(value string) int64 {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "main", "1":
		return 1
	case "event", "2":
		return 2
	case "side", "3":
		return 3
	case "character", "4":
		return 4
	case "special", "5":
		return 5
	default:
		return 0
	}
}

func (s *EpisodeService) episodeDetailAssetSource(episodeMasterID int64) string {
	return fmt.Sprintf("episode-details/%d", episodeMasterID)
}

func (s *EpisodeService) normalizeEpisodeDetailAssetSource(source string) string {
	source = strings.TrimSpace(source)
	if source == "" {
		return source
	}
	if s != nil && s.baseURL != "" {
		source = strings.TrimPrefix(source, s.baseURL)
	}
	return strings.TrimLeft(source, "/")
}

func episodeDetailResultRow(episodeMasterID int64, title string) []any {
	return apiwire.ArrayByKey(24, map[int]any{
		0:  episodeMasterID*1000 + 1,
		1:  episodeMasterID,
		2:  int32(1),
		3:  int32(1),
		4:  "",
		5:  "",
		6:  int32(0),
		7:  title,
		8:  title,
		9:  "",
		10: "",
		11: nil,
		12: "",
		13: "",
		14: "",
		15: "",
		16: nil,
		17: nil,
		18: "",
		19: []any{},
		20: "",
		21: nil,
		22: nil,
		23: nil,
	})
}
