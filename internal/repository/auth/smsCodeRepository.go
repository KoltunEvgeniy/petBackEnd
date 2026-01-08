package auth

import (
	"context"
	"meawby/internal/model/modelAuth"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SmsCodeRepo struct {
	db *sqlx.DB
}

func NewSmsCodeRepo(db *sqlx.DB) *SmsCodeRepo {
	return &SmsCodeRepo{db: db}
}
func (r *SmsCodeRepo) Create(ctx context.Context, code *modelAuth.SMSCode) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO sms_codes(id,account_id,code,expires_at)VALUES($1,$2,$3,$4)",
		code.ID, code.AccountID, code.Code, code.ExpiresAt)
	return err
}

func (r *SmsCodeRepo) GetValidCode(ctx context.Context, accountID uuid.UUID, code string) (*modelAuth.SMSCode, error) {
	var sms modelAuth.SMSCode
	err := r.db.GetContext(ctx, &sms, "SELECT id,account_id,code,expires_at FROM sms_codes WHERE account_id=$1 and code=$2 and expires_at>now()", accountID, code)
	if err != nil {
		return nil, err
	}
	return &sms, err
}

func (r *SmsCodeRepo) Delete(ctx context.Context, smsID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM sms_codes WHERE id = $1", smsID)
	return err
}
