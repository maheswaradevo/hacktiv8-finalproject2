package impl

import (
	"context"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/errors"
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

func (ph *PhotoServiceImpl) ViewPhoto(ctx context.Context) (dto.ViewPhotosResponse, error) {
	count, err := ph.repo.CountPhoto(ctx)
	if err != nil {
		log.Printf("[ViewPhoto] failed to count the photo, err: %v", err)
		return nil, err
	}
	if count == 0 {
		err = errors.ErrDataNotFound
		log.Printf("[ViewPhoto] no data exists in the database: %v", err)
		return nil, err
	}
	res, err := ph.repo.ViewPhoto(ctx)
	if err != nil {
		log.Printf("[ViewPhoto] failed to view the photo, err: %v", err)
		return nil, err
	}
	return dto.NewViewPhotosResponse(res), nil
}
