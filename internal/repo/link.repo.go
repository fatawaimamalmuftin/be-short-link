package repo

import (
	"context"
	"errors"
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

func (r *LinkRepo) CreateLink(ctx context.Context, userID int, req *model.ReqCreateLink) (*model.ResLink, error) {
	query := `
		INSERT INTO links (original_url, slug, user_id) 
		VALUES ($1, $2, $3) 
		RETURNING id_links, original_url, slug, created_at, user_id;
	`

	var link model.ResLink
	err := r.DB.QueryRow(ctx, query, req.OriginalURL, req.Slug, userID).Scan(
		&link.ID,
		&link.OriginalURL,
		&link.Slug,
		&link.CreatedAt,
		&link.UserID,
	)
	if err != nil {
		log.Println("Error CreateLink Repo:", err)
		return nil, err
	}

	return &link, nil
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

func (r *LinkRepo) DeleteLink(ctx context.Context, linkID int, userID int) error {
	query := `DELETE FROM links WHERE id_links = $1 AND user_id = $2;`

	res, err := r.DB.Exec(ctx, query, linkID, userID)
	if err != nil {
		log.Println("Error DeleteLink Repo:", err)
		return err
	}

	if res.RowsAffected() == 0 {
		return errors.New("link not found or unauthorized")
	}

	return nil
}
