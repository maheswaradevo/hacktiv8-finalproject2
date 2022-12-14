package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
)

type SocialMediaRepository interface {
	CreateSocialMedia(ctx context.Context, data models.SocialMedia, userID uint64) (uint64, error)
	ViewSocialMedia(ctx context.Context) (models.PeopleSocialMediaJoined, error)
	CountSocialMedia(ctx context.Context) (int, error)
	UpdateSocialMedia(ctx context.Context, reqData models.SocialMedia, socialMediaID uint64) error
	CheckSocialMedia(ctx context.Context, socialMediaID uint64, userID uint64) (bool, error)
	DeleteSocialMedia(ctx context.Context, socialMediaID uint64) error
}
type SocialMediaImpl struct {
	db *sql.DB
}

func ProvideSocialMediaRepository(db *sql.DB) *SocialMediaImpl {
	return &SocialMediaImpl{
		db: db,
	}
}

var (
	CREATE_SOCIAL_MEDIA = "INSERT INTO `social_media`(name, social_media_url, user_id) VALUES (?, ?, ?);"
	VIEW_SOCIAL_MEDIA   = "SELECT s.id, s.name, s.social_media_url, s.user_id, s.created_at, s.updated_at, u.email, u.username FROM social_media s JOIN `user` u ON u.id = s.user_id ORDER BY s.created_at DESC;"
	COUNT_SOCIAL_MEDIA  = "SELECT COUNT(*) FROM social_media;"
	UPDATE_SOCIAL_MEDIA = "UPDATE social_media SET name = ?, social_media_url = ? WHERE id = ?;"
	CHECK_SOCIAL_MEDIA  = "SELECT id FROM social_media WHERE id = ? AND user_id = ?;"
	DELETE_SOCIAL_MEDIA = "DELETE FROM social_media WHERE id=?;"
)

func (scmd *SocialMediaImpl) CreateSocialMedia(ctx context.Context, data models.SocialMedia, userID uint64) (uint64, error) {
	query := CREATE_SOCIAL_MEDIA
	stmt, err := scmd.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CreateSocialMedia] failed to prepare the statement: %v", err)
		return uint64(0), err
	}

	res, err := stmt.ExecContext(ctx, data.Name, data.SocialMediaURL, userID)
	if err != nil {
		log.Printf("[CreateSocialMedia] failed to insert user to the database: %v", err)
		return uint64(0), err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("[CreateSocialMedia] failed to insert user to the database: %v", err)
		return uint64(id), err
	}

	return uint64(id), nil

}

func (scmd *SocialMediaImpl) ViewSocialMedia(ctx context.Context) (models.PeopleSocialMediaJoined, error) {
	query := VIEW_SOCIAL_MEDIA
	stmt, err := scmd.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[ViewSocialMedia] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[ViewSocialMedia] failed to query to the database, err: %v", err)
		return nil, err
	}
	var peopleSocialMedia models.PeopleSocialMediaJoined
	for rows.Next() {
		personSocialMedia := models.SocialMediaUserJoined{}
		err := rows.Scan(
			&personSocialMedia.SocialMedia.SocialMediaID,
			&personSocialMedia.SocialMedia.Name,
			&personSocialMedia.SocialMedia.SocialMediaURL,
			&personSocialMedia.SocialMedia.UserID,
			&personSocialMedia.SocialMedia.CreatedAt,
			&personSocialMedia.SocialMedia.UpdatedAt,
			&personSocialMedia.User.Email,
			&personSocialMedia.User.Username,
		)
		if err != nil {
			log.Printf("[ViewPhoto] failed to scan the data from the database, err: %v", err)
			return nil, err
		}
		peopleSocialMedia = append(peopleSocialMedia, &personSocialMedia)
	}
	return peopleSocialMedia, nil
}

func (scmd *SocialMediaImpl) CountSocialMedia(ctx context.Context) (int, error) {
	query := COUNT_SOCIAL_MEDIA
	rows := scmd.db.QueryRowContext(ctx, query)
	var count int
	err := rows.Scan(&count)
	if err != nil {
		log.Printf("[CountSocialMedia] failed to scan the data from the database, err: %v", err)
		return 0, err
	}
	return count, nil
}


func (scmd *SocialMediaImpl) UpdateSocialMedia(ctx context.Context, reqData models.SocialMedia, socialMediaID uint64) error {
	query := UPDATE_SOCIAL_MEDIA

	stmt, err := scmd.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[UpdateSocialMedia] failed to prepare the statement, err: %v", err)
		return err
	}
	_, err = stmt.ExecContext(ctx, reqData.Name, reqData.SocialMediaURL, socialMediaID)
	if err != nil {
		log.Printf("[UpdateSocialMedia] failed to store data to the database, err: %v", err)
		return err
	}
	return nil
}

func (scmd *SocialMediaImpl) CheckSocialMedia(ctx context.Context, socialMediaID uint64, userID uint64) (bool, error) {
	query := CHECK_SOCIAL_MEDIA
	stmt, err := scmd.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CheckSocialMedia] failed to prepare the statement, err: %v", err)
		return false, err
	}
	rows, err := stmt.QueryContext(ctx, socialMediaID, userID)
	if err != nil {
		log.Printf("[CheckSocialMedia] failed to query to the database, err: %v", err)
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func (scmd *SocialMediaImpl) DeleteSocialMedia(ctx context.Context, socialMediaID uint64) error {
	query := DELETE_SOCIAL_MEDIA
	stmt, err := scmd.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[DeleteSocialMedia] failed to prepare the statement, err: %v", err)
		return err
	}
	_, err = stmt.QueryContext(ctx, socialMediaID)
	if err != nil {
		log.Printf("[DeleteSocialMedia] failed to delete the data, socialMediaID: %v, err: %v", socialMediaID, err)
		return err
	}
	return nil
}