-- name: CreateAccountMember :exec
-- CreateAccountMember: creates a member for an account
INSERT INTO account_members (id, account_id, user_id, name, role, avatar_url) VALUES (@id, @account_id, @user_id, @name, @role, @avatar_url);

-- name: GetAccountMemberByID :one
-- GetAccountMember: retrieves a member of an account by its id
SELECT * FROM account_members WHERE id = @id;

-- name: GetAccountMembers :many
-- GetAccountMembers: retrieves members list of an account with pagination
SELECT * FROM account_members
WHERE account_id = @account_id
ORDER BY created_at DESC
LIMIT @limit_val OFFSET @offset_val;

-- name: GetAccountMemberByUserID :one
-- GetAccountMemberByUserID: retrieves a member of an account by its user_id
SELECT * FROM account_members WHERE account_id = @account_id AND user_id = @user_id;

-- name: UpdateAccountMember :exec
-- UpdateAccountMember: updates a member of an account
UPDATE account_members SET name = @name, role = @role, avatar_url = @avatar_url WHERE id = @id;

-- name: DeleteAccountMember :exec
-- DeleteAccountMember: deletes a member of an account
DELETE FROM account_members WHERE id = @id;

-- name: DeleteAccountMembersByAccountID :exec
-- DeleteAccountMembersByAccountID: deletes all members of an account
DELETE FROM account_members WHERE account_id = @account_id;

-- name: DeleteAccountMembersByUserID :exec
-- DeleteAccountMembersByUserID: deletes all members across all accounts for a user
DELETE FROM account_members WHERE user_id = @user_id;
