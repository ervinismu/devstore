CREATE TABLE IF NOT EXISTS carts (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    total_price INT,

    user_id BIGINT NOT NULL,

    FOREIGN KEY(user_id) REFERENCES users(id),

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
