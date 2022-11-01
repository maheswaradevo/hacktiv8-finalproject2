package models

import "time"

type SocialMedia struct {
	SocialMediaID  uint64    `db:"id"`
	Name           string    `db:"name"`
	SocialMediaURL string    `db:"social_media_url"`
	UserID         uint64    `db:"user_id"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

type SocialMedias []*SocialMedia
