package persistence

import (
	"context"

	"github.com/jacexh/golang-ddd-template/internal/domain/user"
	"xorm.io/xorm"
)

var (
	User user.Repository = (*userRepository)(nil)
)

type (
	userRepository struct {
		db *xorm.Engine
	}
)

func BuildUserRepository(db *xorm.Engine) user.Repository {
	User = newUserRepository(db)
	return User
}

func newUserRepository(db *xorm.Engine) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) SaveUser(ctx context.Context, u *user.UserEntity) error {
	_, err := ur.db.Context(ctx).Exec(
		"INSERT INTO user (id, name, password, email, version) VALUES (?, ?, ?, ?, 1) ON DUPLICATE KEY UPDATE name=?, password=?, version=version+1",
		u.ID, u.Name, u.Password, u.Email, u.Name, u.Password)
	u.Events.Dispatch(ctx)
	return err
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, mail string) (*user.UserEntity, error) {
	_, err := ur.db.Context(ctx).Exec("select * from user where email=? limit 1", mail)
	return nil, err
}
