
-- +migrate Up
CREATE TABLE authors (
  id   INTEGER PRIMARY KEY,
  name text    NOT NULL,
  bio  text
);

-- +migrate Down
DROP TABLE examples;
