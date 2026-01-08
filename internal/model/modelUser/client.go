package modelUser

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `db:"id"`
	Phone     string    `db:"phone"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
}

type Client struct {
	ID        uuid.UUID `db:"id" json:"id"`
	AccountID uuid.UUID `db:"account_id"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type CreateClientRequest struct {
	Name string `json:"name" binding:"required"`
}

type ClientResponse struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}
