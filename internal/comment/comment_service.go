package comment

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/comment/impl"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
)

type CommentService interface {
	CreateComment(ctx context.Context, data *dto.CreateCommentRequest, userID uint64) (res *dto.CreateCommentResponse, err error)
	ViewComment(ctx context.Context) (dto.ViewCommentsResponse, error)
	UpdateComment(ctx context.Context, commentID uint64, userID uint64, data *dto.EditCommentRequest) (*dto.EditPhotoResponse, error)
	DeleteComment(ctx context.Context, commentID uint64, userID uint64) (*dto.DeleteCommentResponse, error)
}

func ProvideCommentService(db *sql.DB) CommentService {
	repo := impl.ProvideCommentRepository(db)
	return impl.ProvideCommentService(repo)
}
