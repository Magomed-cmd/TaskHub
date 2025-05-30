package service

import (
	"TaskHub/internal/pkg/model"
	"TaskHub/internal/repository/interfaces"
	"TaskHub/internal/utils"
	"context"
	"log"
	"net/mail"

	"github.com/google/uuid"
)

type UserService struct {
	repo interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s UserService) CreateUser(ctx context.Context, user *model.User) error {

	var err error
	_, err = mail.ParseAddress(user.Email)
	if err != nil {
		log.Println("incorrect email: ", err)
		return err
	}

	// Проверка пароля происходит в функции GetHash, поэтому здесь не нужно дополнительно обрабатывать пустой пароль

	hashedPassword, err := utils.GetHash(user.Password)
	if err != nil {
		log.Println("error to get hash: ", err)
		return err
	}

	user.Password = hashedPassword

	user.UUID = uuid.New()

	return s.repo.CreateUser(ctx, user)
}
