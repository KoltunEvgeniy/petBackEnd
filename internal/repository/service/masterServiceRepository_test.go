package service

import (
	"context"
	"meawby/internal/model/modelUser"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func newMockDB(t *testing.T) (*sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	sqlxDB := sqlx.NewDb(db, "postgres")
	return sqlxDB, mock
}
func TestServiceRepo_Add_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterServicesRepo(db)

	ms := &modelUser.MasterService{

		MasterID:  uuid.New(),
		ServiceID: uuid.New(),
	}
	mock.ExpectExec("INSERT INTO master_services").WithArgs(ms.MasterID, ms.ServiceID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Add(context.Background(), ms)

	require.NoError(t, err)
}

func TestServiceRepo_Remove_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterServicesRepo(db)

	masterID := uuid.New()
	serviceID := uuid.New()
	mock.ExpectExec("DELETE FROM master_services").WithArgs(masterID, serviceID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Remove(context.Background(), masterID, serviceID)

	require.NoError(t, err)
}

func TestServiceRepo_ListByMaster_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterServicesRepo(db)

	masterID := uuid.New()
	serviceID := uuid.New()

	rows := sqlmock.NewRows([]string{"service_id"}).AddRow(serviceID)

	mock.ExpectQuery("SELECT service_id from master_services").WithArgs(masterID).WillReturnRows(rows)

	serv, err := repo.ListByMaster(context.Background(), masterID)

	require.NoError(t, err)
	require.Len(t, serv, 1)
	require.Equal(t, serviceID, serv[0])
}

func TestServiceRepo_GetServicesByIDs_Empty(t *testing.T) {
	db, _ := newMockDB(t)
	repo := NewMasterServicesRepo(db)

	serv, err := repo.GetServicesByIDs(context.Background(), []uuid.UUID{})

	require.NoError(t, err)
	require.Len(t, serv, 0)
}

func TestServiceRepo_GetServicesByIDs_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterServicesRepo(db)
	id := uuid.New()
	price := 50
	rows := sqlmock.NewRows([]string{"id", "title", "price"}).AddRow(id, "Manicure", price)

	mock.ExpectQuery("SELECT id,title,price FROM services").WillReturnRows(rows)
	serv, err := repo.GetServicesByIDs(context.Background(), []uuid.UUID{id})

	require.NoError(t, err)
	require.Len(t, serv, 1)
	require.Equal(t, "Manicure", serv[0].Title)
}

func TestServiceRepo_GetServicesByIDs_InError(t *testing.T) {
	db, _ := newMockDB(t)
	repo := NewMasterServicesRepo(db)

	_, err := repo.GetServicesByIDs(context.Background(), nil)

	require.NoError(t, err)

}
