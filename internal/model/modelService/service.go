package modelService

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	DurationMin int       `db:"duration_min" json:"duration_min"`
	Price       int       `db:"price" json:"price"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type CreateServiceRequest struct {
	Title       string `json:"title" binding:"required"`
	DurationMin int    `json:"duration_min" bingind:"required"`
	Price       int    `json:"price" binding:"required"`
}

type ServiceResponce struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	DurationMin int       `json:"duration_min"`
	Price       int       `json:"price"`
	IsActive    bool      `json:"is_active"`
}

func (s *Service) Duretion() time.Duration {
	return time.Minute * time.Duration(s.DurationMin)
}

type ServiceBrief struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Price string    `json:"price"`
}
