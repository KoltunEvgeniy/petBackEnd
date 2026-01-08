package modelUser

import (
	"time"

	"github.com/google/uuid"
)

type Master struct {
	ID        uuid.UUID            `db:"id"`
	AccountID uuid.UUID            `db:"account_id"`
	Name      string               `db:"name"`
	IsActive  bool                 `db:"is_active"`
	CreatedAt time.Time            `db:"created_at"`
	Service   []MasterServiceBrief `db:"services"`
}

type CreateMasterRequest struct {
	Name string `db:"name" binding:"required"`
}

type MasterResponse struct {
	ID       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	IsActive bool      `db:"is_active"`
}

type MasterService struct {
	MasterID  uuid.UUID `db:"master_id"`
	ServiceID uuid.UUID `db:"service_id"`
}

type MasterServiceBrief struct {
	ID    uuid.UUID `db:"id"`
	Title string    `json:"title"`
	Price string    `json:"price"`
}
