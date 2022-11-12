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
	UpdatePhoto(ctx context.Context, reqData models.Photo, photoID uint64) error
	CheckPhoto(ctx context.Context, photoID uint64, userID uint64) (bool, error)
	DeletePhoto(ctx context.Context, photoID uint64) error
}

type photoImpl struct {
	db *sql.DB
}

func ProvidePhotoRepository(db *sql.DB) *photoImpl {
	return &photoImpl{db: db}
}

var (
	INSERT_PHOTO = "INSERT INTO `photo` (title, caption, photo_url, user_id) VALUES (?, ?, ?, ?);"
	VIEW_PHOTO   = "SELECT p.id, p.title, p.caption, p.photo_url, p.user_id, p.created_at, p.updated_at, u.email, u.username FROM photo p JOIN `user` u ON u.id = p.user_id ORDER BY p.created_at DESC;"
	COUNT_PHOTO  = "SELECT COUNT(*) FROM photo;"
	UPDATE_PHOTO = "UPDATE photo SET title = ?, caption = ?, photo_url = ? WHERE id = ?;"
	CHECK_PHOTO  = "SELECT id FROM photo WHERE id = ? AND user_id = ?;"
	DELETE_PHOTO = "DELETE FROM photo WHERE id=?;"
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

func (p photoImpl) UpdatePhoto(ctx context.Context, reqData models.Photo, photoID uint64) error {
	query := UPDATE_PHOTO

	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[UpdatePhoto] failed to prepare the statement, err: %v", err)
		return err
	}
	_, err = stmt.ExecContext(ctx, reqData.Title, reqData.Caption, reqData.PhotoUrl, photoID)
	if err != nil {
		log.Printf("[UpdatePhoto] failed to store data to the database, err: %v", err)
		return err
	}
	return nil
}

func (p photoImpl) CheckPhoto(ctx context.Context, photoID uint64, userID uint64) (bool, error) {
	query := CHECK_PHOTO
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CheckPhoto] failed to prepare the statement, err: %v", err)
		return false, err
	}
	rows, err := stmt.QueryContext(ctx, photoID, userID)
	if err != nil {
		log.Printf("[CheckPhoto] failed to query to the database, err: %v", err)
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func (p photoImpl) DeletePhoto(ctx context.Context, photoID uint64) error {
	query := DELETE_PHOTO
	stmt, err := p.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[DeletePhoto] failed to prepare the statement, err: %v", err)
		return err
	}
	_, err = stmt.QueryContext(ctx, photoID)
	if err != nil {
		log.Printf("[DeletePhoto] failed to delete the data, photoID: %v, err: %v", photoID, err)
		return err
	}
	return nil
}
