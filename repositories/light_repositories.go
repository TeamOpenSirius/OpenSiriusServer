package repositories

import (
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"ymst_server/core/database"
)

type InboxRepository struct{ db *gorm.DB }

func NewInboxRepository(db *gorm.DB) *InboxRepository { return &InboxRepository{db: db} }
func (r *InboxRepository) HasUnreceived(ctx context.Context, userKey string, now time.Time) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&database.GameInboxPackage{}).Where("user_key = ? AND received_at IS NULL AND (expires_at IS NULL OR expires_at > ?)", userKey, now).Count(&count).Error
	return count > 0, err
}
func (r *InboxRepository) ListUnreceived(ctx context.Context, userKey string, now time.Time) ([]database.GameInboxPackage, error) {
	var rows []database.GameInboxPackage
	err := r.db.WithContext(ctx).Where("user_key = ? AND received_at IS NULL AND (expires_at IS NULL OR expires_at > ?)", userKey, now).Order("created_at ASC").Find(&rows).Error
	return rows, err
}
func (r *InboxRepository) FindUnreceived(ctx context.Context, userKey string, packageKey string, now time.Time) (*database.GameInboxPackage, error) {
	var row database.GameInboxPackage
	err := r.db.WithContext(ctx).Where("user_key = ? AND package_key = ? AND received_at IS NULL AND (expires_at IS NULL OR expires_at > ?)", userKey, packageKey, now).First(&row).Error
	if err != nil {
		return nil, err
	}
	return &row, nil
}
func (r *InboxRepository) MarkReceived(ctx context.Context, id uint, at time.Time) error {
	return r.db.WithContext(ctx).Model(&database.GameInboxPackage{}).Where("id = ?", id).Update("received_at", &at).Error
}

type FriendRepository struct{ db *gorm.DB }

func NewFriendRepository(db *gorm.DB) *FriendRepository { return &FriendRepository{db: db} }
func (r *FriendRepository) PendingReceived(ctx context.Context, userKey string) ([]database.GameFriendRequest, error) {
	var rows []database.GameFriendRequest
	err := r.db.WithContext(ctx).Where("to_user_key = ? AND status = ?", userKey, 0).Order("created_at DESC").Find(&rows).Error
	return rows, err
}
func (r *FriendRepository) AcceptedCount(ctx context.Context, userKey string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&database.GameFriendRelation{}).Where("user_key = ?", userKey).Count(&count).Error
	return count, err
}
func (r *FriendRepository) CreateRequest(ctx context.Context, requestKey string, fromUserKey string, toUserKey string, message string) error {
	return r.db.WithContext(ctx).Create(&database.GameFriendRequest{RequestKey: requestKey, FromUserKey: fromUserKey, ToUserKey: toUserKey, Message: message, Status: 0}).Error
}
func (r *FriendRepository) FindRequest(ctx context.Context, requestKey string) (*database.GameFriendRequest, error) {
	var row database.GameFriendRequest
	err := r.db.WithContext(ctx).Where("request_key = ?", requestKey).First(&row).Error
	if err != nil {
		return nil, err
	}
	return &row, nil
}
func (r *FriendRepository) UpdateRequestStatus(ctx context.Context, requestKey string, status int) error {
	return r.db.WithContext(ctx).Model(&database.GameFriendRequest{}).Where("request_key = ?", requestKey).Update("status", status).Error
}
func (r *FriendRepository) CreateRelation(ctx context.Context, userKey string, friendUserKey string) error {
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&database.GameFriendRelation{UserKey: userKey, FriendUserKey: friendUserKey}).
		Error
}
func (r *FriendRepository) DeleteRelation(ctx context.Context, userKey string, friendUserKey string) error {
	return r.db.WithContext(ctx).Where("(user_key = ? AND friend_user_key = ?) OR (user_key = ? AND friend_user_key = ?)", userKey, friendUserKey, friendUserKey, userKey).Delete(&database.GameFriendRelation{}).Error
}

type CircleRepository struct{ db *gorm.DB }

func NewCircleRepository(db *gorm.DB) *CircleRepository { return &CircleRepository{db: db} }
func (r *CircleRepository) PendingInvites(ctx context.Context, userKey string) ([]database.GameCircleInvite, error) {
	var rows []database.GameCircleInvite
	err := r.db.WithContext(ctx).Where("user_key = ? AND status = ?", userKey, 0).Order("created_at DESC").Find(&rows).Error
	return rows, err
}
func (r *CircleRepository) CreateCircle(ctx context.Context, circleID int64, userKey string, name string, description string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&database.GameCircle{CircleID: circleID, CircleName: name, OwnerUserKey: userKey, Description: description}).Error; err != nil {
			return err
		}
		return tx.Create(&database.GameCircleMember{CircleID: circleID, UserKey: userKey, Role: 1, JoinedAt: time.Now().UTC()}).Error
	})
}
func (r *CircleRepository) FindInvite(ctx context.Context, inviteKey string) (*database.GameCircleInvite, error) {
	var row database.GameCircleInvite
	err := r.db.WithContext(ctx).Where("invite_key = ?", inviteKey).First(&row).Error
	if err != nil {
		return nil, err
	}
	return &row, nil
}
func (r *CircleRepository) UpdateInviteStatus(ctx context.Context, inviteKey string, status int) error {
	return r.db.WithContext(ctx).Model(&database.GameCircleInvite{}).Where("invite_key = ?", inviteKey).Update("status", status).Error
}
func (r *CircleRepository) AddMember(ctx context.Context, circleID int64, userKey string, role int) error {
	return r.db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&database.GameCircleMember{CircleID: circleID, UserKey: userKey, Role: role, JoinedAt: time.Now().UTC()}).
		Error
}
func (r *CircleRepository) Leave(ctx context.Context, userKey string) error {
	return r.db.WithContext(ctx).Where("user_key = ?", userKey).Delete(&database.GameCircleMember{}).Error
}

type SupportCompanyRepository struct{ db *gorm.DB }

func NewSupportCompanyRepository(db *gorm.DB) *SupportCompanyRepository {
	return &SupportCompanyRepository{db: db}
}
func (r *SupportCompanyRepository) EnsureForUser(ctx context.Context, userKey string) ([]database.GameSupportCompanyState, error) {
	for company := 0; company < 4; company++ {
		var row database.GameSupportCompanyState
		err := r.db.WithContext(ctx).Where("user_key = ? AND company = ?", userKey, company).First(&row).Error
		if err == gorm.ErrRecordNotFound {
			if err := r.db.WithContext(ctx).Create(&database.GameSupportCompanyState{UserKey: userKey, Company: company, Level: 1}).Error; err != nil {
				return nil, err
			}
		} else if err != nil {
			return nil, err
		}
	}
	var rows []database.GameSupportCompanyState
	err := r.db.WithContext(ctx).Where("user_key = ?", userKey).Order("company ASC").Find(&rows).Error
	return rows, err
}
