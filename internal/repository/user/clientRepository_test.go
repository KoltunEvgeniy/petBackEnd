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

func TestClientRepo_Create_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewClientRepo(db)

	cli := modelUser.Client{
		ID:        uuid.New(),
		AccountID: uuid.New(),
		Name:      "Anton",
	}

	mock.ExpectExec("INSERT INTO clients").WithArgs(cli.ID, cli.AccountID, cli.Name).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(context.Background(), &cli)
	require.NoError(t, err)
}

func TestClientRepo_Create_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewClientRepo(db)

	cli := modelUser.Client{
		ID:        uuid.New(),
		AccountID: uuid.New(),
		Name:      "Anton",
	}

	mock.ExpectExec("INSERT INTO clients").WillReturnError(errors.New("db_error"))

	err := repo.Create(context.Background(), &cli)
	require.Error(t, err)
}

func TestClientRepo_GetByAccountID_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewClientRepo(db)

	accID := uuid.New()
	clientID := uuid.New()
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "account_id", "name", "created_at"}).AddRow(clientID, accID, "Serega", now)

	mock.ExpectQuery("SELECT id, account_id,name,created_at FROM clients").WithArgs(accID).WillReturnRows(rows)

	client, err := repo.GetByAccountID(context.Background(), accID)
	require.NoError(t, err)
	require.NotNil(t, client)
	require.Equal(t, clientID, client.ID)
	require.Equal(t, "Serega", client.Name)
}

func TestClientRepo_GetByAccountID_NotFound(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewClientRepo(db)

	accID := uuid.New()

	mock.ExpectQuery("SELECT id, account_id,name,created_at FROM clients").WithArgs(accID).WillReturnError(modelErrors.ErrAccountNotFound)

	_, err := repo.GetByAccountID(context.Background(), accID)
	require.Error(t, err)

}
