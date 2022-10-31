package dto

import "github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"

type SocialMediaResponse struct {
	ID             uint64 `json:"id"`
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

type SocialMediaResponses []SocialMediaResponse

func CreateSocialMediaResponse(scmd models.SocialMedia) SocialMediaResponse {
	return SocialMediaResponse{
		ID:             scmd.SocialMediaID,
		Name:           scmd.Name,
		SocialMediaURL: scmd.SocialMediaURL,
	}
}

func CreateSocialMediaResponses(t models.SocialMedias) *SocialMediaResponses {
	var socialMediaResponses SocialMediaResponses

	for _, idx := range t {
		todo := CreateSocialMediaResponse(*idx)
		socialMediaResponses = append(socialMediaResponses, todo)
	}
	return &socialMediaResponses
}
