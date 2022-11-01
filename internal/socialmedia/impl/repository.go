package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
)

type SocialMediaRepository interface {
	CreateSocialMedia(ctx context.Context, data models.SocialMedia, userID uint64) (uint64, error)
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
