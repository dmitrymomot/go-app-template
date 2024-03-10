-- name: CreateUser :one
-- CreateUser: Create a new user in the database
INSERT INTO users (id, email, password, created_at) VALUES (@id, @email, @password, @created_at) RETURNING *;

-- name: GetUserByEmail :one
-- GetUserByEmail: Get a user by email
SELECT * FROM users WHERE email = @email;

-- name: GetUserByID :one
-- GetUserByID: Get a user by id
SELECT * FROM users WHERE id = @id;

-- name: UpdateUserPassword :exec
-- UpdateUserPassword: Update a user's password
UPDATE users SET password = @password WHERE id = @id;

-- name: UpdateUserEmail :exec
-- UpdateUserEmail: Update a user's email
UPDATE users SET email = @email, verified_at = NULL WHERE id = @id;

-- name: VerifyUser :exec
-- VerifyUser: Verify a user's email
UPDATE users SET verified_at = @verified_at WHERE id = @id;

-- name: DeleteUser :exec
-- DeleteUser: Delete a user from the database
DELETE FROM users WHERE id = @id;

-- name: GetUsers :many
-- GetUsers: Get user list with pagination
SELECT * FROM users ORDER BY created_at DESC LIMIT @limit_val OFFSET @offset_val;

-- name: GetUsersCount :one
-- GetUsersCount: Get total user count
SELECT COUNT(id) FROM users;
