package modelSchedule

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	Id        uuid.UUID `db:"id" json:"id"`
	ClientID  uuid.UUID `db:"client_id" json:"client_id"`
	MasterID  uuid.UUID `db:"master_id" json:"master_id"`
	ServiceID uuid.UUID `db:"service_id" json:"service_id"`
	StartAt   time.Time `db:"start_at" json:"start_at"`
	EndAt     time.Time `db:"end_at" json:"end_at"`
	Status    string    `db:"status" json:"status"`
	Price     int       `db:"price" json:"price"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type AppointmentRequest struct {
	MasterID  uuid.UUID `json:"master_id" binding:"required"`
	ServiceId uuid.UUID `json:"service_id" binding:"required"`
	Date      string    `json:"date" binding:"required"`
	StartTime string    `json:"start_at" binding:"required"`
}

type AppointmentResponse struct {
	Id        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	MasterID  uuid.UUID `json:"master_id"`
	ServiceId uuid.UUID `json:"service_id"`
	StartAt   time.Time `json:"start_at"`
	EndAt     time.Time `json:"end_at"`
	Status    string    `json:"status"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
