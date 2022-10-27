package dto

import "github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"

type UserRegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
