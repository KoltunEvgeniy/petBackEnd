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

type MasterRepo struct {
	db *sqlx.DB
}

func NewMasterRepo(db *sqlx.DB) *MasterRepo {
	return &MasterRepo{db: db}
}

func (r *MasterRepo) Create(ctx context.Context, master *modelUser.Master) error {
	query := "INSERT INTO masters(id,account_id,name) VALUES($1,$2,$3)"
	_, err := r.db.ExecContext(ctx, query, master.ID, master.AccountID, master.Name)
	return err
}

func (r *MasterRepo) GetByAccountId(ctx context.Context, accountID uuid.UUID) (*modelUser.Master, error) {
	var master modelUser.Master
	err := r.db.GetContext(ctx, &master, "SELECT id,account_id,name,is_active,created_at FROM masters WHERE account_id = $1", accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, modelErrors.ErrMaster
		}
		return nil, err
	}
	return &master, nil
}

func (r *MasterRepo) GetAllMasters(ctx context.Context) ([]modelUser.Master, error) {
	var m []modelUser.Master
	query := "SELECT id,account_id,name,is_active,created_at FROM masters where is_active = true"
	err := r.db.SelectContext(ctx, &m, query)
	return m, err
}
