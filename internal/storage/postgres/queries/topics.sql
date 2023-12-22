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
         INNER JOIN topic_last_update tlu ON (t.topic_id = tlu.topic_id)
ORDER BY tlu.last_updated_at DESC
OFFSET sqlc.arg('offset')::BIGINT LIMIT sqlc.arg('limit')::BIGINT;

-- name: TopicsWithUnreadCommentsNumber :many
SELECT
    t.topic_id,
    t.title,
    t.content,
    t.created_at,
    t.created_by,
    u.username AS author_username,
    (
        SELECT
            COUNT(*)
        FROM
            comments c
        WHERE
            c.topic_id = t.topic_id
            AND (lrc.comment_id IS NULL
                OR c.comment_id > lrc.comment_id)) AS unread_comments_count
FROM
    topics t
    INNER JOIN users u ON t.created_by = u.user_id
    INNER JOIN topic_last_update tlu ON t.topic_id = tlu.topic_id
    LEFT JOIN last_read_comments lrc ON t.topic_id = lrc.topic_id
        AND lrc.user_id = @read_by_user_id
    ORDER BY
        tlu.last_updated_at DESC OFFSET sqlc.arg ('offset')::bigint
LIMIT sqlc.arg ('limit')::bigint;

-- name: TopicsGetByID :one
SELECT t.topic_id,
       t.title,
       t.content,
       t.created_at,
       t.created_by,
       u.username AS author_username
FROM topics t
         INNER JOIN users u ON (t.created_by = u.user_id)
WHERE topic_id = @topic_id;

-- name: TopicsGetByIDWithLastReadComment :one
SELECT t.topic_id,
       t.title,
       t.content,
       t.created_at,
       t.created_by,
       u.username AS author_username,
       lrc.comment_id AS last_read_comment_id
FROM topics t
         INNER JOIN users u ON t.created_by = u.user_id
         LEFT JOIN last_read_comments lrc ON t.topic_id = lrc.topic_id AND lrc.user_id = @user_id
WHERE t.topic_id = @topic_id;
