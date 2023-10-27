// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
    user_id,
    category_id,
    title,
    type,
    description,
    value, 
    date
)VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING id, user_id, category_id, title, type, description, value, date, created_at
`

type CreateAccountParams struct {
	UserID      int32     `json:"user_id"`
	CategoryID  int32     `json:"category_id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Value       int32     `json:"value"`
	Date        time.Time `json:"date"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.UserID,
		arg.CategoryID,
		arg.Title,
		arg.Type,
		arg.Description,
		arg.Value,
		arg.Date,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, user_id, category_id, title, type, description, value, date, created_at FROM accounts WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}

const getAccounts = `-- name: GetAccounts :many
SELECT 
    a.id,
    a.user_id,
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.created_at,
    c.title AS category_title
FROM accounts a LEFT JOIN 
categories c ON c.id = a.category_id 
WHERE a.user_id = $1 AND a.type = $2 AND a.category_id = $3 
AND a.title LIKE $4 AND a.description LIKE $5 AND a.date = $6
`

type GetAccountsParams struct {
	UserID      int32     `json:"user_id"`
	Type        string    `json:"type"`
	CategoryID  int32     `json:"category_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type GetAccountsRow struct {
	ID            int32          `json:"id"`
	UserID        int32          `json:"user_id"`
	Title         string         `json:"title"`
	Type          string         `json:"type"`
	Description   string         `json:"description"`
	Value         int32          `json:"value"`
	Date          time.Time      `json:"date"`
	CreatedAt     time.Time      `json:"created_at"`
	CategoryTitle sql.NullString `json:"category_title"`
}

func (q *Queries) GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccounts,
		arg.UserID,
		arg.Type,
		arg.CategoryID,
		arg.Title,
		arg.Description,
		arg.Date,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsRow{}
	for rows.Next() {
		var i GetAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreatedAt,
			&i.CategoryTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountsGraph = `-- name: GetAccountsGraph :one
SELECT COUNT(*) FROM accounts
WHERE user_id = $1 AND type = $2
`

type GetAccountsGraphParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) GetAccountsGraph(ctx context.Context, arg GetAccountsGraphParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountsGraph, arg.UserID, arg.Type)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAccountsReports = `-- name: GetAccountsReports :one
SELECT SUM(value) AS sum_value FROM accounts
WHERE user_id = $1 AND type = $2
`

type GetAccountsReportsParams struct {
	UserID int32  `json:"user_id"`
	Type   string `json:"type"`
}

func (q *Queries) GetAccountsReports(ctx context.Context, arg GetAccountsReportsParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountsReports, arg.UserID, arg.Type)
	var sum_value int64
	err := row.Scan(&sum_value)
	return sum_value, err
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts SET title = $2, 
description = $3, value = $4  WHERE id = $1 
RETURNING id, user_id, category_id, title, type, description, value, date, created_at
`

type UpdateAccountParams struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Value,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoryID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreatedAt,
	)
	return i, err
}
