package service

import (
	"context"
	"errors"
	"meawby/internal/model/modelService"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestServiceRepo_Create_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewServiceRepo(db)

	s := &modelService.Service{
		ID:          uuid.New(),
		Title:       "Lix",
		DurationMin: 120,
		Price:       50,
		IsActive:    true,
	}
	mock.ExpectExec("INSERT INTO services").WithArgs(s.ID, s.Title, s.DurationMin, s.Price, s.IsActive).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(context.Background(), s)

	require.NoError(t, err)
}

func TestServiceRepo_Create_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewServiceRepo(db)

	s := &modelService.Service{
		ID:          uuid.New(),
		Title:       "Lix",
		DurationMin: 120,
		Price:       50,
		IsActive:    true,
	}
	mock.ExpectExec("INSERT INTO services").WithArgs(s.ID, s.Title, s.DurationMin, s.Price, s.IsActive).
		WillReturnError(errors.New("Db_err"))

	err := repo.Create(context.Background(), s)

	require.Error(t, err)
}

func TestServiceRepo_GetAll_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewServiceRepo(db)

	id := uuid.New()
	createdAtL := time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "title", "duration_min", "price", "is_active", "created_at",
	}).AddRow(id, "Manic", 120, 50, true, createdAtL)

	mock.ExpectQuery("SELECT id,title,duration_min,price,is_active,created_at FROM services").WillReturnRows(rows)
	serv, err := repo.GetAll(context.Background())

	require.NoError(t, err)
	require.Len(t, serv, 1)
	require.Equal(t, id, serv[0].ID)
	require.Equal(t, "Manic", serv[0].Title)
	require.True(t, serv[0].IsActive)
}

func TestServiceRepo_GetAll_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewServiceRepo(db)

	mock.ExpectQuery("SELECT id,title,duration_min,price,is_active,created_at FROM services").WillReturnError(errors.New("db_err"))
	_, err := repo.GetAll(context.Background())

	require.Error(t, err)
}

func TestServiceRepo_GetByID_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewServiceRepo(db)

	id := uuid.New()
	createdAtL := time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "title", "duration_min", "price", "is_active", "created_at",
	}).AddRow(id, "Manic", 120, 50, true, createdAtL)

	mock.ExpectQuery("SELECT id,title,duration_min,price,is_active,created_at FROM services").WithArgs(id).WillReturnRows(rows)
	serv, err := repo.GetById(context.Background(), id)

	require.NoError(t, err)
	require.NotNil(t, serv)
	require.Equal(t, id, serv.ID)
}

func TestServiceRepo_GetByID_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewServiceRepo(db)

	id := uuid.New()
	mock.ExpectQuery("SELECT id,title,duration_min,price,is_active,created_at FROM services").WithArgs(id).WillReturnError(errors.New("dbError"))
	_, err := repo.GetById(context.Background(), id)

	require.Error(t, err)

}
