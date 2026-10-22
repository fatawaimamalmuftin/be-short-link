package model

import "time"

type ResLink struct {
	ID          int       `json:"id_links" db:"id_links"`
	OriginalURL string    `json:"original_url" db:"original_url"`
	Slug        string    `json:"slug" db:"slug"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UserID      int       `json:"user_id" db:"user_id"`
}