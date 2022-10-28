package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
)

type UserSignInResponse struct {
	AccessToken string `json:"access_token"`
}

type UserEditProfileResponse struct {
	UserID    uint64    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserSignInResponse(ac string) *UserSignInResponse {
	return &UserSignInResponse{
		AccessToken: ac,
	}
}

func NewUserEditProfileResponse(usr models.User, userID uint64) *UserEditProfileResponse {
	return &UserEditProfileResponse{
		UserID:    userID,
		Email:     usr.Email,
		Username:  usr.Username,
		Age:       usr.Age,
		UpdatedAt: time.Now(),
	}
}
