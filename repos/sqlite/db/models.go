// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
)

type Client struct {
	ID           string
	CreatedAt    int64
	Name         string
	Description  string
	Website      string
	RedirectUris []byte
	SecretHash   []byte
	UserID       string
}

type Oauth struct {
	CreatedAt   int64
	ClientID    string
	Category    string
	TokenHash   []byte
	RedirectUri string
	UserID      string
	Scopes      string
	Data        []byte
	Expires     int64
	Used        bool
}

type Passkey struct {
	ID         string
	CredID     []byte
	Name       string
	CreatedAt  int64
	UserID     string
	Credential []byte
}

type Permission struct {
	CreatedAt int64
	ClientID  string
	UserID    string
	Scopes    string
}

type RecoveryCode struct {
	CreatedAt int64
	UserID    string
	CodeHash  []byte
}

type Remember2fa struct {
	CreatedAt int64
	UserID    string
	CodeHash  []byte
	Expires   int64
}

type RsaKey struct {
	Name      string
	CreatedAt int64
	Private   []byte
	Public    []byte
}

type Session struct {
	Token   string
	Data    []byte
	Expires int64
}

type Token struct {
	CreatedAt int64
	Category  string
	TokenKey  string
	ValueHash []byte
	Expires   int64
}

type User struct {
	ID              string
	CreatedAt       int64
	Name            string
	Email           string
	EmailConfirmed  bool
	PasswordHash    []byte
	OtpActive       bool
	OtpUrl          string
	NewEmail        sql.NullString
	NewEmailToken   []byte
	NewEmailExpires sql.NullInt64
	Admin           bool
}
