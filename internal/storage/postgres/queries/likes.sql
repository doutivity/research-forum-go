-- name: LikesUpsert :exec
INSERT INTO likes (comment_id, active, created_at, created_by, updated_at, updated_by)
VALUES (@comment_id, @active, @created_at, @created_by, @updated_at, @updated_by)
ON CONFLICT (comment_id, created_by)
    DO UPDATE SET active     = EXCLUDED.active,
                  updated_at = EXCLUDED.updated_at,
                  updated_by = EXCLUDED.updated_by;

-- name: LikesByCommentIDs :many
SELECT likes.comment_id,
       users.user_id,
       users.username,
       likes.created_at
FROM likes
         INNER JOIN users ON likes.created_by = users.user_id
WHERE likes.comment_id = ANY (@comment_ids::BIGINT[])
  AND likes.active = TRUE;
