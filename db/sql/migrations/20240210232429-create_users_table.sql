
-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    email VARCHAR(255) NOT NULL CHECK (email <> ''),
	password bytea NOT NULL CHECK (password <> ''),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS users_email_idx ON users (email);
-- +migrate StatementEnd


-- +migrate Down
-- +migrate StatementBegin
DROP TABLE IF EXISTS users;
-- +migrate StatementEnd
