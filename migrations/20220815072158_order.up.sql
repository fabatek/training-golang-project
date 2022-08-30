CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    sum_price REAL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);