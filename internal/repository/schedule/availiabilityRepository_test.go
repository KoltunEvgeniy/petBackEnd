package schedule

import (
	"context"
	"database/sql"
	"errors"
	"meawby/internal/model/modelSchedule"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAvailiabilityRepo_Upset_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAvailiabilityRepo(db)

	a := modelSchedule.Availiability{
		ID:          uuid.New(),
		MasterID:    uuid.New(),
		Date:        time.Now(),
		IsAvailable: true,
	}
	mock.ExpectExec("INSERT INTO availability").
		WithArgs(a.ID, a.MasterID, a.Date, a.IsAvailable).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := repo.Upset(context.Background(), a)
	require.NoError(t, err)
}

func TestAvailiabilityRepo_Upset_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAvailiabilityRepo(db)
	mock.ExpectExec("INSERT INTO availability").
		WillReturnError(errors.New("upset_filed"))
	err := repo.Upset(context.Background(), modelSchedule.Availiability{})
	require.Error(t, err)
}

func TestAvailiabilityRepo_GetRange_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAvailiabilityRepo(db)

	masterID := uuid.New()
	from := time.Now()
	to := from.AddDate(0, 0, 7)

	rows := sqlmock.NewRows([]string{"date", "is_available"}).
		AddRow(from, true).
		AddRow(from.AddDate(0, 0, 1), false)

	mock.ExpectQuery("SELECT date,is_available FROM availability").
		WithArgs(masterID, from, to).
		WillReturnRows(rows)
	res, err := repo.GetRange(context.Background(), masterID, from, to)
	require.NoError(t, err)
	require.Len(t, res, 2)
	require.True(t, res[0].IsAvailable)
}

func TestAvailiabilityRepo_GetRange_Empty(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAvailiabilityRepo(db)

	masterID := uuid.New()
	from := time.Now()
	to := from.AddDate(0, 0, 7)

	rows := sqlmock.NewRows([]string{"date", "is_available"})
	mock.ExpectQuery("SELECT date,is_available FROM availability").
		WillReturnRows(rows)
	res, err := repo.GetRange(context.Background(), masterID, from, to)
	require.NoError(t, err)
	require.Empty(t, res)
}

func TestAvailiabilityRepo_GetByDay_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAvailiabilityRepo(db)

	masterID := uuid.New()
	date := time.Now()

	rows := sqlmock.NewRows([]string{"id", "master_id", "date", "is_available"}).
		AddRow(uuid.New(), uuid.New(), date, true)

	mock.ExpectQuery("SELECT id,master_id,date,is_available FROM availability").
		WithArgs(masterID, date).
		WillReturnRows(rows)
	res, err := repo.GetByDay(context.Background(), masterID, date)
	require.NoError(t, err)
	require.True(t, res.IsAvailable)
}

func TestAvailiabilityRepo_GetByDay_NotFound(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAvailiabilityRepo(db)

	mock.ExpectQuery("SELECT id,master_id,date,is_available FROM availability").
		WillReturnError(sql.ErrNoRows)
	res, err := repo.GetByDay(context.Background(), uuid.New(), time.Now())
	require.Error(t, err)
	require.Nil(t, res)
}

func TestAvailiabilityRepo_Update_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAvailiabilityRepo(db)

	masterID := uuid.New()
	date := time.Now()
	isAvaible := true

	mock.ExpectExec("UPDATE availability SET is_available").
		WithArgs(masterID, date, isAvaible).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := repo.Update(context.Background(), masterID, date, isAvaible)
	require.NoError(t, err)
}

func TestAvailiabilityRepo_Update_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAvailiabilityRepo(db)

	mock.ExpectExec("UPDATE availability SET is_available").
		WillReturnError(errors.New("db_error"))
	err := repo.Update(context.Background(), uuid.New(), time.Now(), true)
	require.Error(t, err)
}
