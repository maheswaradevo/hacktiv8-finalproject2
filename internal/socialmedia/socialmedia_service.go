package socialmedia

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/socialmedia/impl"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
)

type SocialMediaService interface {
	CreateSocialMedia(ctx context.Context, data *dto.CreateSocialMediaRequest, userID uint64) (res *dto.CreateSocialMediaResponse, err error)
	ViewSocialMedia(ctx context.Context) (dto.ViewSocialMediasResponse, error)
	UpdateSocialMedia(ctx context.Context, data *dto.EditSocialMediaRequest, socialMediaID uint64, userID uint64) (*dto.EditSocialMediaResponse, error)
	DeleteSocialMedia(ctx context.Context, socialMediaID uint64, userID uint64) (*dto.DeleteSocialMediaResponse, error)
}

func ProvideSocialMediaService(db *sql.DB) SocialMediaService {
	repo := impl.ProvideSocialMediaRepository(db)
	return impl.ProvideSocialMediaService(repo)
}

