-- name: LikesNew :one
INSERT INTO likes (comment_id, active, created_at, created_by, updated_at, updated_by)
VALUES (@comment_id, @active, @created_at, @created_by, @updated_at, @updated_by)
ON CONFLICT (comment_id, created_by) 
DO UPDATE SET 
    active = NOT likes.active, 
    updated_at = EXCLUDED.updated_at, 
    updated_by = EXCLUDED.updated_by
RETURNING comment_id, created_by;


-- name: LikesByCommentID :many
SELECT likes.*, users.*
FROM likes
JOIN users ON likes.created_by = users.user_id
WHERE likes.comment_id = @comment_id AND likes.active = TRUE;
