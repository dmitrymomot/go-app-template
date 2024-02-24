
-- +migrate Up
CREATE TABLE user_social_profiles (
	user_id CHARACTER(36) NOT NULL REFERENCES users(id),
	social_id VARCHAR(255) NOT NULL,
	social_name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX user_social_profiles_user_id_social_id_uindex ON user_social_profiles (user_id, social_id);

-- +migrate Down
DROP TABLE user_social_profiles;
