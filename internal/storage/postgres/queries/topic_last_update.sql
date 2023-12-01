-- name: TopicLastUpdateNew :exec
INSERT INTO topic_last_update (topic_id, last_updated_at)
VALUES (@topic_id, @last_updated_at)
ON CONFLICT (topic_id) 
DO UPDATE SET last_updated_at = EXCLUDED.last_updated_at;
