
-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS accounts (
  id CHARACTER(36) PRIMARY KEY NOT NULL CHECK(length(id) == 36),
  name VARCHAR(255) DEFAULT NULL CHECK(length(name) <= 255),
  slug VARCHAR(50) NOT NULL CHECK(length(slug) <= 50 AND length(slug) >= 5),
  logo_url VARCHAR(255) DEFAULT NULL CHECK(length(logo_url) <= 255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS accounts_slug_uindex ON accounts (slug);
CREATE INDEX IF NOT EXISTS accounts_created_at_index ON accounts (created_at);
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE accounts;
