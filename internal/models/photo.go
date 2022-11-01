package models

import "time"

type Photo struct {
	PhotoID   uint64    `db:"id"`
	Title     string    `db:"title"`
	Caption   string    `db:"caption"`
	PhotoUrl  string    `db:"photo_url"`
	UserID    uint64    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Photos []*Photo
