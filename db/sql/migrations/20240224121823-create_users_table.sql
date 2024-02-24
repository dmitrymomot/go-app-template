
-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id CHARACTER(36) PRIMARY KEY NOT NULL CHECK(length(id) == 36),
  email VARCHAR(255) NOT NULL CHECK(length(email) <= 255 AND length(email) >= 6),
  password VARCHAR(255) DEFAULT NULL CHECK(length(password) <= 255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  verified_at TIMESTAMP DEFAULT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS users_email_uindex ON users (email);
CREATE INDEX IF NOT EXISTS users_created_at_index ON users (created_at);
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE users;
