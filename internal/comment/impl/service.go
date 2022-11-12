package impl

import (
	"context"
	"log"

	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/dto"
	"github.com/maheswaradevo/hacktiv8-finalproject2/pkg/errors"
)

type CommentServiceImpl struct {
	repo CommentRepository
}

func ProvideCommentService(repo CommentRepository) *CommentServiceImpl {
	return &CommentServiceImpl{
		repo: repo,
	}
}

func (cmt *CommentServiceImpl) CreateComment(ctx context.Context, data *dto.CreateCommentRequest, userID uint64) (res *dto.CreateCommentResponse, err error) {
	commentData := data.ToCommentEntity()
	commentData.UserID = userID
	commentID, err := cmt.repo.CreateComment(ctx, *commentData)
	if err != nil {
		log.Printf("[CreateComment] failed to store user data to database: %v", err)
		return
	}
	return dto.NewCommentCreateResponse(*commentData, userID, commentID), nil
}

func (cmt *CommentServiceImpl) ViewComment(ctx context.Context) (dto.ViewCommentsResponse, error) {
	count, err := cmt.repo.CountComment(ctx)
	if err != nil {
		log.Printf("[ViewComment] failed to count the comment, err: %v", err)
		return nil, err
	}
	if count == 0 {
		err = errors.ErrDataNotFound
		log.Printf("[ViewComment] no data exists in the database: %v", err)
		return nil, err
	}
	res, err := cmt.repo.ViewComment(ctx)
	if err != nil {
		log.Printf("[ViewComment] failed to view the comment, err: %v", err)
		return nil, err
	}
	return dto.NewViewCommentsResponse(res), nil
}

func (cmt *CommentServiceImpl) UpdateComment(ctx context.Context, commentID uint64, userID uint64, data *dto.EditCommentRequest) (*dto.EditPhotoResponse, error) {
	editedComment := data.ToCommentEntity()

	check, err := cmt.repo.CheckComment(ctx, commentID, userID)
	if err != nil {
		log.Printf("[UpdateComment] failed to check comment with, userID: %v, err: %v", userID, err)
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		log.Printf("[UpdateComment] no comment in userID: %v", userID)
		return nil, err
	}
	err = cmt.repo.UpdateComment(ctx, *editedComment, commentID, userID)
	if err != nil {
		log.Printf("[UpdateComment] failed to update comment, err: %v", err)
		return nil, err
	}
	photo, err := cmt.repo.GetPhotoByID(ctx, commentID, userID)
	if err != nil {
		log.Printf("[UpdateComment] failed to get photo, err: %v", err)
		return nil, err
	}
	return photo, nil
}

func (cmt *CommentServiceImpl) DeleteComment(ctx context.Context, commentID uint64, userID uint64) (*dto.DeleteCommentResponse, error) {
	check, err := cmt.repo.CheckComment(ctx, commentID, userID)
	if err != nil {
		log.Printf("[DeleteComment] failed to check comment with, userID: %v, err: %v", userID, err)
		return nil, err
	}
	if !check {
		err = errors.ErrDataNotFound
		log.Printf("[DeleteComment] no comment in userID: %v", userID)
		return nil, err
	}

	err = cmt.repo.DeleteComment(ctx, commentID, userID)
	if err != nil {
		log.Printf("[DeleteComment] failed to delete comment, id: %v", commentID)
		return nil, err
	}
	message := "Your comment has been successfully deleted"
	return dto.NewDeleteCommentResponse(message), nil
}
