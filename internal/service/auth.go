package service

import (
	"TaskHub/internal/pkg/model"
	"TaskHub/internal/repository/interfaces"
	"TaskHub/internal/utils"
	"context"
	"log"
	"net/mail"
)

type AuthService struct {
	repo      interfaces.UserRepository
	jwtSecret string
}

func NewAuthService(repo interfaces.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{repo: repo, jwtSecret: jwtSecret}
}

func (s AuthService) Login(ctx context.Context, InputData model.LoginInput) (*string, error) {

	if _, err := mail.ParseAddress(InputData.Email); err != nil {
		log.Println("Invalid email address", err)
		return nil, err
	}

	user, err := s.repo.GetUserByEmail(ctx, InputData.Email)
	if err != nil {
		log.Println("error to get user by email: ", err)
		return nil, err
	}
	log.Printf("Loaded user: %+v", user)
	log.Println("user.Password =", user.Password)

	err = utils.CheckPassword(InputData.Password, user.Password)
	if err != nil {
		log.Println("Incorrect password or email", err)
		return nil, err
	}

	token, err := utils.GenerateJWT(*user, s.jwtSecret)
	if err != nil {
		log.Println("error to generate jwt token: ", err)
		return nil, err
	}

	return token, nil
}
