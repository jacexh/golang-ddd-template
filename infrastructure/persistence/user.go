package persistence

import (
	"context"

	"{{.Module}}/domain/user"
	"{{.Module}}/logger"
	"{{.Module}}/trace"
	"go.uber.org/zap"
	"xorm.io/xorm"
)

type (
	userRepository struct {
		db *xorm.Engine
	}
)

func BuildUserRepository(db *xorm.Engine) user.UserRepository {
	return newUserRepository(db)
}

func newUserRepository(db *xorm.Engine) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) SaveUser(context.Context, *user.UserEntity) error {
	return nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, uid string) (*user.UserEntity, error) {
	logger.Logger.Info("get user by id", zap.String("user_id", uid), trace.MustExtractRequestIndexFromCtxAsField(ctx))
	_, err := ur.db.Context(ctx).Exec("select * from user where id=? limit 1", uid)
	return nil, err
}
