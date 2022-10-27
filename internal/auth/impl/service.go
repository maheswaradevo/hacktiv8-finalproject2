package impl

import (
	"context"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	repo AuthRepository
}

func ProvideAuthService(repo AuthRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		repo: repo,
	}
}

func (auth *AuthServiceImpl) RegisterUser(ctx context.Context, data *dto.UserRegistrationRequest) error {
	userData := data.ToEntity()

	exist, err := auth.repo.GetUserEmail(ctx, userData.Email)
	if err != nil && err != errors.ErrInvalidResources {
		log.Printf("[RegisterUser] failed to check duplicate email: %v", err)
		return err
	}

	if exist != nil {
		err = errors.ErrUserExists
		log.Printf("[RegisterUser] user with email %v already existed", data.Email)
		return err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[RegisterUser] failed to hashed the password: %v", err)
		return err
	}
	userData.Password = string(hashed)
	err = auth.repo.InsertUser(ctx, *userData)
	if err != nil {
		log.Printf("[RegisterUser] failed to store user data to database: %v", err)
		return err
	}
	return nil
}
