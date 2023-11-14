# DOU Forum on PostgreSQL
- [Збереження стану онлайну користувача в Redis](https://dou.ua/forums/topic/35260/)
- [Hash, Set чи Sorted set. Який тип даних вибрати для збереження стану онлайну користувача в Redis?](https://dou.ua/forums/topic/44655/)
- [Batch UPDATE в PostgreSQL](https://dou.ua/forums/topic/35261/)

# Support Ukraine 🇺🇦
- Help Ukraine via [SaveLife fund](https://savelife.in.ua/en/donate-en/)
- Help Ukraine via [Dignitas fund](https://dignitas.fund/donate/)
- Help Ukraine via [National Bank of Ukraine](https://bank.gov.ua/en/news/all/natsionalniy-bank-vidkriv-spetsrahunok-dlya-zboru-koshtiv-na-potrebi-armiyi)
- More info on [war.ukraine.ua](https://war.ukraine.ua/) and [MFA of Ukraine](https://twitter.com/MFA_Ukraine)

# Testing
```bash
make env-up
make docker-go-version
make docker-pg-version
make migrate-up
make go-test
make env-down
```

# Schema
```sql
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
```
