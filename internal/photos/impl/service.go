package impl

import (
	"context"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
)

type PhotoServiceImpl struct {
	repo PhotoRepository
}

func ProvidePhotoService(repo PhotoRepository) *PhotoServiceImpl {
	return &PhotoServiceImpl{repo: repo}
}

func (ph *PhotoServiceImpl) PostPhoto(ctx context.Context, data *dto.PostPhotoRequest, userID uint64) (*dto.PostPhotoResponse, error) {
	photoData := data.ToEntity()
	photoData.UserID = userID
	res, err := ph.repo.PostPhoto(ctx, *photoData)
	if err != nil {
		log.Printf("[PostPhoto] failed to post the photo, err: %v", err)
		return nil, err
	}
	photoData.PhotoID = res
	return dto.NewPostPhotoResponse(*photoData, userID), nil
}
