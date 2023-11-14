-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    user_id  BIGSERIAL NOT NULL PRIMARY KEY,
    username VARCHAR   NOT NULL UNIQUE
);

CREATE TABLE topics
(
    topic_id   BIGSERIAL                NOT NULL PRIMARY KEY,
    title      VARCHAR                  NOT NULL,
    content    VARCHAR                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by BIGINT                   NOT NULL REFERENCES users (user_id),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by BIGINT                   NOT NULL REFERENCES users (user_id),
    deleted_at TIMESTAMP WITH TIME ZONE NULL,
    deleted_by BIGINT                   NULL REFERENCES users (user_id)
);

CREATE TABLE comments
(
    comment_id        BIGSERIAL                NOT NULL PRIMARY KEY,
    parent_comment_id BIGINT                   NULL REFERENCES comments (comment_id),
    topic_id          BIGINT                   NOT NULL REFERENCES topics (topic_id),
    content           VARCHAR                  NOT NULL,
    created_at        TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by        BIGINT                   NOT NULL REFERENCES users (user_id),
    updated_at        TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by        BIGINT                   NOT NULL REFERENCES users (user_id),
    deleted_at        TIMESTAMP WITH TIME ZONE NULL,
    deleted_by        BIGINT                   NULL REFERENCES users (user_id)
);

CREATE TABLE likes
(
    comment_id BIGINT                   NOT NULL REFERENCES comments (comment_id),
    active     BOOLEAN                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_by BIGINT                   NOT NULL REFERENCES users (user_id),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_by BIGINT                   NOT NULL REFERENCES users (user_id),
    PRIMARY KEY (comment_id, created_by)
);

CREATE TABLE last_read_comments
(
    user_id    BIGINT NOT NULL REFERENCES users (user_id),
    topic_id   BIGINT NOT NULL REFERENCES topics (topic_id),
    comment_id BIGINT NOT NULL REFERENCES comments (comment_id),
    PRIMARY KEY (user_id, topic_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE last_read_comments;
DROP TABLE likes;
DROP TABLE comments;
DROP TABLE topics;
DROP TABLE users;
-- +goose StatementEnd
