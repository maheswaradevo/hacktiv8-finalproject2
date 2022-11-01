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

func (scmd *SocialMediaServiceImpl) CreateSocialMedia(ctx context.Context, data *dto.CreateSocialMediaRequest, userID uint64) (res *dto.CreateSocialMediaResponse, err error) {
	socialMediaData := data.ToEntity()

	socialMediaID, err := scmd.repo.CreateSocialMedia(ctx, *socialMediaData, userID) 
	if err != nil {
		log.Printf("[CreateSocialMedia] failed to store user data to database: %v", err)
		return 
	}

	return dto.NewSocialMediaCreateResponse(*socialMediaData, userID, socialMediaID), nil
}
