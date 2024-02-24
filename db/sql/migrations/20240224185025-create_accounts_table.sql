
-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS accounts (
  id CHARACTER(36) PRIMARY KEY NOT NULL CHECK(length(id) == 36),
  name VARCHAR(50) NOT NULL CHECK(length(name) <= 50 AND length(name) >= 5),
  title VARCHAR(255) DEFAULT NULL CHECK(length(title) <= 255),
  logo_url VARCHAR(255) DEFAULT NULL CHECK(length(logo_url) <= 255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS accounts_name_uindex ON accounts (name);
CREATE INDEX IF NOT EXISTS accounts_created_at_index ON accounts (created_at);
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE accounts;
