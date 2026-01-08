package payment

import (
	"context"

	"meawby/internal/model/modelPayment"
	"meawby/internal/repository"

	"github.com/google/uuid"
)

type PaymentServ struct {
	repo repository.PaymentRepository
}

func NewPaymentServ(repo *repository.Repository) *PaymentServ {
	return &PaymentServ{repo: repo.PaymentRepository}
}

func (s *PaymentServ) Create(ctx context.Context, p *modelPayment.Payment) error {
	return s.repo.Create(ctx, p)
}

func (s *PaymentServ) GetByAppointment(ctx context.Context, appointmentID uuid.UUID) ([]modelPayment.Payment, error) {
	return s.repo.GetByAppointmentID(ctx, appointmentID)
}
