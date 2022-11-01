package photos

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/photos/impl"
)

type PhotoService interface {
	PostPhoto(ctx context.Context, data *dto.PostPhotoRequest, userID uint64) (*dto.PostPhotoResponse, error)
}

func ProvidePhotoService(db *sql.DB) PhotoService {
	repo := impl.ProvidePhotoRepository(db)
	return impl.ProvidePhotoService(repo)
}
