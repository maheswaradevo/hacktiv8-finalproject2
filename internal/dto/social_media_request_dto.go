package dto

import "github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"

type CreateSocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required"`
}

func (dto *CreateSocialMediaRequest) ToEntity() (scmd *models.SocialMedia) {
	scmd = &models.SocialMedia{
		Name:           dto.Name,
		SocialMediaURL: dto.SocialMediaURL,
	}
	return
}

type EditSocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

func (dto *EditSocialMediaRequest) ToEntity() *models.SocialMedia {
	return &models.SocialMedia{
		Name:           dto.Name,
		SocialMediaURL: dto.SocialMediaUrl,
	}
}
