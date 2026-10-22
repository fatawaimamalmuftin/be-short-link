package repo

import (
	"context"
	"log"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo struct {
	DB *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{DB: db}
}

func (r *AuthRepo) CreateUser(ctx context.Context, user *model.ReqRegis) error {
	query := `INSERT INTO users (email, password_hash) VALUES ($1, $2);`

	_, err := r.DB.Exec(ctx, query, user.Email, user.Password)
	if err != nil {
		log.Println("Error CreateUser Repo:", err)
		return err
	}

	return nil
}
