-- +goose Up
-- +goose StatementBegin
CREATE TABLE topic_last_update
(
    topic_id   BIGINT                NOT NULL REFERENCES topics (topic_id) PRIMARY KEY,
    last_updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE topic_last_update;
-- +goose StatementEnd
