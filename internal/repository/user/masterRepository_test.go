package user

import (
	"context"
	"errors"
	"meawby/internal/model/modelErrors"
	"meawby/internal/model/modelUser"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestMasterRepo_Create_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterRepo(db)

	cli := modelUser.Master{
		ID:        uuid.New(),
		AccountID: uuid.New(),
		Name:      "Anton",
	}

	mock.ExpectExec("INSERT INTO masters").WithArgs(cli.ID, cli.AccountID, cli.Name).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(context.Background(), &cli)
	require.NoError(t, err)
}

func TestMasterRepo_Create_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterRepo(db)

	cli := modelUser.Master{
		ID:        uuid.New(),
		AccountID: uuid.New(),
		Name:      "Anton",
	}

	mock.ExpectExec("INSERT INTO masters").WillReturnError(errors.New("db_error"))

	err := repo.Create(context.Background(), &cli)
	require.Error(t, err)
}

func TestMasterRepo_GetByAccountID_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterRepo(db)

	accID := uuid.New()
	masterID := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "account_id", "name", "is_active", "created_at"}).AddRow(masterID, accID, "Serega", true, now)

	mock.ExpectQuery("SELECT id,account_id,name,is_active,created_at FROM masters").WithArgs(accID).WillReturnRows(rows)

	client, err := repo.GetByAccountId(context.Background(), accID)
	require.NoError(t, err)
	require.NotNil(t, client)
	require.Equal(t, masterID, client.ID)
	require.Equal(t, "Serega", client.Name)
}

func TestMasterRepo_GetByAccountID_NotFound(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterRepo(db)

	accID := uuid.New()

	mock.ExpectQuery("SELECT id,account_id,name,is_active,created_at FROM masters").WithArgs(accID).WillReturnError(modelErrors.ErrAccountNotFound)

	_, err := repo.GetByAccountId(context.Background(), accID)
	require.Error(t, err)

}

func TestMasterRepo_GetAll_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterRepo(db)

	rows := sqlmock.NewRows([]string{"id", "account_id", "name", "is_active", "created_at"}).AddRow(uuid.New(), uuid.New(), "Luisa", true, time.Now()).AddRow(uuid.New(), uuid.New(), "Luisa2", true, time.Now())

	mock.ExpectQuery("SELECT id,account_id,name,is_active,created_at").
		WillReturnRows(rows)

	accs, err := repo.GetAllMasters(context.Background())

	require.NoError(t, err)
	require.Len(t, accs, 2)
}
