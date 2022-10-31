package impl

import (
	"context"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
)

type SocialMediaServiceImpl struct {
	repo SocialMediaRepository
}

func ProvideSocialMediaService(repo SocialMediaRepository) *SocialMediaServiceImpl {
	return &SocialMediaServiceImpl{
		repo: repo,
	}
}

func (scmd *SocialMediaServiceImpl) CreateSocialMedia(ctx context.Context, data *dto.CreateSocialMediaRequest, userID uint64) error {
	socialMediaData := data.ToEntity()

	err := scmd.repo.CreateSocialMedia(ctx, *socialMediaData, userID)
	if err != nil {
		log.Printf("[RegisterUser] failed to store user data to database: %v", err)
		return err
	}
	return nil
}
