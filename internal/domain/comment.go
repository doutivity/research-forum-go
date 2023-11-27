package domain

import "time"

type CommentCreate struct {
	ParentCommentID *int64
	Content         string
	TopicID         int64
}

type CommentUpdate struct {
	ID      int64
	Content string
}

type CommentAuthor struct {
	ID       int64
	Username string
}

type Comment struct {
	ID              int64
	ParentCommentID *int64
	Content         string
	CreatedAt       time.Time
	Author          *CommentAuthor
	Likes           []*Like
}
