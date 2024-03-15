// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package repository

import (
	"context"
)

type Querier interface {
	// CountAccountMembers: retrieves members number of an account
	CountAccountMembers(ctx context.Context, accountID string) (int64, error)
	// CountUserAccounts: retrieves accounts number of a user
	CountUserAccounts(ctx context.Context, userID string) (int64, error)
	// CreateAccount: creates an account for a user
	CreateAccount(ctx context.Context, arg CreateAccountParams) error
	// CreateAccountMember: creates a member for an account
	CreateAccountMember(ctx context.Context, arg CreateAccountMemberParams) error
	// CreateUser: Create a new user in the database
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	// CreateUserSocialProfile: Link a user to a social profile
	CreateUserSocialProfile(ctx context.Context, arg CreateUserSocialProfileParams) error
	// DeleteAccount: deletes an account
	DeleteAccount(ctx context.Context, id string) error
	// DeleteAccountMember: deletes a member of an account
	DeleteAccountMember(ctx context.Context, id string) error
	// DeleteAccountMembersByAccountID: deletes all members of an account
	DeleteAccountMembersByAccountID(ctx context.Context, accountID string) error
	// DeleteAccountMembersByUserID: deletes all members across all accounts for a user
	DeleteAccountMembersByUserID(ctx context.Context, userID string) error
	// DeleteUser: Delete a user from the database
	DeleteUser(ctx context.Context, id string) error
	// DeleteUserSocialProfileBySocialID: Delete a user's social profile by social id and social name
	DeleteUserSocialProfileBySocialID(ctx context.Context, arg DeleteUserSocialProfileBySocialIDParams) error
	// DeleteUserSocialProfilesByUserID: Delete a user's social profiles by user id
	DeleteUserSocialProfilesByUserID(ctx context.Context, userID string) error
	// GetAccount: retrieves an account by its id
	GetAccount(ctx context.Context, id string) (Account, error)
	// GetAccountBySlug: retrieves an account by its slug
	GetAccountBySlug(ctx context.Context, slug string) (Account, error)
	// GetAccountMember: retrieves a member of an account by its id
	GetAccountMemberByID(ctx context.Context, id string) (AccountMember, error)
	// GetAccountMemberByUserID: retrieves a member of an account by its user_id
	GetAccountMemberByUserID(ctx context.Context, arg GetAccountMemberByUserIDParams) (AccountMember, error)
	// GetAccountMembers: retrieves members list of an account with pagination
	GetAccountMembers(ctx context.Context, arg GetAccountMembersParams) ([]AccountMember, error)
	// GetAccountUsers: retrieves users list of an account with pagination
	GetAccountUsers(ctx context.Context, arg GetAccountUsersParams) ([]GetAccountUsersRow, error)
	// GetUserAccounts: retrieves accounts list of a user with pagination
	GetUserAccounts(ctx context.Context, arg GetUserAccountsParams) ([]GetUserAccountsRow, error)
	// GetUserByEmail: Get a user by email
	GetUserByEmail(ctx context.Context, email string) (User, error)
	// GetUserByID: Get a user by id
	GetUserByID(ctx context.Context, id string) (User, error)
	// GetUserSocialProfileBySocialID: Get a user's social profile by social id and social name
	GetUserSocialProfileBySocialID(ctx context.Context, arg GetUserSocialProfileBySocialIDParams) (UserExternalProfile, error)
	// GetUserSocialProfileByUserID: Get a user's social profile by user id and social name
	GetUserSocialProfileByUserID(ctx context.Context, arg GetUserSocialProfileByUserIDParams) (UserExternalProfile, error)
	// GetUserSocialProfilesByUserID: Get a user's social profiles by user id
	GetUserSocialProfilesByUserID(ctx context.Context, userID string) ([]UserExternalProfile, error)
	// GetUsers: Get user list with pagination
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	// GetUsersCount: Get total user count
	GetUsersCount(ctx context.Context) (int64, error)
	// UpdateAccount: updates an account
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) error
	// UpdateAccountMember: updates a member of an account
	UpdateAccountMember(ctx context.Context, arg UpdateAccountMemberParams) error
	// UpdateUserEmail: Update a user's email
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error
	// UpdateUserPassword: Update a user's password
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	// VerifyUser: Verify a user's email
	VerifyUser(ctx context.Context, arg VerifyUserParams) error
}

var _ Querier = (*Queries)(nil)
