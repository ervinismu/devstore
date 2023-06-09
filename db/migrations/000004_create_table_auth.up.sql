CREATE TABLE IF NOT EXISTS auths (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    token VARCHAR NOT NULL,
    auth_type VARCHAR NOT NULL,
    expired_at TIMESTAMPTZ NOT NULL,
    user_id BIGINT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)
);
