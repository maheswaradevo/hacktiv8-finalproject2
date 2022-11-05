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

type ViewPhotoResponse struct {
	PhotoID   uint64       `json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url"`
	UserID    uint64       `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      userResponse `json:"user"`
}

type EditPhotoResponse struct {
	PhotoID   uint64    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint64    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type userResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ViewPhotosResponse []*ViewPhotoResponse

func NewEditPhotoResponse(ph models.Photo, userID uint64) *EditPhotoResponse {
	return &EditPhotoResponse{
		PhotoID:   ph.PhotoID,
		Title:     ph.Title,
		Caption:   ph.Caption,
		PhotoUrl:  ph.PhotoUrl,
		UserID:    userID,
		UpdatedAt: time.Now(),
	}
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

func NewViewPhotoResponse(pu models.PhotoUserJoined) *ViewPhotoResponse {
	return &ViewPhotoResponse{
		PhotoID:   pu.Photo.PhotoID,
		Title:     pu.Photo.Title,
		Caption:   pu.Photo.Caption,
		PhotoUrl:  pu.Photo.PhotoUrl,
		UserID:    pu.Photo.UserID,
		CreatedAt: pu.Photo.CreatedAt,
		UpdatedAt: pu.Photo.UpdatedAt,
		User: userResponse{
			Email:    pu.User.Email,
			Username: pu.User.Username,
		},
	}
}

func NewViewPhotosResponse(pp models.PeoplePhotoJoined) ViewPhotosResponse {
	var viewPhotosResponse ViewPhotosResponse

	for idx := range pp {
		peoplePhoto := NewViewPhotoResponse(*pp[idx])
		viewPhotosResponse = append(viewPhotosResponse, peoplePhoto)
	}
	return viewPhotosResponse
}
