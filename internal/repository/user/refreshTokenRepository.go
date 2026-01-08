package user

import (
	"context"
	"database/sql"
	"errors"
	"meawby/internal/model/modelAuth"
	"meawby/internal/model/modelErrors"

	"github.com/jmoiron/sqlx"
)

type RefreshTokenRepo struct {
	db *sqlx.DB
}

func NewRefreshTokenRepo(db *sqlx.DB) *RefreshTokenRepo {
	return &RefreshTokenRepo{db: db}
}

func (r *RefreshTokenRepo) Create(ctx context.Context, token *modelAuth.RefreshToken) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO refresh_tokens(id,account_id,token,expires_at)VALUES($1,$2,$3,$4)",
		token.ID, token.AccountID, token.Token, token.ExpiresAt)
	return err
}

func (r *RefreshTokenRepo) GetByToken(ctx context.Context, token string) (*modelAuth.RefreshToken, error) {
	var rt modelAuth.RefreshToken
	err := r.db.GetContext(ctx, &rt, "SELECT id,account_id,token,expires_at FROM refresh_tokens WHERE token=$1", token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, modelErrors.ErrMaster //ErrToken
		}
		return nil, err
	}
	return &rt, nil
}
