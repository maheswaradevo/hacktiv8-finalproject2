package dto

import "github.com/maheswaradevo/hacktiv8-finalproject2/internal/models"

type CreateCommentRequest struct {
	Message string `json:"message"`
	PhotoID uint64 `json:"photo_id"`
}

func (dto *CreateCommentRequest) ToCommentEntity() (cmt *models.Comment) {
	cmt = &models.Comment{
		Message: dto.Message,
		PhotoID: dto.PhotoID,
	}
	return
}

type EditCommentRequest struct {
	Message string `json:"message"`
}

func (dto *EditCommentRequest) ToCommentEntity() *models.CommentUserJoined {
	return &models.CommentUserJoined{
		Comment: models.Comment{
			Message: dto.Message,
		},
	}
}
