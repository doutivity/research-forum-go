package domain

import "time"

type Like struct {
	CommentID  int64
	CreatedAt  time.Time
	LikeAuthor *LikeAuthor
}

type LikeAuthor struct {
	ID       int64
	Username string
}

type LikeUpdate struct {
	CommentID  int64
	LikeAuthor *LikeAuthor
}
