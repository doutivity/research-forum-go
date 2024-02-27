// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: topic_last_update.sql

package dbs

import (
	"context"
	"time"
)

const topicLastUpdateNew = `-- name: TopicLastUpdateNew :exec
INSERT INTO topic_last_update (topic_id, last_updated_at)
    VALUES ($1, $2)
ON CONFLICT (topic_id)
    DO UPDATE SET
        last_updated_at = EXCLUDED.last_updated_at
`

type TopicLastUpdateNewParams struct {
	TopicID       int64
	LastUpdatedAt time.Time
}

func (q *Queries) TopicLastUpdateNew(ctx context.Context, arg TopicLastUpdateNewParams) error {
	_, err := q.exec(ctx, q.topicLastUpdateNewStmt, topicLastUpdateNew, arg.TopicID, arg.LastUpdatedAt)
	return err
}
