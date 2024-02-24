
-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE user_external_profiles (
	user_id CHARACTER(36) NOT NULL,
	provider_id VARCHAR(255) NOT NULL,
	provider_type VARCHAR(50) NOT NULL,
	external_account_id VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE UNIQUE INDEX user_external_profiles_user_id_social_id_uindex
	ON user_external_profiles (user_id, external_account_id);
-- +migrate StatementEnd

-- +migrate Down
DROP TABLE user_external_profiles;
