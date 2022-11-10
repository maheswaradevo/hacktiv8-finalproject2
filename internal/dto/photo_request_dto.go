package dto

import "github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"

type PostPhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url" validate:"required"`
}

type EditPhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

func (dto *EditPhotoRequest) ToEntity() *models.Photo {
	return &models.Photo{
		Title:    dto.Title,
		Caption:  dto.Caption,
		PhotoUrl: dto.PhotoUrl,
	}
}

func (dto *PostPhotoRequest) ToEntity() *models.Photo {
	return &models.Photo{
		Title:    dto.Title,
		Caption:  dto.Caption,
		PhotoUrl: dto.PhotoUrl,
	}
}
