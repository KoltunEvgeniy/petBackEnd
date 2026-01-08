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

type AccountRepo struct {
	db *sqlx.DB
}

func NewAccountRepo(db *sqlx.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (r *AccountRepo) GetByPhone(ctx context.Context, phone string) (*modelUser.Account, error) {
	var acc modelUser.Account
	err := r.db.GetContext(ctx, &acc, "SELECT id,phone,role,created_at FROM accounts WHERE phone=$1", phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &acc, nil
}

func (r *AccountRepo) Create(ctx context.Context, acc *modelUser.Account) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO accounts(id,phone)VALUES($1,$2)", acc.ID, acc.Phone)
	return err
}

func (r *AccountRepo) GetById(ctx context.Context, accountID uuid.UUID) (*modelUser.Account, error) {
	var account modelUser.Account
	err := r.db.GetContext(ctx, &account, "SELECT id,phone,role,created_at FROM accounts WHERE id=$1", accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepo) GetAll(ctx context.Context) ([]modelUser.Account, error) {
	var accounts []modelUser.Account
	err := r.db.SelectContext(ctx, &accounts, "SELECT id,phone,role,created_at from accounts where role IN ('client','master')")
	return accounts, err
}

func (r *AccountRepo) DeleteAccountByID(ctx context.Context, accountID uuid.UUID) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM accounts WHERE id = $1", accountID)
	if err != nil {
		return err
	}
	rf, _ := res.RowsAffected()
	if rf == 0 {
		return modelErrors.ErrAccountNotFound
	}
	return nil
}

func (r *AccountRepo) UpdateRole(ctx context.Context, accountID uuid.UUID, role string) error {
	_, err := r.db.ExecContext(ctx, "UPDATE accounts SET role = $1 WHERE id = $2", role, accountID)
	return err
}
