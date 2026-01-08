package modelPayment

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID            uuid.UUID `db:"id" json:"id"`
	AppointmentID uuid.UUID `db:"appointment_id" json:"appointment_id"`
	Amount        int       `db:"amount" json:"amount"`
	Status        string    `db:"status" json:"status"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

type PaymentRequest struct {
	AppointmentID uuid.UUID `json:"appointment_id" binding:"required"`
	Amount        int       `json:"amount" binding:"required"`
}

type PaymentResponse struct {
	ID            uuid.UUID `json:"id"`
	AppointmentID uuid.UUID `json:"appointment_id"`
	Amount        int       `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}
