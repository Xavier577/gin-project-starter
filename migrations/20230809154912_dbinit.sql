
-- +goose Up
CREATE TABLE users (
   id uuid NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
   email text NOT NULL UNIQUE,
   password text
);


-- +goose Down
DROP TABLE users;

