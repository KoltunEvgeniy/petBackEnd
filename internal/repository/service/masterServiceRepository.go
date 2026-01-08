package service

import (
	"context"
	"meawby/internal/model/modelUser"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MasterServicesRepo struct {
	db *sqlx.DB
}

func NewMasterServicesRepo(db *sqlx.DB) *MasterServicesRepo {
	return &MasterServicesRepo{db: db}
}

func (r *MasterServicesRepo) Add(ctx context.Context, ms *modelUser.MasterService) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO master_services(master_id,service_id) values($1,$2)", ms.MasterID, ms.ServiceID)
	return err
}

func (r *MasterServicesRepo) Remove(ctx context.Context, masterID, serviceID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM master_services WHERE master_id = $1 and service_id = $2", masterID, serviceID)
	return err
}

func (r *MasterServicesRepo) ListByMaster(ctx context.Context, masterID uuid.UUID) ([]uuid.UUID, error) {
	var services []uuid.UUID
	err := r.db.SelectContext(ctx, &services, "SELECT service_id from master_services where master_id = $1", masterID)
	return services, err
}

func (r *MasterServicesRepo) ListByClient(ctx context.Context, masterID uuid.UUID) ([]uuid.UUID, error) {
	var services []uuid.UUID
	err := r.db.SelectContext(ctx, &services, "SELECT service_id from master_services where master_id = $1", masterID)
	return services, err
}

func (r *MasterServicesRepo) GetServicesByIDs(ctx context.Context, ids []uuid.UUID) ([]modelUser.MasterServiceBrief, error) {
	if len(ids) == 0 {
		return []modelUser.MasterServiceBrief{}, nil
	}
	query, args, err := sqlx.In("SELECT id,title,price FROM services WHERE ID In(?)", ids)
	if err != nil {
		return nil, err
	}
	query = r.db.Rebind(query)
	var services []modelUser.MasterServiceBrief
	err = r.db.SelectContext(ctx, &services, query, args...)
	if err != nil {
		return nil, err
	}
	return services, nil
}
