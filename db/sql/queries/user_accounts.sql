-- name: GetUserAccounts :many
-- GetUserAccounts: retrieves accounts list of a user with pagination
SELECT sqlc.embed(accounts), sqlc.embed(account_members)
FROM account_members
JOIN accounts ON accounts.id = account_members.account_id
WHERE account_members.user_id = ?
ORDER BY account_members.created_at DESC
LIMIT ? OFFSET ?;

-- name: GetAccountUsers :many
-- GetAccountUsers: retrieves users list of an account with pagination
SELECT sqlc.embed(users), sqlc.embed(account_members)
FROM account_members
JOIN users ON users.id = account_members.user_id
WHERE account_members.account_id = ?
ORDER BY account_members.created_at DESC
LIMIT ? OFFSET ?;
