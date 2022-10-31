package socialmedia

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/socialmedia/impl"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
)

type SocialMediaService interface {
	CreateSocialMedia(ctx context.Context, data *dto.CreateSocialMediaRequest, userID uint64) error
}

func ProvideSocialMediaService(db *sql.DB) SocialMediaService {
	repo := impl.ProvideSocialMediaRepository(db)
	return impl.ProvideSocialMediaService(repo)
}

