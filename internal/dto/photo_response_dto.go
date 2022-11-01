package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
)

type PostPhotoResponse struct {
	PhotoID   uint64    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPostPhotoResponse(ph models.Photo, userID uint64) *PostPhotoResponse {
	return &PostPhotoResponse{
		PhotoID:   ph.PhotoID,
		Title:     ph.Title,
		Caption:   ph.Caption,
		PhotoUrl:  ph.PhotoUrl,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
}
