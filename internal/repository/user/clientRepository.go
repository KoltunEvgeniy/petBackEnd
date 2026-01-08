package user

import (
	"context"
	"database/sql"
	"errors"
	"meawby/internal/model/modelErrors"
	"meawby/internal/model/modelUser"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ClientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) *ClientRepo {
	return &ClientRepo{db: db}
}

func (r *ClientRepo) Create(ctx context.Context, client *modelUser.Client) error {
	query := "INSERT INTO clients(id,account_id,name) VALUES($1,$2,$3)"
	_, err := r.db.ExecContext(ctx, query, client.ID, client.AccountID, client.Name)
	return err
}

func (r *ClientRepo) GetByAccountID(ctx context.Context, accountID uuid.UUID) (*modelUser.Client, error) {
	var c modelUser.Client
	if err := r.db.GetContext(ctx, &c, "SELECT id, account_id,name,created_at FROM clients WHERE account_id =$1", accountID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, modelErrors.ErrAccountNotFound
		}
		return nil, err
	}
	return &c, nil

}
