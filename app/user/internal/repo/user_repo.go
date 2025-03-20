package repo

import (
	"context"
	"gorm.io/gorm"
)

var _ domain.UserRepo = &MysqlUserRepo{}

type MysqlUserRepo struct {
	db *gorm.DB
}

func NewMysqlUserRepo(db *gorm.DB) *MysqlUserRepo {
	return &MysqlUserRepo{
		db: db,
	}
}
func (r *MysqlUserRepo) CreateUser(ctx context.Context, user *domain.User) (err error) {
	return r.db.WithContext(ctx).Model(&domain.User{}).Create(user).Error
}
func (r *MysqlUserRepo) FindUserByEmail(ctx context.Context, email string) (user *domain.User, err error) {
	user = &domain.User{}
	if err = r.db.
		WithContext(ctx).
		Model(&domain.User{}).
		Find(user, "email = ?", email).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}
