package persistence

import (
	"context"

	"github.com/jacexh/golang-ddd-template/domain/event"
	"github.com/jacexh/golang-ddd-template/domain/user"
	"github.com/jacexh/golang-ddd-template/logger"
	"github.com/jacexh/golang-ddd-template/trace"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

var (
	User user.UserRepository = (*userRepository)(nil)
)

type (
	userRepository struct {
		db *xorm.Engine
	}
)

func BuildUserRepository(db *xorm.Engine) user.UserRepository {
	User = newUserRepository(db)
	return User
}

func newUserRepository(db *xorm.Engine) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) SaveUser(ctx context.Context, u *user.UserEntity) error {
	_, _ = ur.db.Context(ctx).Exec("select * from user")
	for ev, got := u.Events.Next(); got; {
		event.Publish(ctx, ev)
	}
	return nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, uid string) (*user.UserEntity, error) {
	logger.Logger.Info("get user by id", zap.String("user_id", uid), trace.MustExtractRequestIndexFromCtxAsField(ctx))
	_, err := ur.db.Context(ctx).Exec("select * from user where id=? limit 1", uid)
	return nil, err
}
