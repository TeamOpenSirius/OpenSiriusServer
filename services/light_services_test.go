package services

import (
	"context"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"ymst_server/core/database"
	"ymst_server/repositories"
)

func newLightServiceDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file:light_services_"+time.Now().Format("150405.000000000")+"?mode=memory&cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := database.AutoMigrate(db); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestAccountExtraConfirmationCodeReusesActiveCode(t *testing.T) {
	db := newLightServiceDB(t)
	svc := NewAccountExtraService(repositories.NewAccountRepository(db), repositories.NewSessionRepository(db), repositories.NewConfirmationCodeRepository(db))
	code1, remaining1, err := svc.GetConfirmationCode(context.Background(), "usr_a")
	if err != nil {
		t.Fatal(err)
	}
	code2, remaining2, err := svc.GetConfirmationCode(context.Background(), "usr_a")
	if err != nil {
		t.Fatal(err)
	}
	if code1 == "" || code1 != code2 || remaining1 <= 0 || remaining2 <= 0 {
		t.Fatalf("unexpected codes %q %q", code1, code2)
	}
}

func TestLoginAndHomePayloadShapes(t *testing.T) {
	login := NewLoginService(nil)
	payload, updates, err := login.Login(context.Background(), "usr_a", "push")
	if err != nil {
		t.Fatal(err)
	}
	if len(payload) != 7 || len(updates) != 0 {
		t.Fatalf("unexpected login payload")
	}
	home := NewHomeService()
	bonus, bonusUpdates, err := home.CheckReceiveLoginBonus(context.Background(), "usr_a")
	if err != nil {
		t.Fatal(err)
	}
	if len(bonus) != 0 || len(bonusUpdates) != 0 {
		t.Fatalf("unexpected bonus payload")
	}
	payment, _, err := home.CheckExternalPayment(context.Background(), "usr_a")
	if err != nil {
		t.Fatal(err)
	}
	if len(payment) != 1 {
		t.Fatalf("unexpected payment payload")
	}
}

func TestInboxFriendCircleDefaults(t *testing.T) {
	db := newLightServiceDB(t)
	inbox := NewInboxService(repositories.NewInboxRepository(db))
	has, err := inbox.CheckPackages(context.Background(), "usr_a")
	if err != nil {
		t.Fatal(err)
	}
	if has {
		t.Fatalf("empty inbox returned true")
	}
	friend := NewFriendService(repositories.NewFriendRepository(db))
	results, count, err := friend.GetReceivedRequests(context.Background(), "usr_a")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 0 || count != 0 {
		t.Fatalf("unexpected friend result")
	}
	circle := NewCircleService(repositories.NewCircleRepository(db), repositories.NewSupportCompanyRepository(db))
	invited, err := circle.GetInvitedCircles(context.Background(), "usr_a")
	if err != nil {
		t.Fatal(err)
	}
	if len(invited) != 0 {
		t.Fatalf("unexpected invited circles")
	}
	support, err := circle.GetSupportAndTheaterLevelInformation(context.Background(), "usr_a")
	if err != nil {
		t.Fatal(err)
	}
	if len(support) != 4 {
		t.Fatalf("support len = %d", len(support))
	}
}

func TestInboxReceivePackageIsUserScoped(t *testing.T) {
	db := newLightServiceDB(t)
	now := time.Now().UTC()
	if err := db.Create(&database.GameInboxPackage{UserKey: "usr_a", PackageKey: "pkg_same", Title: "A"}).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&database.GameInboxPackage{UserKey: "usr_b", PackageKey: "pkg_same", Title: "B"}).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&database.GameInboxPackage{UserKey: "usr_a", PackageKey: "pkg_old", ExpiresAt: &now}).Error; err != nil {
		t.Fatal(err)
	}
	inbox := NewInboxService(repositories.NewInboxRepository(db))
	has, err := inbox.CheckPackages(context.Background(), "usr_a")
	if err != nil || !has {
		t.Fatalf("usr_a package check = %v, %v", has, err)
	}
	payload, _, err := inbox.ReceivePackage(context.Background(), "usr_a", "pkg_same")
	if err != nil || len(payload) != 2 {
		t.Fatalf("receive package failed: %#v %v", payload, err)
	}
	hasA, err := inbox.CheckPackages(context.Background(), "usr_a")
	if err != nil {
		t.Fatal(err)
	}
	hasB, err := inbox.CheckPackages(context.Background(), "usr_b")
	if err != nil {
		t.Fatal(err)
	}
	if hasA || !hasB {
		t.Fatalf("package isolation failed: usr_a=%v usr_b=%v", hasA, hasB)
	}
}

func TestFriendRequestLifecycleIsUserScoped(t *testing.T) {
	db := newLightServiceDB(t)
	friend := NewFriendService(repositories.NewFriendRepository(db))
	if _, err := friend.SendRequest(context.Background(), "usr_a", "usr_b", "hello"); err != nil {
		t.Fatal(err)
	}
	var req database.GameFriendRequest
	if err := db.Where("from_user_key = ? AND to_user_key = ?", "usr_a", "usr_b").First(&req).Error; err != nil {
		t.Fatal(err)
	}
	if payload, err := friend.AcceptRequest(context.Background(), "usr_c", req.RequestKey); err != nil || payload[0] != int64(4) {
		t.Fatalf("wrong user accepted request: %#v %v", payload, err)
	}
	if payload, err := friend.AcceptRequest(context.Background(), "usr_b", req.RequestKey); err != nil || payload[0] != int64(1) {
		t.Fatalf("accept failed: %#v %v", payload, err)
	}
	_, countA, err := friend.GetFriends(context.Background(), "usr_a")
	if err != nil {
		t.Fatal(err)
	}
	_, countB, err := friend.GetFriends(context.Background(), "usr_b")
	if err != nil {
		t.Fatal(err)
	}
	if countA != 1 || countB != 1 {
		t.Fatalf("friend relation count = %d/%d", countA, countB)
	}
	if _, err := friend.DeleteFriend(context.Background(), "usr_a", "usr_b"); err != nil {
		t.Fatal(err)
	}
	_, countA, _ = friend.GetFriends(context.Background(), "usr_a")
	_, countB, _ = friend.GetFriends(context.Background(), "usr_b")
	if countA != 0 || countB != 0 {
		t.Fatalf("delete friend failed: %d/%d", countA, countB)
	}
}

func TestCircleInviteAndMembershipAreUserScoped(t *testing.T) {
	db := newLightServiceDB(t)
	circle := NewCircleService(repositories.NewCircleRepository(db), repositories.NewSupportCompanyRepository(db))
	if _, err := circle.CreateCircle(context.Background(), "usr_owner", "Test Circle", "desc"); err != nil {
		t.Fatal(err)
	}
	var created database.GameCircle
	if err := db.Where("owner_user_key = ?", "usr_owner").First(&created).Error; err != nil {
		t.Fatal(err)
	}
	invite := database.GameCircleInvite{InviteKey: "invite_a", UserKey: "usr_a", CircleID: created.CircleID, CircleName: created.CircleName}
	if err := db.Create(&invite).Error; err != nil {
		t.Fatal(err)
	}
	if payload, err := circle.AcceptInvite(context.Background(), "usr_b", "invite_a"); err != nil || payload[0] != int64(18) {
		t.Fatalf("wrong user accepted invite: %#v %v", payload, err)
	}
	if payload, err := circle.AcceptInvite(context.Background(), "usr_a", "invite_a"); err != nil || payload[0] != int64(12) {
		t.Fatalf("accept invite failed: %#v %v", payload, err)
	}
	var count int64
	if err := db.Model(&database.GameCircleMember{}).Where("circle_id = ? AND user_key = ?", created.CircleID, "usr_a").Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("circle member count = %d", count)
	}
	if _, err := circle.LeaveCircle(context.Background(), "usr_a"); err != nil {
		t.Fatal(err)
	}
	if err := db.Model(&database.GameCircleMember{}).Where("circle_id = ? AND user_key = ?", created.CircleID, "usr_a").Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Fatalf("circle leave failed, count = %d", count)
	}
}
