package impl

import (
	"context"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/errors"
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

func (scmd *SocialMediaServiceImpl) ViewSocialMedia(ctx context.Context) (dto.ViewSocialMediasResponse, error) {
	count, err := scmd.repo.CountSocialMedia(ctx)
	if err != nil {
		log.Printf("[ViewSocialMedia] failed to count the social media, err: %v", err)
		return nil, err
	}
	if count == 0 {
		err = errors.ErrDataNotFound
		log.Printf("[ViewSocialMedia] no data exists in the database: %v", err)
		return nil, err
	}
	res, err := scmd.repo.ViewSocialMedia(ctx)
	if err != nil {
		log.Printf("[ViewSocialMedia] failed to view the social media, err: %v", err)
		return nil, err
	}
	return dto.NewViewSocialMediasResponse(res), nil
}

func (scmd *SocialMediaServiceImpl) UpdateSocialMedia(ctx context.Context, data *dto.EditSocialMediaRequest, socialMediaID uint64, userID uint64) (*dto.EditSocialMediaResponse, error) {
	editSocialMedia := data.ToEntity()
	check, err := scmd.repo.CheckSocialMedia(ctx, socialMediaID, userID)
	if err != nil {
		log.Printf("[UpdateSocialMedia] failed to check social media with, userID: %v, err: %v", userID, err)
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		log.Printf("[UpdateSocialMedia] no social media in userID: %v", userID)
		return nil, err
	}
	editSocialMedia.SocialMediaID = socialMediaID

	err = scmd.repo.UpdateSocialMedia(ctx, *editSocialMedia, socialMediaID)
	if err != nil {
		log.Printf("[UpdateSocialMedia] failed to update the social media, id: %v, err: %v", socialMediaID, err)
		return nil, err
	}
	return dto.NewEditSocialMediaResponse(*editSocialMedia, userID), nil
}
