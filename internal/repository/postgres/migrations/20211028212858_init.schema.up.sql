CREATE TABLE users
(
    id         UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    username   VARCHAR(200) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP    NOT NULL DEFAULT NOW()
);

CREATE TABLE topics
(
    id          UUID PRIMARY KEY       DEFAULT gen_random_uuid(),
    user_id     UUID          NOT NULL,
    name        VARCHAR(300)  NOT NULL,
    description VARCHAR(1500) NOT NULL,
    created_at  TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP     NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET DEFAULT
);

CREATE TABLE comments
(
    id         UUID PRIMARY KEY       DEFAULT gen_random_uuid(),
    topic_id   UUID          NOT NULL,
    user_id    UUID          NOT NULL,
    content    VARCHAR(1800) NOT NULL,
    created_at TIMESTAMP     NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP     NOT NULL DEFAULT NOW(),
    FOREIGN KEY (topic_id) REFERENCES topics (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET DEFAULT
);

create index topics__user_id_idx on topics using btree(user_id);
create index comments__user_id_idx on comments using btree(user_id);
create index comments__topic_id_idx on comments using btree(topic_id);