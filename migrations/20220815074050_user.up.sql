CREATE TABLE IF NOT EXISTS  users (
    id VARCHAR(255) NOT NULL,
    name VARCHAR(255),
    email VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);