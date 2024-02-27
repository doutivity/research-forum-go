-- name: CommentsNew :one
INSERT INTO comments (parent_comment_id, topic_id, content, created_at, created_by, updated_at, updated_by)
VALUES (@parent_comment_id, @topic_id, @content, @created_at, @created_by, @updated_at, @updated_by)
RETURNING comment_id;

-- name: CommentsByTopic :many
SELECT c.comment_id,
       c.parent_comment_id,
       c.content,
       c.created_by,
       c.created_at,
       u.username AS author_username
FROM comments c
         INNER JOIN users u ON c.created_by = u.user_id
WHERE c.topic_id = @topic_id::BIGINT
OFFSET sqlc.arg('offset') LIMIT sqlc.arg('limit');

-- name: CommentsByID :one
SELECT c.comment_id,
       c.parent_comment_id,
       c.content,
       c.created_by,
       c.created_at,
       u.username AS author_username
FROM comments c
         INNER JOIN users u ON c.created_by = u.user_id
WHERE c.comment_id = @comment_id;

-- name: CommentUpdate :exec
UPDATE comments
SET content    = @content,
    updated_at = @updated_at,
    updated_by = @updated_by
WHERE comment_id = @comment_id;
