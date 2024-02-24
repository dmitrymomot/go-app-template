-- name: GetUserAccounts :many
-- GetUserAccounts: retrieves accounts list of a user with pagination
SELECT sqlc.embed(accounts), sqlc.embed(account_members)
FROM account_members
JOIN accounts ON accounts.id = account_members.account_id
WHERE account_members.user_id = @user_id
ORDER BY account_members.created_at DESC
LIMIT @limit_val OFFSET @offset_val;
