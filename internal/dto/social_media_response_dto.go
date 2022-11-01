package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
)

type CreateSocialMediaResponse struct {
	SocialMediaID  uint64    `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint64    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

func NewSocialMediaCreateResponse(scmd models.SocialMedia, userID uint64, socialMediaID uint64) *CreateSocialMediaResponse {
	return &CreateSocialMediaResponse{
		SocialMediaID: 	socialMediaID,
		Name:           scmd.Name,
		SocialMediaURL: scmd.SocialMediaURL,
		UserID:         userID,
		CreatedAt:      time.Now(),
	}
}
