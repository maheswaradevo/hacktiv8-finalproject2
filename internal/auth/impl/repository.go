package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/errors"
)

type AuthRepository interface {
	InsertUser(ctx context.Context, data models.User) error
	GetUserEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, userID uint64, data models.User) error
}
type AuthImpl struct {
	db *sql.DB
}

func ProvideAuthRepository(db *sql.DB) *AuthImpl {
	return &AuthImpl{
		db: db,
	}
}

var (
	INSERT_USER    = "INSERT INTO `user`(email, password, age, username) VALUES (?, ?, ?, ?);"
	GET_USER_EMAIL = "SELECT id, email, username, password FROM user WHERE email=?;"
	UPDATE_USER    = "UPDATE `user` SET email = ?, username = ?, age = ? WHERE id = ?;"
)

func (auth *AuthImpl) InsertUser(ctx context.Context, data models.User) error {
	query := INSERT_USER
	stmt, err := auth.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[InserUser] failed to prepare the statement: %v", err)
		return err
	}
	_, err = stmt.ExecContext(ctx, data.Email, data.Password, data.Age, data.Username)
	if err != nil {
		log.Printf("[InserUser] failed to insert user to the database: %v", err)
		return err
	}
	return nil
}

func (auth *AuthImpl) GetUserEmail(ctx context.Context, email string) (*models.User, error) {
	query := GET_USER_EMAIL
	res := auth.db.QueryRowContext(ctx, query, email)
	user := &models.User{}

	err := res.Scan(&user.UserID, &user.Email, &user.Username, &user.Password)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("[GetUserEmail] failed to scan the data: %v", err)
		return nil, err
	} else if err == sql.ErrNoRows {
		log.Printf("[GetUserEmail] no data existed in the database\n")
		return nil, errors.ErrInvalidResources
	}
	return user, nil
}

func (auth *AuthImpl) UpdateUser(ctx context.Context, userID uint64, data models.User) error {
	query := UPDATE_USER
	stmt, err := auth.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("[UpdateUser] failed to prepare the statement: %v", err)
		return err
	}
	_, err = stmt.ExecContext(ctx, data.Email, data.Username, data.Age, userID)
	if err != nil {
		log.Printf("[UpdateUser] failed to insert the data to the database, err: %v", err)
		return err
	}
	return nil
}
