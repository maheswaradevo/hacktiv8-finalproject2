package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
)

type PhotoRepository interface {
	PostPhoto(ctx context.Context, data models.Photo) (photoID uint64, err error)
	ViewPhoto(ctx context.Context) (models.PeoplePhotoJoined, error)
	CountPhoto(ctx context.Context) (int, error)
}

type photoImpl struct {
	db *sql.DB
}

func ProvidePhotoRepository(db *sql.DB) *photoImpl {
	return &photoImpl{db: db}
}

var (
	INSERT_PHOTO = "INSERT INTO `photo` (title, caption, photo_url, user_id) VALUES (?, ?, ?, ?);"
	VIEW_PHOTO   = "SELECT p.id, p.title, p.caption, p.photo_url, p.user_id, p.created_at, p.updated_at, u.email, u.username FROM photo p JOIN `user` u ON u.id = p.user_id;"
	COUNT_PHOTO  = "SELECT COUNT(*) FROM photo;"
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

func (p photoImpl) ViewPhoto(ctx context.Context) (models.PeoplePhotoJoined, error) {
	query := VIEW_PHOTO
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[ViewPhoto] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[ViewPhoto] failed to query to the database, err: %v", err)
		return nil, err
	}
	var peoplePhoto models.PeoplePhotoJoined
	for rows.Next() {
		personPhoto := models.PhotoUserJoined{}
		err := rows.Scan(
			&personPhoto.Photo.PhotoID,
			&personPhoto.Photo.Title,
			&personPhoto.Photo.Caption,
			&personPhoto.Photo.PhotoUrl,
			&personPhoto.Photo.UserID,
			&personPhoto.Photo.CreatedAt,
			&personPhoto.Photo.UpdatedAt,
			&personPhoto.User.Email,
			&personPhoto.User.Username,
		)
		if err != nil {
			log.Printf("[ViewPhoto] failed to scan the data from the database, err: %v", err)
			return nil, err
		}
		peoplePhoto = append(peoplePhoto, &personPhoto)
	}
	return peoplePhoto, nil
}

func (p photoImpl) CountPhoto(ctx context.Context) (int, error) {
	query := COUNT_PHOTO
	rows := p.db.QueryRowContext(ctx, query)
	var count int
	err := rows.Scan(&count)
	if err != nil {
		log.Printf("[CountPhoto] failed to scan the data from the database, err: %v", err)
		return 0, err
	}
	return count, nil
}
