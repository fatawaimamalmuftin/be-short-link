package repo

import (
	"context"
	"log"

	"github.com/fatawaimamalmuftin/be-shortlink/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkRepo struct {
	DB *pgxpool.Pool
}

func NewLinkRepo(db *pgxpool.Pool) *LinkRepo {
	return &LinkRepo{DB: db}
}

func (r *LinkRepo) GetLinksByUserID(ctx context.Context, userID int) ([]model.ResLink, error) {
	query := `SELECT id_links, original_url, slug, created_at, user_id FROM links WHERE user_id = $1 ORDER BY created_at DESC;`

	rows, err := r.DB.Query(ctx, query, userID)
	if err != nil {
		log.Println("Error GetLinksByUserID Query Repo:", err)
		return nil, err
	}
	defer rows.Close()

	links := []model.ResLink{}
	for rows.Next() {
		var link model.ResLink
		if err := rows.Scan(&link.ID, &link.OriginalURL, &link.Slug, &link.CreatedAt, &link.UserID); err != nil {
			log.Println("Error Scan GetLinksByUserID Repo:", err)
			return nil, err
		}
		links = append(links, link)
	}

	return links, nil
}
