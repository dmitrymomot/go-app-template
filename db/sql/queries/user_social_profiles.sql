-- name: CreateUserSocialProfile :exec
-- CreateUserSocialProfile: Link a user to a social profile
INSERT INTO user_external_profiles (user_id, provider_id, provider_type, external_account_id) VALUES (?, ?, ?, ?);

-- name: GetUserSocialProfileBySocialID :one
-- GetUserSocialProfileBySocialID: Get a user's social profile by social id and social name
SELECT * FROM user_external_profiles
WHERE external_account_id = ?
AND provider_type = ?
AND provider_id = ?;

-- name: GetUserSocialProfileByUserID :one
-- GetUserSocialProfileByUserID: Get a user's social profile by user id and social name
SELECT * FROM user_external_profiles
WHERE user_id = ?
AND provider_type = ?
AND provider_id = ?;

-- name: GetUserSocialProfilesByUserID :many
-- GetUserSocialProfilesByUserID: Get a user's social profiles by user id
SELECT * FROM user_external_profiles WHERE user_id = ? ORDER BY created_at DESC;

-- name: DeleteUserSocialProfileBySocialID :exec
-- DeleteUserSocialProfileBySocialID: Delete a user's social profile by social id and social name
DELETE FROM user_external_profiles
WHERE external_account_id = ?
AND provider_type = ?
AND provider_id = ?;

-- name: DeleteUserSocialProfilesByUserID :exec
-- DeleteUserSocialProfilesByUserID: Delete a user's social profiles by user id
DELETE FROM user_external_profiles WHERE user_id = ?;
