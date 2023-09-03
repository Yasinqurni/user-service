-- migrate:up
-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT ,
    address CHAR(100) NOT NULL,
    name CHAR(36) NOT NULL,
    email CHAR(36) NOT NULL,
    created_at TIMESTAMP NULL ,
    updated_at TIMESTAMP NULL ,
    deleted_at TIMESTAMP NULL , PRIMARY KEY (id)
);

-- migrate:down
-- +goose Down
DROP TABLE users;
