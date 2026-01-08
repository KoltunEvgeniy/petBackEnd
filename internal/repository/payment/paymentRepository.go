package payment

import (
	"context"
	"meawby/internal/model/modelPayment"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type PaymentRepo struct {
	db *sqlx.DB
}

func NewPaymentRepo(db *sqlx.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) Create(ctx context.Context, p *modelPayment.Payment) error {
	query := "INSERT INTO payments(id,appointment_id,amount,status)VALUES($1,$2,$3,$4)"
	_, err := r.db.ExecContext(ctx, query, p.ID, p.AppointmentID, p.Amount, p.Status)
	return err
}

func (r *PaymentRepo) GetByAppointmentID(ctx context.Context, appointmentID uuid.UUID) ([]modelPayment.Payment, error) {
	var payments []modelPayment.Payment
	query := "SELECT id,appointment_id,amount,status,created_at FROM payments where appointment_id = $1"
	err := r.db.SelectContext(ctx, &payments, query, appointmentID)
	return payments, err
}
