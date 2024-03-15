-- name: CreateAccountMember :exec
-- CreateAccountMember: creates a member for an account
INSERT INTO account_members (id, account_id, user_id, name, role, avatar_url) VALUES (?, ?, ?, ?, ?, ?);

-- name: GetAccountMemberByID :one
-- GetAccountMember: retrieves a member of an account by its id
SELECT * FROM account_members WHERE id = ?;

-- name: GetAccountMembers :many
-- GetAccountMembers: retrieves members list of an account with pagination
SELECT * FROM account_members
WHERE account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetAccountMemberByUserID :one
-- GetAccountMemberByUserID: retrieves a member of an account by its user_id
SELECT * FROM account_members WHERE account_id = ? AND user_id = ?;

-- name: UpdateAccountMember :exec
-- UpdateAccountMember: updates a member of an account
UPDATE account_members SET name = ?, role = ?, avatar_url = ? WHERE id = ?;

-- name: DeleteAccountMember :exec
-- DeleteAccountMember: deletes a member of an account
DELETE FROM account_members WHERE id = ?;

-- name: DeleteAccountMembersByAccountID :exec
-- DeleteAccountMembersByAccountID: deletes all members of an account
DELETE FROM account_members WHERE account_id = ?;

-- name: DeleteAccountMembersByUserID :exec
-- DeleteAccountMembersByUserID: deletes all members across all accounts for a user
DELETE FROM account_members WHERE user_id = ?;

-- name: CountAccountMembers :one
-- CountAccountMembers: retrieves members number of an account
SELECT COUNT(user_id) as count
FROM account_members
WHERE account_id = ?;

-- name: CountUserAccounts :one
-- CountUserAccounts: retrieves accounts number of a user
SELECT COUNT(account_id) as count
FROM account_members
WHERE user_id = ?;
