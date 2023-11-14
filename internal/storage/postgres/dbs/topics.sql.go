// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: topics.sql

package dbs

import (
	"context"
	"time"
)

const topics = `-- name: Topics :many
SELECT t.topic_id,
       t.title,
       t.content,
       t.created_at,
       t.created_by,
       u.username AS author_username
FROM topics t
         INNER JOIN users u ON (t.created_by = u.user_id)
ORDER BY t.topic_id DESC
OFFSET $1::BIGINT LIMIT $2::BIGINT
`

type TopicsParams struct {
	Offset int64
	Limit  int64
}

type TopicsRow struct {
	TopicID        int64
	Title          string
	Content        string
	CreatedAt      time.Time
	CreatedBy      int64
	AuthorUsername string
}

func (q *Queries) Topics(ctx context.Context, arg TopicsParams) ([]TopicsRow, error) {
	rows, err := q.query(ctx, q.topicsStmt, topics, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TopicsRow
	for rows.Next() {
		var i TopicsRow
		if err := rows.Scan(
			&i.TopicID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.AuthorUsername,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const topicsNew = `-- name: TopicsNew :one
INSERT INTO topics (title, content, created_at, created_by, updated_at, updated_by)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING topic_id
`

type TopicsNewParams struct {
	Title     string
	Content   string
	CreatedAt time.Time
	CreatedBy int64
	UpdatedAt time.Time
	UpdatedBy int64
}

func (q *Queries) TopicsNew(ctx context.Context, arg TopicsNewParams) (int64, error) {
	row := q.queryRow(ctx, q.topicsNewStmt, topicsNew,
		arg.Title,
		arg.Content,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
	var topic_id int64
	err := row.Scan(&topic_id)
	return topic_id, err
}