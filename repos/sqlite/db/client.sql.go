// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: client.sql

package db

import (
	"context"
	"database/sql"
)

const createClient = `-- name: CreateClient :one
INSERT INTO clients (
  id, created_at, name, description, website, redirect_uris, secret_hash, user_id
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?  
) RETURNING id, created_at, name, description, website, redirect_uris, secret_hash, user_id
`

type CreateClientParams struct {
	ID           string
	CreatedAt    int64
	Name         string
	Description  string
	Website      string
	RedirectUris []byte
	SecretHash   []byte
	UserID       string
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (Client, error) {
	row := q.db.QueryRowContext(ctx, createClient,
		arg.ID,
		arg.CreatedAt,
		arg.Name,
		arg.Description,
		arg.Website,
		arg.RedirectUris,
		arg.SecretHash,
		arg.UserID,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Website,
		&i.RedirectUris,
		&i.SecretHash,
		&i.UserID,
	)
	return i, err
}

const deleteClient = `-- name: DeleteClient :execresult
DELETE FROM clients WHERE user_id = ? AND id = ?
`

type DeleteClientParams struct {
	UserID string
	ID     string
}

func (q *Queries) DeleteClient(ctx context.Context, arg DeleteClientParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteClient, arg.UserID, arg.ID)
}

const findClient = `-- name: FindClient :one
SELECT id, created_at, name, description, website, redirect_uris, secret_hash, user_id FROM clients WHERE id = ?
`

func (q *Queries) FindClient(ctx context.Context, id string) (Client, error) {
	row := q.db.QueryRowContext(ctx, findClient, id)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Website,
		&i.RedirectUris,
		&i.SecretHash,
		&i.UserID,
	)
	return i, err
}

const findClientByUser = `-- name: FindClientByUser :many
SELECT id, created_at, name, description, website, redirect_uris, secret_hash, user_id FROM clients WHERE user_id = ?
`

func (q *Queries) FindClientByUser(ctx context.Context, userID string) ([]Client, error) {
	rows, err := q.db.QueryContext(ctx, findClientByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Client
	for rows.Next() {
		var i Client
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.Website,
			&i.RedirectUris,
			&i.SecretHash,
			&i.UserID,
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

const findClientByUserAndID = `-- name: FindClientByUserAndID :one
SELECT id, created_at, name, description, website, redirect_uris, secret_hash, user_id FROM clients WHERE user_id = ? AND id = ?
`

type FindClientByUserAndIDParams struct {
	UserID string
	ID     string
}

func (q *Queries) FindClientByUserAndID(ctx context.Context, arg FindClientByUserAndIDParams) (Client, error) {
	row := q.db.QueryRowContext(ctx, findClientByUserAndID, arg.UserID, arg.ID)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Website,
		&i.RedirectUris,
		&i.SecretHash,
		&i.UserID,
	)
	return i, err
}

const updateClient = `-- name: UpdateClient :one
UPDATE clients SET
  name = ?, description = ?, website = ?, redirect_uris = ?
WHERE user_id = ? AND id = ?
RETURNING id, created_at, name, description, website, redirect_uris, secret_hash, user_id
`

type UpdateClientParams struct {
	Name         string
	Description  string
	Website      string
	RedirectUris []byte
	UserID       string
	ID           string
}

func (q *Queries) UpdateClient(ctx context.Context, arg UpdateClientParams) (Client, error) {
	row := q.db.QueryRowContext(ctx, updateClient,
		arg.Name,
		arg.Description,
		arg.Website,
		arg.RedirectUris,
		arg.UserID,
		arg.ID,
	)
	var i Client
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Website,
		&i.RedirectUris,
		&i.SecretHash,
		&i.UserID,
	)
	return i, err
}

const updateClientSecret = `-- name: UpdateClientSecret :execresult
UPDATE clients SET secret_hash = ? WHERE user_id = ? AND id = ?
`

type UpdateClientSecretParams struct {
	SecretHash []byte
	UserID     string
	ID         string
}

func (q *Queries) UpdateClientSecret(ctx context.Context, arg UpdateClientSecretParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateClientSecret, arg.SecretHash, arg.UserID, arg.ID)
}
