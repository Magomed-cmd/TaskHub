package postgres

import (
	"TaskHub/pkg/model"
	"context"
	"gorm.io/gorm"
	"log"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) CreateUser(ctx context.Context, user *model.User) error {

	err := u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		log.Printf("failed to create user: %+v, error: %v", user, err)
		return err
	}
	return nil
}

func (u *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		log.Printf("failed to get user by email %s: %v", email, err)
		return nil, err
	}
	return &user, nil
}
