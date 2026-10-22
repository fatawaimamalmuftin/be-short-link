package service

import (
	"context"
	"log"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/model"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/repo"
	"github.com/fatawaimamalmuftin/be-shortlink/pkg/hasing"
)

type AuthService struct {
	Repo *repo.AuthRepo
}

func NewAuthService(r *repo.AuthRepo) *AuthService {
	return &AuthService{Repo: r}
}

func (s *AuthService) RegisterUser(ctx context.Context, newUser *model.ReqRegis) error {
	hashConfig := hasing.GenRecomConfHash()

	hashedPassword, err := hashConfig.GenPasHash(newUser.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return err
	}

	newUser.Password = hashedPassword

	if err := s.Repo.CreateUser(ctx, newUser); err != nil {
		log.Println("Error RegisterUser Service:", err)
		return err
	}

	return nil
}
