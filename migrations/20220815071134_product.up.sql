CREATE TABLE IF NOT EXISTS products (
    id VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    price REAL,
    quantity BIGINT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);