CREATE TABLE IF NOT EXISTS  order_items (
    id VARCHAR(255) NOT NULL,
    product_id VARCHAR(255),
    quantity BIGINT,
    order_id VARCHAR(255),
    price REAL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);