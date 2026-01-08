package admin

import (
	"context"
	"meawby/internal/model/modelUser"
	"meawby/internal/repository"

	"github.com/google/uuid"
)

type AdminService struct {
	client  repository.ClientRepository
	master  repository.MasterRepository
	account repository.AccountRepository
}

func NewAdminService(repo *repository.Repository) *AdminService {
	return &AdminService{client: repo.ClientRepository, master: repo.MasterRepository, account: repo.AccountRepository}
}

func (s *AdminService) UpdateRole(ctx context.Context, accountID uuid.UUID, role string) error {
	return s.account.UpdateRole(ctx, accountID, role)
}

func (s *AdminService) GetAll(ctx context.Context) ([]modelUser.Account, error) {
	return s.account.GetAll(ctx)
}

func (s *AdminService) DeleteAccountByID(ctx context.Context, accountID uuid.UUID) error {
	return s.account.DeleteAccountByID(ctx, accountID)
}
