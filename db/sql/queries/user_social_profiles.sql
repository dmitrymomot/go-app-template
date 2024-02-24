-- name: CreateUserSocialProfile :exec
-- CreateUserSocialProfile: Link a user to a social profile
INSERT INTO user_social_profiles (user_id, social_id, social_name) VALUES (@user_id, @social_id, @social_name);

-- name: GetUserSocialProfileBySocialID :one
-- GetUserSocialProfileBySocialID: Get a user's social profile by social id and social name
SELECT * FROM user_social_profiles WHERE social_id = @social_id AND social_name = @social_name;

-- name: GetUserSocialProfilesByUserID :many
-- GetUserSocialProfilesByUserID: Get a user's social profiles by user id
SELECT * FROM user_social_profiles WHERE user_id = @user_id ORDER BY created_at DESC;

-- name: DeleteUserSocialProfileBySocialID :exec
-- DeleteUserSocialProfileBySocialID: Delete a user's social profile by social id and social name
DELETE FROM user_social_profiles WHERE social_id = @social_id AND social_name = @social_name;

-- name: DeleteUserSocialProfilesByUserID :exec
-- DeleteUserSocialProfilesByUserID: Delete a user's social profiles by user id
DELETE FROM user_social_profiles WHERE user_id = @user_id;
