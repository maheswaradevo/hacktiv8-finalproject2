package photos

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/photos/impl"
)

type PhotoService interface {
	PostPhoto(ctx context.Context, data *dto.PostPhotoRequest, userID uint64) (*dto.PostPhotoResponse, error)
	ViewPhoto(ctx context.Context) (dto.ViewPhotosResponse, error)
	UpdatePhoto(ctx context.Context, data *dto.EditPhotoRequest, photoID uint64, userID uint64) (*dto.EditPhotoResponse, error)
}

func ProvidePhotoService(db *sql.DB) PhotoService {
	repo := impl.ProvidePhotoRepository(db)
	return impl.ProvidePhotoService(repo)
}
