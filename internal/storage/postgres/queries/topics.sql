-- name: TopicsNew :one
INSERT INTO topics (title, content, created_at, created_by, updated_at, updated_by)
VALUES (@title, @content, @created_at, @created_by, @updated_at, @updated_by)
RETURNING topic_id;

-- name: Topics :many
SELECT t.topic_id,
       t.title,
       t.content,
       t.created_at,
       t.created_by,
       u.username AS author_username
FROM topics t
         INNER JOIN users u ON (t.created_by = u.user_id)
ORDER BY t.topic_id DESC
OFFSET sqlc.arg('offset')::BIGINT LIMIT sqlc.arg('limit')::BIGINT;

-- name: TopicsGetByID :one
SELECT t.topic_id,
       t.title,
       t.content,
       t.created_at,
       t.created_by,
       u.username AS author_username
FROM topics t
         INNER JOIN users u ON (t.created_by = u.user_id)
WHERE topic_id = sqlc.arg('topic_id')::BIGINT;

