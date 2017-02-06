
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id SERIAL NOT NULL PRIMARY KEY,
  email TEXT NOT NULL DEFAULT ''
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
