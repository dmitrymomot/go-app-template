-- name: CreateAccount :exec
-- CreateAccount: creates an account for a user
INSERT INTO accounts (id, name, slug, logo_url) VALUES (@id, @name, @slug, @logo_url);

-- name: GetAccount :one
-- GetAccount: retrieves an account by its id
SELECT * FROM accounts WHERE id = @id;

-- name: GetAccountBySlug :one
-- GetAccountBySlug: retrieves an account by its slug
SELECT * FROM accounts WHERE slug = @slug;

-- name: UpdateAccount :exec
-- UpdateAccount: updates an account
UPDATE accounts SET name = @name, slug = @slug, logo_url = @logo_url WHERE id = @id;

-- name: DeleteAccount :exec
-- DeleteAccount: deletes an account
DELETE FROM accounts WHERE id = @id;
