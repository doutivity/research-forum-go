// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package dbs

import (
	"database/sql"
	"time"
)

type Comment struct {
	CommentID       int64
	ParentCommentID sql.NullInt64
	TopicID         int64
	Content         string
	CreatedAt       time.Time
	CreatedBy       int64
	UpdatedAt       time.Time
	UpdatedBy       int64
	DeletedAt       sql.NullTime
	DeletedBy       sql.NullInt64
}

type LastReadComment struct {
	UserID    int64
	TopicID   int64
	CommentID int64
}

type Like struct {
	CommentID int64
	Active    bool
	CreatedAt time.Time
	CreatedBy int64
	UpdatedAt time.Time
	UpdatedBy int64
}

type Topic struct {
	TopicID   int64
	Title     string
	Content   string
	CreatedAt time.Time
	CreatedBy int64
	UpdatedAt time.Time
	UpdatedBy int64
	DeletedAt sql.NullTime
	DeletedBy sql.NullInt64
}

type TopicCommentStat struct {
	TopicID      int64
	CommentCount int64
}

type TopicLastUpdate struct {
	TopicID       int64
	LastUpdatedAt time.Time
}

type User struct {
	UserID   int64
	Username string
}
