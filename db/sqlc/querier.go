// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int32) error
	DeleteCategory(ctx context.Context, id int32) error
	GetAccount(ctx context.Context, id int32) (Account, error)
	GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error)
	GetAccountsByUserIdAndType(ctx context.Context, arg GetAccountsByUserIdAndTypeParams) ([]GetAccountsByUserIdAndTypeRow, error)
	GetAccountsByUserIdAndTypeAndCategoryId(ctx context.Context, arg GetAccountsByUserIdAndTypeAndCategoryIdParams) ([]GetAccountsByUserIdAndTypeAndCategoryIdRow, error)
	GetAccountsByUserIdAndTypeAndCategoryIdAndTitle(ctx context.Context, arg GetAccountsByUserIdAndTypeAndCategoryIdAndTitleParams) ([]GetAccountsByUserIdAndTypeAndCategoryIdAndTitleRow, error)
	GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescription(ctx context.Context, arg GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescriptionParams) ([]GetAccountsByUserIdAndTypeAndCategoryIdAndTitleAndDescriptionRow, error)
	GetAccountsByUserIdAndTypeAndDate(ctx context.Context, arg GetAccountsByUserIdAndTypeAndDateParams) ([]GetAccountsByUserIdAndTypeAndDateRow, error)
	GetAccountsByUserIdAndTypeAndDescription(ctx context.Context, arg GetAccountsByUserIdAndTypeAndDescriptionParams) ([]GetAccountsByUserIdAndTypeAndDescriptionRow, error)
	GetAccountsByUserIdAndTypeAndTitle(ctx context.Context, arg GetAccountsByUserIdAndTypeAndTitleParams) ([]GetAccountsByUserIdAndTypeAndTitleRow, error)
	GetAccountsGraph(ctx context.Context, arg GetAccountsGraphParams) (int64, error)
	GetAccountsReports(ctx context.Context, arg GetAccountsReportsParams) (int64, error)
	GetCategories(ctx context.Context, arg GetCategoriesParams) ([]Category, error)
	GetCategory(ctx context.Context, id int32) (Category, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUserById(ctx context.Context, id int32) (User, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
}

var _ Querier = (*Queries)(nil)
