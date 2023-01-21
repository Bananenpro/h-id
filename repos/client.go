package repos

import "context"

type ClientModel struct {
	BaseModel
	Name         string      `db:"name"`
	Description  string      `db:"description"`
	RedirectURIs StringSlice `db:"redirect_uris"`
	SecretHash   []byte      `db:"secret_hash"`
	UserID       string      `db:"user_id"`
}

type ClientRepository interface {
	Find(ctx context.Context, userID, id string) (*ClientModel, error)
	FindByUserID(ctx context.Context, userID string) ([]*ClientModel, error)
	Create(ctx context.Context, userID, name, description string, redirectURIs []string, secretHash []byte) (*ClientModel, error)
	Update(ctx context.Context, userID, id, name, description string, redirectURIs []string) error
	UpdateSecret(ctx context.Context, userID, id string, newSecretHash []byte) error
	Delete(ctx context.Context, userID, id string) error
}
