-- name: CreateUser :one
INSERT INTO users (id, email) VALUES (?, ?) RETURNING *;
