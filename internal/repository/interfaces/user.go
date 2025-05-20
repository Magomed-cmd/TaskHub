package interfaces

import (
	"TaskHub/internal/pkg/model"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}
