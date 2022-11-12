package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, data models.Comment) (commentID uint64, err error)
	ViewComment(ctx context.Context) (models.PeopleCommentJoined, error)
	CountComment(ctx context.Context) (int, error)
	UpdateComment(ctx context.Context, reqData models.CommentUserJoined, commentID uint64, userID uint64) error
	DeleteComment(ctx context.Context, commentID uint64, userID uint64) error
	CheckComment(ctx context.Context, commentID uint64, userID uint64) (bool, error)
	GetPhotoByID(ctx context.Context, commentID uint64, userID uint64) (*dto.EditPhotoResponse, error)
}

type CommentImplRepo struct {
	db *sql.DB
}

func ProvideCommentRepository(db *sql.DB) *CommentImplRepo {
	return &CommentImplRepo{
		db: db,
	}
}

var (
	CREATE_COMMENT  = "INSERT INTO `comment`(user_id, photo_id, message) VALUES (?, ?, ?);"
	VIEW_COMMENT    = "SELECT c.id, c.message, c.photo_id, c.user_id, c.updated_at, c.created_at, u.id, u.email, u.username, p.id, p.title, p.caption, p.photo_url, p.user_id FROM comment c INNER JOIN `user` u ON u.id = c.user_id INNER JOIN `photo` p ON p.id = c.photo_id ORDER BY c.created_at DESC;"
	COUNT_COMMENT   = "SELECT COUNT(*) FROM comment;"
	CHECK_COMMENT   = "SELECT id FROM comment WHERE id = ? AND user_id = ?;"
	UPDATE_COMMENT  = "UPDATE comment SET message=? WHERE id=? AND user_id=?;"
	DELETE_COMMENT  = "DELETE FROM comment c WHERE c.id = ? AND c.user_id = ?"
	GET_PHOTO_BY_ID = "SELECT p.id, p.title, p.caption, p.photo_url, p.user_id, p.updated_at FROM `comment` c INNER join `photo` p ON p.id = c.photo_id WHERE c.id = ? AND c.user_id = ?"
)

func (cmt CommentImplRepo) CreateComment(ctx context.Context, data models.Comment) (commentID uint64, err error) {
	query := CREATE_COMMENT
	stmt, err := cmt.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CreateComment] failed to prepare statement: %v", err)
		return uint64(0), err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, data.UserID, data.PhotoID, data.Message)
	if err != nil {
		log.Printf("[CreateComment] failed to insert user to the database: %v", err)
		return uint64(0), err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("[CreateComment] failed to insert user to the database: %v", err)
		return uint64(id), err
	}
	commentID = uint64(id)

	return commentID, nil
}

func (cmt CommentImplRepo) ViewComment(ctx context.Context) (models.PeopleCommentJoined, error) {
	query := VIEW_COMMENT
	stmt, err := cmt.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[ViewPhoto] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("[ViewPhoto] failed to query to the database, err: %v", err)
		return nil, err
	}
	var peopleComment models.PeopleCommentJoined
	for rows.Next() {
		personComment := models.CommentUserJoined{}
		err := rows.Scan(
			&personComment.Comment.CommentID,
			&personComment.Comment.Message,
			&personComment.Comment.PhotoID,
			&personComment.Comment.UserID,
			&personComment.Comment.UpdatedAt,
			&personComment.Comment.CreatedAt,
			&personComment.User.UserID,
			&personComment.User.Email,
			&personComment.User.Username,
			&personComment.Photo.PhotoID,
			&personComment.Photo.Title,
			&personComment.Photo.Caption,
			&personComment.Photo.PhotoUrl,
			&personComment.Photo.UserID,
		)
		if err != nil {
			log.Printf("[ViewComment] failed to scan the data from the database, err: %v", err)
			return nil, err
		}
		peopleComment = append(peopleComment, &personComment)
	}
	return peopleComment, nil
}

func (cmt CommentImplRepo) CountComment(ctx context.Context) (int, error) {
	query := COUNT_COMMENT
	rows := cmt.db.QueryRowContext(ctx, query)
	var count int
	err := rows.Scan(&count)
	if err != nil {
		log.Printf("[CountComment] failed to scan the data from the database, err: %v", err)
		return 0, err
	}
	return count, nil
}

func (cmt CommentImplRepo) UpdateComment(ctx context.Context, reqData models.CommentUserJoined, commentID uint64, userID uint64) error {
	query := UPDATE_COMMENT

	stmt, err := cmt.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[UpdateComment] failed to prepare the statement, err: %v", err)
		return err
	}
	_, err = stmt.ExecContext(ctx, reqData.Comment.Message, commentID, userID)
	if err != nil {
		log.Printf("[UpdateComment] failed to store data to the database, err: %v", err)
		return err
	}
	return nil
}

func (cmt CommentImplRepo) GetPhotoByID(ctx context.Context, commentID uint64, userID uint64) (*dto.EditPhotoResponse, error) {
	query := GET_PHOTO_BY_ID
	stmt, err := cmt.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[GetPhotoByID] failed to prepare the statement, err: %v", err)
		return nil, err
	}
	rows := stmt.QueryRowContext(ctx, commentID, userID)
	if err != nil {
		log.Printf("[GetPhotoByID] failed to query to the database, err: %v", err)
		return nil, err
	}
	personComment := models.CommentUserJoined{}
	err = rows.Scan(
		&personComment.Photo.PhotoID,
		&personComment.Photo.Title,
		&personComment.Photo.Caption,
		&personComment.Photo.PhotoUrl,
		&personComment.Photo.UserID,
		&personComment.Photo.UpdatedAt,
	)
	if err != nil {
		log.Printf("[GetPhotoByID] failed to scan the data from the database, err: %v", err)
		return nil, err
	}
	return dto.NewEditPhotoResponse(personComment.Photo, userID), err
}

func (cmt CommentImplRepo) CheckComment(ctx context.Context, commentID uint64, userID uint64) (bool, error) {
	query := CHECK_COMMENT
	stmt, err := cmt.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[CheckComment] failed to prepare the statement, err: %v", err)
		return false, err
	}
	rows, err := stmt.QueryContext(ctx, commentID, userID)
	if err != nil {
		log.Printf("[CheckComment] failed to query to the database, err: %v", err)
		return false, err
	}
	for rows.Next() {
		return true, nil
	}
	return false, nil
}

func (cmt CommentImplRepo) DeleteComment(ctx context.Context, commentID uint64, userID uint64) error {
	query := DELETE_COMMENT

	stmt, err := cmt.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[DeleteComment] failed to prepare the statement, err: %v", err)
		return err
	}

	_, err = stmt.QueryContext(ctx, commentID, userID)
	if err != nil {
		log.Printf("[DeleteComment] failed to delete the comment, err: %v", err)
		return err
	}
	return nil
}
