-- +goose Up
-- +goose StatementBegin 
CREATE TABLE users (
    pk SERIAL PRIMARY KEY,
    user_id uuid NOT NULL UNIQUE, 
    email TEXT NOT NULL UNIQUE,
    "password" TEXT NOT NULL

);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd


