package service

import (
	"context"
	"log"

	customerror "github.com/fatawaimamalmuftin/be-shortlink/internal/customError"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/model"
	"github.com/fatawaimamalmuftin/be-shortlink/internal/repo"
	jwtpkg "github.com/fatawaimamalmuftin/be-shortlink/pkg"
	"github.com/fatawaimamalmuftin/be-shortlink/pkg/blacklist"
	"github.com/fatawaimamalmuftin/be-shortlink/pkg/hasing"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	Repo *repo.AuthRepo
	RDB  *redis.Client
}

func NewAuthService(r *repo.AuthRepo, rdb *redis.Client) *AuthService {
	return &AuthService{
		Repo: r,
		RDB:  rdb,
	}
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

func (s *AuthService) LoginUser(ctx context.Context, req *model.ReqLogin) (*model.ResLogin, error) {
	user, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, customerror.PasMissMach
	}

	if err := hasing.ComparePassHash(req.Password, user.PasswordHash); err != nil {
		return nil, customerror.PasMissMach
	}

	claims := jwtpkg.NewJWTclem(user.ID, user.Email)
	token, err := claims.GenToken()
	if err != nil {
		log.Println("Error GenerateToken Service:", err)
		return nil, err
	}

	return &model.ResLogin{
		Token: token,
	}, nil
}

func (s *AuthService) LogoutUser(ctx context.Context, token string) error {
	if err := blacklist.AddToken(token, s.RDB, ctx); err != nil {
		log.Println("Error AddToken Blacklist Service:", err)
		return err
	}

	return nil
}
