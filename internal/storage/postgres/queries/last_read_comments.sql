-- name: LastReadCommentsNew :exec
INSERT INTO last_read_comments (user_id, topic_id, comment_id)
    VALUES (@user_id, @topic_id, @comment_id)
ON CONFLICT (user_id, topic_id)
    DO UPDATE SET
        comment_id = EXCLUDED.comment_id;

