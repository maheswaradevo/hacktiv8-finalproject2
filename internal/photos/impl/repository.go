package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
)

type PhotoRepository interface {
	PostPhoto(ctx context.Context, data models.Photo) (photoID uint64, err error)
}

type photoImpl struct {
	db *sql.DB
}

func ProvidePhotoRepository(db *sql.DB) *photoImpl {
	return &photoImpl{db: db}
}

var (
	INSERT_PHOTO = "INSERT INTO `photo` (title, caption, photo_url, user_id) VALUES (?, ?, ?, ?);"
)

func (p photoImpl) PostPhoto(ctx context.Context, data models.Photo) (photoID uint64, err error) {
	query := INSERT_PHOTO
	rows, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[PostPhoto] failed to prepare the statement, err: %v", err)
		return
	}
	res, err := rows.ExecContext(ctx, data.Title, data.Caption, data.PhotoUrl, data.UserID)
	if err != nil {
		log.Printf("[PostPhoto] failed to insert the data to database, err: %v", err)
		return
	}
	id, _ := res.LastInsertId()
	photoID = uint64(id)
	return photoID, nil
}
