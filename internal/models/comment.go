package models

import "time"

type Comment struct {
	CommentID uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	PhotoID   uint64    `db:"photo_id"`
	Message   string    `db:"message"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CommentUserJoined struct {
	Comment Comment
	Photo   Photo
	User    User
}

type PeopleCommentJoined []*CommentUserJoined
type Comments []*Comment
