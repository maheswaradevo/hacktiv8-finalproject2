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

type ViewSocialMediaResponse struct {
	SocialMediaID  uint64       `json:"id"`
	Name           string       `json:"name"`
	SocialMediaURL string       `json:"social_media_url"`
	UserID         uint64       `json:"user_id"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	User           userResponse `json:"user"`
}

type ViewSocialMediasResponse []*ViewSocialMediaResponse

type EditSocialMediaResponse struct {
	SocialMediaID  uint64    `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         uint64    `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewSocialMediaCreateResponse(scmd models.SocialMedia, userID uint64, socialMediaID uint64) *CreateSocialMediaResponse {
	return &CreateSocialMediaResponse{
		SocialMediaID:  socialMediaID,
		Name:           scmd.Name,
		SocialMediaURL: scmd.SocialMediaURL,
		UserID:         userID,
		CreatedAt:      time.Now(),
	}
}

func NewViewSocialMediaResponse(su models.SocialMediaUserJoined) *ViewSocialMediaResponse {
	return &ViewSocialMediaResponse{
		SocialMediaID:  su.SocialMedia.SocialMediaID,
		Name:           su.SocialMedia.Name,
		SocialMediaURL: su.SocialMedia.SocialMediaURL,
		UserID:         su.SocialMedia.UserID,
		CreatedAt:      su.SocialMedia.CreatedAt,
		UpdatedAt:      su.SocialMedia.UpdatedAt,
		User: userResponse{
			Email:    su.User.Email,
			Username: su.User.Username,
		},
	}
}

func NewViewSocialMediasResponse(pp models.PeopleSocialMediaJoined) ViewSocialMediasResponse {
	var viewSocialMediasResponse ViewSocialMediasResponse

	for idx := range pp {
		peopleSocialMedia := NewViewSocialMediaResponse(*pp[idx])
		viewSocialMediasResponse = append(viewSocialMediasResponse, peopleSocialMedia)
	}
	return viewSocialMediasResponse
}

func NewEditSocialMediaResponse(scmd models.SocialMedia, userID uint64) *EditSocialMediaResponse {
	return &EditSocialMediaResponse{
		SocialMediaID:  scmd.SocialMediaID,
		Name:           scmd.Name,
		SocialMediaURL: scmd.SocialMediaURL,
		UserID:         userID,
		UpdatedAt:      time.Now(),
	}
}
