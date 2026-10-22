package model

import "time"

type ReqCreateLink struct {
	OriginalURL string `json:"original_url" binding:"required,url"`
	Slug        string `json:"slug" binding:"required,alphanum,min=3,max=50"`
}

type ResLink struct {
	ID          int       `json:"id_links" db:"id_links"`
	OriginalURL string    `json:"original_url" db:"original_url"`
	Slug        string    `json:"slug" db:"slug"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UserID      int       `json:"user_id" db:"user_id"`
}
