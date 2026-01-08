package user

import (
	"context"
	"meawby/internal/model/modelErrors"
	"meawby/internal/model/modelUser"
	"testing"
	"time"

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

func TestAccountRepo_GetByPhone_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAccountRepo(db)

	phone := "+37513771377"
	id := uuid.New()
	createdAt := time.Now()

	rows := sqlmock.NewRows([]string{"id", "phone", "role", "created_at"}).
		AddRow(id, phone, "client", createdAt)

	mock.ExpectQuery("SELECT id,phone,role,created_at FROM accounts").WithArgs(phone).
		WillReturnRows(rows)

	acc, err := repo.GetByPhone(context.Background(), phone)

	require.NoError(t, err)
	require.NotNil(t, acc)
	require.Equal(t, phone, acc.Phone)
}

func TestAccountRepo_GetByPhone_NotFound(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAccountRepo(db)

	phone := "+37513771377"

	mock.ExpectQuery("SELECT id,phone,role,created_at FROM accounts").WithArgs(phone).
		WillReturnError(modelErrors.ErrAccountNotFound)

	acc, err := repo.GetByPhone(context.Background(), phone)

	require.Error(t, err)
	require.Nil(t, acc)
}

func TestAccountRepo_Create_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAccountRepo(db)

	acc := modelUser.Account{
		ID:    uuid.New(),
		Phone: "13771377",
	}

	mock.ExpectExec("INSERT INTO accounts").WithArgs(acc.ID, acc.Phone).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(context.Background(), &acc)

	require.NoError(t, err)
}

func TestAccountRepo_GetByID_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAccountRepo(db)

	phone := "+37513771377"
	id := uuid.New()
	createdAt := time.Now()

	rows := sqlmock.NewRows([]string{"id", "phone", "role", "created_at"}).
		AddRow(id, phone, "client", createdAt)

	mock.ExpectQuery("SELECT id,phone,role,created_at FROM accounts").WithArgs(id).
		WillReturnRows(rows)

	acc, err := repo.GetById(context.Background(), id)

	require.NoError(t, err)
	require.NotNil(t, acc)
	require.Equal(t, id, acc.ID)
}

func TestAccountRepo_GetAll_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAccountRepo(db)

	rows := sqlmock.NewRows([]string{"id", "phone", "role", "created_at"}).AddRow(uuid.New(), "1234", "client", time.Now()).AddRow(uuid.New(), "1377", "master", time.Now())

	mock.ExpectQuery("SELECT id,phone,role,created_at").
		WillReturnRows(rows)

	accs, err := repo.GetAll(context.Background())

	require.NoError(t, err)
	require.Len(t, accs, 2)
}

func TestAccountRepo_DeleteAccountByID_NotFound(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAccountRepo(db)

	id := uuid.New()
	mock.ExpectExec("DELETE FROM accounts").WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.DeleteAccountByID(context.Background(), id)

	require.ErrorIs(t, err, modelErrors.ErrAccountNotFound)
}

func TestAccountRepo_UpdateRole_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAccountRepo(db)

	id := uuid.New()
	role := "admin"

	mock.ExpectExec("UPDATE accounts SET role").WithArgs(role, id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.UpdateRole(context.Background(), id, role)

	require.NoError(t, err)
}
