package repo

import (
	"context"
	"errors"
	"log"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/model"
	"github.com/jackc/pgx/v5"
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

func (r *AuthRepo) GetUserByEmail(ctx context.Context, email string) (*model.UserFromDB, error) {
	query := `SELECT id_users, email, password_hash FROM users WHERE email = $1;`

	var user model.UserFromDB
	err := r.DB.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // User tidak ditemukan
		}
		log.Println("Error GetUserByEmail Repo:", err)
		return nil, err
	}

	return &user, nil
}
