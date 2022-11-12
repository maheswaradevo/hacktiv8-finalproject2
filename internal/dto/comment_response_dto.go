package dto

import (
	"time"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"
)

type CreateCommentResponse struct {
	CommentID uint64    `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint64    `json:"photo_id"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type ViewCommentResponse struct {
	CommentID uint64                   `json:"id"`
	Message   string                   `json:"message"`
	PhotoID   uint64                   `json:"photo_id"`
	UserID    uint64                   `json:"user_id"`
	UpdatedAt time.Time                `json:"updated_at"`
	CreatedAt time.Time                `json:"created_at"`
	User      ViewCommentUserResponse  `json:"user"`
	Photo     ViewCommentPhotoResponse `json:"photo"`
}

type ViewCommentUserResponse struct {
	UserID   uint64 `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ViewCommentPhotoResponse struct {
	PhotoID  uint64 `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   uint64 `json:"user_id"`
}

type ViewCommentsResponse []*ViewCommentResponse

type EditCommentResponse struct {
	PhotoID   uint64    `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint64    `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}

func NewCommentCreateResponse(cmt models.Comment, userID uint64, commentID uint64) *CreateCommentResponse {
	return &CreateCommentResponse{
		CommentID: commentID,
		Message:   cmt.Message,
		PhotoID:   cmt.PhotoID,
		UserID:    userID,
		CreatedAt: time.Now(),
	}
}

func NewViewCommentResponse(cmt models.CommentUserJoined) *ViewCommentResponse {
	return &ViewCommentResponse{
		CommentID: cmt.Comment.CommentID,
		Message:   cmt.Comment.Message,
		PhotoID:   cmt.Photo.PhotoID,
		UserID:    cmt.User.UserID,
		UpdatedAt: cmt.Comment.UpdatedAt,
		CreatedAt: cmt.Comment.CreatedAt,
		User: ViewCommentUserResponse{
			UserID:   cmt.User.UserID,
			Email:    cmt.User.Email,
			Username: cmt.User.Username,
		},
		Photo: ViewCommentPhotoResponse{
			PhotoID:  cmt.Photo.PhotoID,
			Title:    cmt.Photo.Title,
			Caption:  cmt.Photo.Caption,
			PhotoUrl: cmt.Photo.PhotoUrl,
			UserID:   cmt.Photo.UserID,
		},
	}
}

func NewViewCommentsResponse(cmt models.PeopleCommentJoined) ViewCommentsResponse {
	var viewCommentsResponse ViewCommentsResponse

	for idx := range cmt {
		peopleComment := NewViewCommentResponse(*cmt[idx])
		viewCommentsResponse = append(viewCommentsResponse, peopleComment)
	}
	return viewCommentsResponse
}

func NewEditCommentResponse(ph models.CommentUserJoined, userID uint64) *EditCommentResponse {
	return &EditCommentResponse{
		PhotoID:   ph.Photo.PhotoID,
		Title:     ph.Photo.Title,
		Caption:   ph.Photo.Caption,
		PhotoUrl:  ph.Photo.PhotoUrl,
		UserID:    userID,
		UpdatedAt: time.Now(),
	}
}

func NewDeleteCommentResponse(message string) *DeleteCommentResponse {
	return &DeleteCommentResponse{
		Message: message,
	}
}
