package dto

import "github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"

type PostPhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

func (dto *PostPhotoRequest) ToEntity() *models.Photo {
	return &models.Photo{
		Title:    dto.Title,
		Caption:  dto.Caption,
		PhotoUrl: dto.PhotoUrl,
	}
}
