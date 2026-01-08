package service

import (
	"context"

	"meawby/internal/model/modelService"
	"meawby/internal/repository"

	"github.com/google/uuid"
)

type ServiceServ struct {
	service repository.ServiceRepository
}

func NewServiceService(repo *repository.Repository) *ServiceServ {
	return &ServiceServ{service: repo.ServiceRepository}
}

func (s *ServiceServ) Create(ctx context.Context, title string, duration int, price int) error {
	service := &modelService.Service{
		ID:          uuid.New(),
		Title:       title,
		DurationMin: duration,
		Price:       price,
		IsActive:    true,
	}
	return s.service.Create(ctx, service)
}

func (s *ServiceServ) GetAll(ctx context.Context) ([]modelService.Service, error) {
	return s.service.GetAll(ctx)
}
