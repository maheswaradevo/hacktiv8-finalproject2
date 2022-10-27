package impl

import (
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/global/config"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
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

func (auth *AuthServiceImpl) LoginUser(ctx context.Context, data *dto.UserSignInRequest) (res *dto.UserSignInResponse, err error) {
	userInfo := data.ToEntity()

	userCred, err := auth.repo.GetUserEmail(ctx, userInfo.Email)
	if err != nil {
		log.Printf("[LoginUser] failed to fetch user with email: %v, err: %v", userInfo.Email, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(userCred.Password), []byte(userInfo.Password))
	if err != nil {
		log.Printf("[LoginUser] failed to compare the hashed password and the password, err: %v", err)
		err = errors.ErrInvalidCred
		return
	}

	token, err := auth.createAccessToken(userCred)
	if err != nil {
		log.Printf("[LoginUser] failed to create new token, err: %v", err)
		return nil, err
	}
	return dto.NewUserSignInResponse(token), nil
}

func (auth AuthServiceImpl) createAccessToken(user *models.User) (string, error) {
	cfg := config.GetConfig()

	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["exp"] = time.Now().Add(time.Minute * 1).Unix()
	claim["user_id"] = user.UserID

	token := jwt.NewWithClaims(cfg.JWT_SIGNING_METHOD, claim)
	signedToken, err := token.SignedString([]byte(cfg.API_SECRET_KEY))
	if err != nil {
		log.Printf("[createAccessToken] failed to create new token, err: %v", err)
		return "", nil
	}
	return signedToken, nil
}
