-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    user_id serial PRIMARY KEY,
    username varchar(50) UNIQUE NOT NULL
);

CREATE TABLE topics (
    topic_id serial PRIMARY KEY,
    title varchar(255) NOT NULL,
    created_by integer REFERENCES users (user_id),
    created_at timestamp WITH time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH time zone,
    is_deleted boolean DEFAULT FALSE,
    content text NOT NULL
);

CREATE TABLE comments (
    comment_id serial PRIMARY KEY,
    topic_id integer REFERENCES topics (topic_id) ON DELETE CASCADE,
    parent_comment_id integer REFERENCES comments (comment_id),
    user_id integer REFERENCES users (user_id),
    content text NOT NULL,
    created_at timestamp WITH time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp WITH time zone,
    is_deleted boolean DEFAULT FALSE
);

CREATE TABLE likes (
    like_id serial PRIMARY KEY,
    comment_id integer REFERENCES comments (comment_id),
    liked_by integer REFERENCES users (user_id),
    liked_at timestamp WITH time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE last_read_comments (
    user_id integer REFERENCES users (user_id),
    comment_id integer REFERENCES comments (comment_id),
    topic_id integer REFERENCES topics (topic_id) ON DELETE CASCADE,
    PRIMARY KEY (user_id, comment_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS last_read_comments;

DROP TABLE IF EXISTS likes;

DROP TABLE IF EXISTS comments;

DROP TABLE IF EXISTS topics;

DROP TABLE IF EXISTS users;

-- +goose StatementEnd
