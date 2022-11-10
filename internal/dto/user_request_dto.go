package dto

import "github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"

type UserRegistrationRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Username string `json:"username" validate:"required"`
	Age      int    `json:"age" validate:"required,numeric,min=9"`
}

type UserSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserEditProfileRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Age      int    `json:"age"`
}

func (dto *UserRegistrationRequest) ToEntity() (usr *models.User) {
	usr = &models.User{
		Email:    dto.Email,
		Password: dto.Password,
		Username: dto.Username,
		Age:      dto.Age,
	}
	return
}

func (dto *UserSignInRequest) ToEntity() (usr *models.User) {
	usr = &models.User{
		Email:    dto.Email,
		Password: dto.Password,
	}
	return
}

func (dto *UserEditProfileRequest) ToEntity() (usr *models.User) {
	usr = &models.User{
		Email:    dto.Email,
		Username: dto.Username,
		Age:      dto.Age,
	}
	return
}
