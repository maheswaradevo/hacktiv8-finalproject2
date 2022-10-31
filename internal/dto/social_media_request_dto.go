package dto

import "github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"

type CreateSocialMediaRequest struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

func (dto *CreateSocialMediaRequest) ToEntity() (scmd *models.SocialMedia) {
	scmd = &models.SocialMedia{
		Name: dto.Name,
		SocialMediaURL: dto.SocialMediaURL,
	}
	return
}