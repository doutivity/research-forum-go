package domain

import "time"

type Like struct {
	CreatedAt  time.Time
	LikeAuthor *LikeAuthor
}

type LikeAuthor struct {
	ID       int64
	Username string
}

type LikeCreate struct {
	CommentID  int64
	LikeAuthor *LikeAuthor
}
