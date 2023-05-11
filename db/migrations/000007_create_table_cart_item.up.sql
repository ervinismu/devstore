CREATE TABLE IF NOT EXISTS cart_items (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    quantity INT,
    base_price INT,
    total_price INT,

    cart_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,

    FOREIGN KEY(cart_id) REFERENCES carts(id),
    FOREIGN KEY(product_id) REFERENCES products(id),

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);
