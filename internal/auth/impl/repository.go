package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
)

type AuthRepository interface {
	InsertUser(ctx context.Context, data models.User) error
	GetUserEmail(ctx context.Context, email string) (*models.User, error)
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
	GET_USER_EMAIL = "SELECT id, email, username FROM user WHERE email=?;"
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

	err := res.Scan(&user.UserID, &user.Email, &user.Username)
	if err != nil {
		log.Printf("[GetUserEmail] failed to scan the data: %v", err)
		return nil, err
	}
	return user, nil
}
