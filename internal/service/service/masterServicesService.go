package service

import (
	"context"

	"meawby/internal/model/modelUser"
	"meawby/internal/repository"

	"github.com/google/uuid"
)

type MasterServicesServ struct {
	masterService repository.MasterServiceRepository
	master        repository.MasterRepository
}

func NewMasterServicesServ(repo *repository.Repository) *MasterServicesServ {
	return &MasterServicesServ{masterService: repo.MasterServiceRepository, master: repo.MasterRepository}
}

func (s *MasterServicesServ) Add(ctx context.Context, masterID, serviceID uuid.UUID) error {
	ms := &modelUser.MasterService{
		MasterID:  masterID,
		ServiceID: serviceID,
	}
	return s.masterService.Add(ctx, ms)
}

func (s *MasterServicesServ) Remove(ctx context.Context, masterID, serviceID uuid.UUID) error {
	return s.masterService.Remove(ctx, masterID, serviceID)
}

func (s *MasterServicesServ) GetProfle(ctx context.Context, accountID uuid.UUID) (*modelUser.Master, error) {
	master, service, err := s.GetList(ctx, accountID)
	if err != nil {
		return nil, err
	}
	master.Service = service
	return master, nil
}

func (s *MasterServicesServ) GetList(ctx context.Context, accountID uuid.UUID) (*modelUser.Master, []modelUser.MasterServiceBrief, error) {
	master, err := s.master.GetByAccountId(ctx, accountID)
	if err != nil {
		return nil, nil, err
	}
	serviceIDs, err := s.masterService.ListByMaster(ctx, master.ID)
	if err != nil {
		return nil, nil, err
	}
	service, err := s.masterService.GetServicesByIDs(ctx, serviceIDs)
	if err != nil {
		return nil, nil, err
	}
	return master, service, nil
}

func (s *MasterServicesServ) GetListForClient(ctx context.Context, masterID uuid.UUID) (*modelUser.Master, []modelUser.MasterServiceBrief, error) {
	serviceIDs, err := s.masterService.ListByClient(ctx, masterID)
	if err != nil {
		return nil, nil, err
	}
	service, err := s.masterService.GetServicesByIDs(ctx, serviceIDs)
	if err != nil {
		return nil, nil, err
	}
	return nil, service, nil
}
