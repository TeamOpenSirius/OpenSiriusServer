package repositories

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"

	"ymst_server/apiwire"
	"ymst_server/core/database"
)

type UserStateMutator struct {
	repo *UserRepository
}

func (r *UserRepository) Mutator() *UserStateMutator {
	return &UserStateMutator{repo: r}
}

func (m *UserStateMutator) UpdateObjectByUnion(ctx context.Context, unionID int64, match func(data []any) bool, mutate func(data []any) ([]any, error)) ([]any, error) {
	if m == nil || m.repo == nil {
		return nil, fmt.Errorf("user state mutator not initialized")
	}
	if err := m.repo.EnsureState(ctx); err != nil {
		return nil, err
	}
	var update any
	err := m.repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var rows []database.GameRawDataObject
		if err := tx.Where("user_key = ? AND union_id = ? AND is_null = ?", m.repo.userKey, unionID, false).Order("object_index ASC").Find(&rows).Error; err != nil {
			return err
		}
		for _, row := range rows {
			data, err := rawRowData(row)
			if err != nil {
				return err
			}
			if match != nil && !match(data) {
				continue
			}
			next, err := mutate(data)
			if err != nil {
				return err
			}
			if err := m.repo.upsertRawObjectWithKey(tx, row.ObjectKey, row.ObjectIndex, unionID, next); err != nil {
				return err
			}
			update = apiwire.DataObject(unionID, next)
			return nil
		}
		return gorm.ErrRecordNotFound
	})
	if err != nil {
		return nil, err
	}
	return []any{update}, nil
}

func (m *UserStateMutator) InsertObject(ctx context.Context, unionID int64, data []any) (any, error) {
	if m == nil || m.repo == nil {
		return nil, fmt.Errorf("user state mutator not initialized")
	}
	if err := m.repo.EnsureState(ctx); err != nil {
		return nil, err
	}
	index, err := m.repo.nextRawObjectIndex(ctx)
	if err != nil {
		return nil, err
	}
	if err := m.repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return m.repo.UpsertRawObject(tx, index, unionID, data)
	}); err != nil {
		return nil, err
	}
	return apiwire.DataObject(unionID, data), nil
}

func (r *UserRepository) touchStateMeta(tx *gorm.DB, objectDelta int) error {
	now := time.Now().UTC()
	var meta database.GameUserStateMeta
	err := tx.Where("user_key = ?", r.userKey).First(&meta).Error
	if err == nil {
		updates := map[string]any{"last_updated_at": now}
		if objectDelta != 0 {
			updates["object_count"] = meta.ObjectCount + objectDelta
		}
		return tx.Model(&meta).Updates(updates).Error
	}
	if err != gorm.ErrRecordNotFound {
		return err
	}
	seedMode := r.seedMode
	if seedMode == "" {
		seedMode = "generated"
	}
	return tx.Create(&database.GameUserStateMeta{UserKey: r.userKey, SeedMode: seedMode, StateVersion: 1, ObjectCount: objectDelta, LastBuiltAt: now, LastUpdatedAt: now}).Error
}
