-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    email text NOT NULL

);

-- +goose Down
DROP TABLE users;
