package services

import (
	"context"
	"strconv"
	"time"

	"ymst_server/apiwire"
	"ymst_server/repositories"
)

type InboxService struct{ inbox *repositories.InboxRepository }

func NewInboxService(inbox *repositories.InboxRepository) *InboxService {
	return &InboxService{inbox: inbox}
}
func (s *InboxService) CheckPackages(ctx context.Context, userKey string) (bool, error) {
	if s == nil || s.inbox == nil {
		return false, nil
	}
	return s.inbox.HasUnreceived(ctx, userKey, time.Now().UTC())
}
func (s *InboxService) GetPackages(ctx context.Context, userKey string) ([]any, error) {
	if s == nil || s.inbox == nil {
		return []any{}, nil
	}
	_, err := s.inbox.ListUnreceived(ctx, userKey, time.Now().UTC())
	if err != nil {
		return nil, err
	}
	return []any{}, nil
}
func (s *InboxService) ReceivePackage(ctx context.Context, userKey string, packageKey string) ([]any, []any, error) {
	if s == nil || s.inbox == nil || packageKey == "" {
		return apiwire.InboxReceiveResult([]any{}, false), []any{}, nil
	}
	row, err := s.inbox.FindUnreceived(ctx, userKey, packageKey, time.Now().UTC())
	if err != nil {
		return apiwire.InboxReceiveResult([]any{}, false), []any{}, nil
	}
	if err := s.inbox.MarkReceived(ctx, row.ID, time.Now().UTC()); err != nil {
		return nil, nil, err
	}
	hasRemaining, err := s.inbox.HasUnreceived(ctx, userKey, time.Now().UTC())
	if err != nil {
		return nil, nil, err
	}
	return apiwire.InboxReceiveResult([]any{}, hasRemaining), []any{}, nil
}
func (s *InboxService) ReceiveAllPackages(ctx context.Context, userKey string) ([]any, []any, error) {
	rows, err := s.inbox.ListUnreceived(ctx, userKey, time.Now().UTC())
	if err != nil {
		return nil, nil, err
	}
	for _, row := range rows {
		if err := s.inbox.MarkReceived(ctx, row.ID, time.Now().UTC()); err != nil {
			return nil, nil, err
		}
	}
	return apiwire.InboxReceiveResult([]any{}, false), []any{}, nil
}

type FriendService struct {
	friends *repositories.FriendRepository
}

func NewFriendService(friends *repositories.FriendRepository) *FriendService {
	return &FriendService{friends: friends}
}
func (s *FriendService) GetReceivedRequests(ctx context.Context, userKey string) ([]any, int64, error) {
	if s == nil || s.friends == nil {
		return []any{}, 0, nil
	}
	_, err := s.friends.PendingReceived(ctx, userKey)
	if err != nil {
		return nil, 0, err
	}
	count, err := s.friends.AcceptedCount(ctx, userKey)
	if err != nil {
		return nil, 0, err
	}
	return []any{}, count, nil
}
func (s *FriendService) GetFriends(ctx context.Context, userKey string) ([]any, int64, error) {
	count, err := s.friends.AcceptedCount(ctx, userKey)
	if err != nil {
		return nil, 0, err
	}
	return []any{}, count, nil
}
func (s *FriendService) SendRequest(ctx context.Context, fromUserKey string, toUserKey string, message string) ([]any, error) {
	key, err := repositories.NewToken("freq")
	if err != nil {
		return nil, err
	}
	if err := s.friends.CreateRequest(ctx, key, fromUserKey, toUserKey, message); err != nil {
		return nil, err
	}
	return apiwire.FriendRequestResult(1), nil
}
func (s *FriendService) AcceptRequest(ctx context.Context, userKey string, requestKey string) ([]any, error) {
	req, err := s.friends.FindRequest(ctx, requestKey)
	if err != nil {
		return apiwire.FriendAcceptResult(4), nil
	}
	if req.ToUserKey != userKey || req.Status != 0 {
		return apiwire.FriendAcceptResult(4), nil
	}
	if err := s.friends.UpdateRequestStatus(ctx, requestKey, 1); err != nil {
		return nil, err
	}
	if err := s.friends.CreateRelation(ctx, req.FromUserKey, req.ToUserKey); err != nil {
		return nil, err
	}
	if err := s.friends.CreateRelation(ctx, req.ToUserKey, req.FromUserKey); err != nil {
		return nil, err
	}
	return apiwire.FriendAcceptResult(1), nil
}
func (s *FriendService) RejectRequest(ctx context.Context, userKey string, requestKey string) ([]any, error) {
	req, err := s.friends.FindRequest(ctx, requestKey)
	if err != nil || req.ToUserKey != userKey {
		return apiwire.FriendAcceptResult(4), nil
	}
	return apiwire.FriendAcceptResult(1), s.friends.UpdateRequestStatus(ctx, requestKey, 2)
}
func (s *FriendService) DeleteFriend(ctx context.Context, userKey string, friendUserKey string) ([]any, error) {
	return apiwire.BooleanResult(true), s.friends.DeleteRelation(ctx, userKey, friendUserKey)
}

type CircleService struct {
	circles *repositories.CircleRepository
	support *repositories.SupportCompanyRepository
}

func NewCircleService(circles *repositories.CircleRepository, support *repositories.SupportCompanyRepository) *CircleService {
	return &CircleService{circles: circles, support: support}
}
func (s *CircleService) GetInvitedCircles(ctx context.Context, userKey string) ([]any, error) {
	if s == nil || s.circles == nil {
		return []any{}, nil
	}
	_, err := s.circles.PendingInvites(ctx, userKey)
	if err != nil {
		return nil, err
	}
	return []any{}, nil
}
func (s *CircleService) GetSupportAndTheaterLevelInformation(ctx context.Context, userKey string) ([]any, error) {
	rows, err := s.support.EnsureForUser(ctx, userKey)
	if err != nil {
		return nil, err
	}
	statuses := make([][]any, 4)
	defaultLevelLimitStatus := apiwire.SupportCompanyLevelLimitStatus(0, []any{})
	defaultLevelUppedAt := time.Date(2026, 5, 16, 12, 0, 0, 0, time.UTC)
	for i := 0; i < 4; i++ {
		statuses[i] = apiwire.SupportCompanyLevelStatus(1, 0, defaultLevelUppedAt, 0, defaultLevelLimitStatus)
	}
	for _, row := range rows {
		if row.Company >= 0 && row.Company < 4 {
			lastLevelUppedAt := defaultLevelUppedAt
			if row.LastLevelUppedAt != nil {
				lastLevelUppedAt = *row.LastLevelUppedAt
			}
			statuses[row.Company] = apiwire.SupportCompanyLevelStatus(row.Level, row.CurrentSupportPoint, lastLevelUppedAt, row.LevelLimit, defaultLevelLimitStatus)
		}
	}
	return apiwire.SupportCompanyInformation(statuses[0], statuses[1], statuses[2], statuses[3]), nil
}
func (s *CircleService) CreateCircle(ctx context.Context, userKey string, name string, description string) ([]any, error) {
	id := time.Now().UnixNano()
	if name == "" {
		name = "Circle"
	}
	hashedID := "circle_" + strconv.FormatInt(id, 10)
	return apiwire.CreateCircleResult(14, hashedID), s.circles.CreateCircle(ctx, id, userKey, name, description)
}
func (s *CircleService) AcceptInvite(ctx context.Context, userKey string, inviteKey string) ([]any, error) {
	inv, err := s.circles.FindInvite(ctx, inviteKey)
	if err != nil || inv.UserKey != userKey {
		return apiwire.CircleResult(18), nil
	}
	if err := s.circles.UpdateInviteStatus(ctx, inviteKey, 1); err != nil {
		return nil, err
	}
	return apiwire.CircleResult(12), s.circles.AddMember(ctx, inv.CircleID, userKey, 0)
}
func (s *CircleService) RejectInvite(ctx context.Context, userKey string, inviteKey string) ([]any, error) {
	inv, err := s.circles.FindInvite(ctx, inviteKey)
	if err != nil || inv.UserKey != userKey {
		return apiwire.CircleResult(18), nil
	}
	return apiwire.CircleResult(22), s.circles.UpdateInviteStatus(ctx, inviteKey, 2)
}
func (s *CircleService) LeaveCircle(ctx context.Context, userKey string) ([]any, error) {
	return apiwire.CircleResult(16), s.circles.Leave(ctx, userKey)
}
func (s *CircleService) GetMyCircle(ctx context.Context, userKey string) ([]any, error) {
	return []any{}, nil
}
