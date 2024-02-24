
-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS account_members (
  id CHARACTER(36) PRIMARY KEY NOT NULL CHECK(length(id) == 36),
  account_id CHARACTER(36) NOT NULL CHECK(length(account_id) == 36),
  user_id CHARACTER(36) NOT NULL CHECK(length(user_id) == 36),
  name VARCHAR(50) DEFAULT NULL CHECK(length(name) <= 100 AND length(name) >= 2),
  role VARCHAR(50) DEFAULT NULL CHECK(length(role) <= 50),
  avatar_url VARCHAR(255) DEFAULT NULL CHECK(length(avatar_url) <= 255),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (account_id) REFERENCES accounts (id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX account_members_account_id_user_id_uindex ON account_members (account_id, user_id);
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE account_members;
