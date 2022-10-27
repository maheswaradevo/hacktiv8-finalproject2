package auth

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/auth/impl"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
)

type AuthService interface {
	RegisterUser(ctx context.Context, data *dto.UserRegistrationRequest) error
	LoginUser(ctx context.Context, data *dto.UserSignInRequest) (res *dto.UserSignInResponse, err error)
}

func ProvideAuthService(db *sql.DB) AuthService {
	repo := impl.ProvideAuthRepository(db)
	return impl.ProvideAuthService(repo)
}
