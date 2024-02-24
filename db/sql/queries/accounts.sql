-- name: CreateAccount :exec
-- CreateAccount: creates an account for a user
INSERT INTO accounts (id, name, title, logo_url) VALUES (@id, @name, @title, @logo_url);

-- name: GetAccount :one
-- GetAccount: retrieves an account by its id
SELECT * FROM accounts WHERE id = @id;

-- name: GetAccountByName :one
-- GetAccountByName: retrieves an account by its name
SELECT * FROM accounts WHERE name = @name;

-- name: UpdateAccount :exec
-- UpdateAccount: updates an account
UPDATE accounts SET name = @name, title = @title, logo_url = @logo_url WHERE id = @id;

-- name: DeleteAccount :exec
-- DeleteAccount: deletes an account
DELETE FROM accounts WHERE id = @id;
