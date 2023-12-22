-- +goose Up
-- +goose StatementBegin
CREATE TABLE topic_comment_stats
(
    topic_id      BIGINT NOT NULL PRIMARY KEY,
    comment_count BIGINT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE topic_comment_stats;
-- +goose StatementEnd
