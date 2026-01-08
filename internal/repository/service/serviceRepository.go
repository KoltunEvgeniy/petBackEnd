package service

import (
	"context"
	"meawby/internal/model/modelService"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ServiceRepo struct {
	db *sqlx.DB
}

func NewServiceRepo(db *sqlx.DB) *ServiceRepo {
	return &ServiceRepo{db: db}
}

func (r *ServiceRepo) Create(ctx context.Context, s *modelService.Service) error {
	query := "INSERT INTO services(id,title,duration_min,price,is_active)values($1,$2,$3,$4,$5)"
	_, err := r.db.ExecContext(ctx, query, s.ID, s.Title, s.DurationMin, s.Price, s.IsActive)
	return err
}

func (r *ServiceRepo) GetAll(ctx context.Context) ([]modelService.Service, error) {
	var ss []modelService.Service
	query := "SELECT id,title,duration_min,price,is_active,created_at FROM services WHERE is_active = true"
	err := r.db.SelectContext(ctx, &ss, query)
	return ss, err
}

func (r *ServiceRepo) GetById(ctx context.Context, serviceID uuid.UUID) (*modelService.Service, error) {
	var s modelService.Service
	query := "SELECT id,title,duration_min,price,is_active,created_at FROM services WHERE is_active = true and id=$1"
	err := r.db.GetContext(ctx, &s, query, serviceID)
	return &s, err
}
