package user

import (
	"context"
	"meawby/internal/model/modelUser"
	"meawby/internal/repository"

	"github.com/google/uuid"
)

// import (
// 	"meawby/internal/model"
// 	"meawby/internal/repository"
// 	"time"

// 	"github.com/google/uuid"
// )

type MasterServ struct {
	master  repository.MasterRepository
	account repository.AccountRepository
	service repository.MasterServiceRepository
}

func NewMasterServ(repo *repository.Repository) *MasterServ {
	return &MasterServ{
		master:  repo.MasterRepository,
		account: repo.AccountRepository,
		service: repo.MasterServiceRepository}
}

func (s *MasterServ) Create(ctx context.Context, accountID uuid.UUID, name string) error {
	_, err := s.account.GetById(ctx, accountID)
	if err != nil {
		return err
	}
	master := &modelUser.Master{
		ID:        uuid.New(),
		AccountID: accountID,
		Name:      name,
	}
	return s.master.Create(ctx, master)

}

func (s *MasterServ) GetByAccountID(ctx context.Context, accountID uuid.UUID) (*modelUser.Master, error) {
	return s.master.GetByAccountId(ctx, accountID)
}

func (s *MasterServ) GetAll(ctx context.Context) ([]modelUser.Master, error) {
	masters, _ := s.master.GetAllMasters(ctx)
	for i := range masters {
		servisesIDS, _ := s.service.ListByClient(ctx, masters[i].ID)
		servises, _ := s.service.GetServicesByIDs(ctx, servisesIDS)
		masters[i].Service = servises
	}
	return masters, nil
}
