package impl

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
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

	validate := validator.New()
	validateError := validate.Struct(data)
	if validateError != nil {
		validateError = errors.ErrInvalidRequestBody
		log.Printf("[PostPhoto] there's data that not through the validate process")
		return nil, validateError
	}
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

func (ph *PhotoServiceImpl) UpdatePhoto(ctx context.Context, data *dto.EditPhotoRequest, photoID uint64, userID uint64) (*dto.EditPhotoResponse, error) {
	editPhoto := data.ToEntity()

	validate := validator.New()
	validateError := validate.Struct(data)
	if validateError != nil {
		validateError = errors.ErrInvalidRequestBody
		log.Printf("[UpdatePhoto] there's data that not through the validate process")
		return nil, validateError
	}
	check, err := ph.repo.CheckPhoto(ctx, photoID, userID)
	if err != nil {
		log.Printf("[UpdatePhoto] failed to check photo with, userID: %v, err: %v", userID, err)
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		log.Printf("[UpdatePhoto] no photo in userID: %v", userID)
		return nil, err
	}
	editPhoto.PhotoID = photoID

	err = ph.repo.UpdatePhoto(ctx, *editPhoto, photoID)
	if err != nil {
		log.Printf("[UpdatePhoto] failed to update the photo, id: %v, err: %v", photoID, err)
		return nil, err
	}
	return dto.NewEditPhotoResponse(*editPhoto, userID), nil
}

func (ph *PhotoServiceImpl) DeletePhoto(ctx context.Context, photoID uint64, userID uint64) (*dto.DeletePhotoResponse, error) {
	check, err := ph.repo.CheckPhoto(ctx, photoID, userID)
	if err != nil {
		log.Printf("[DeletePhoto] failed to check photoID: %v, err: %v", photoID, err)
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		log.Printf("[DeletePhoto] data with photoID %v not found", photoID)
		return nil, err
	}
	err = ph.repo.DeletePhoto(ctx, photoID)
	if err != nil {
		log.Printf("[DeletePhoto] failed to delete photo, photoID: %v, err: %v", photoID, err)
		return nil, err
	}
	msg := "Your photo has been successfully deleted"
	return dto.NewDeletePhotoResponse(msg), nil
}
