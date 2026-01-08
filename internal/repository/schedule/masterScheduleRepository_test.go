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

func TestMasterScheduleRepo_ListByMaster_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)

	masterID := uuid.New()

	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 20, 0, 0, 0, time.UTC)

	rows := sqlmock.NewRows([]string{"day_of_week", "start_time", "end_time"}).
		AddRow(1, start, end)

	mock.ExpectQuery("SELECT day_of_week,start_time,end_time from master_schedules").
		WithArgs(masterID).WillReturnRows(rows)

	res, err := repo.ListByMaster(context.Background(), masterID)

	require.NoError(t, err)
	require.Len(t, res, 1)
	require.Equal(t, 1, res[0].DayOfWeek)
}

func TestMasterScheduleRepo_ListByMaster_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)

	masterID := uuid.New()

	mock.ExpectQuery("SELECT day_of_week,start_time,end_time from master_schedules").
		WithArgs(masterID).WillReturnError(errors.New("db_error"))

	res, err := repo.ListByMaster(context.Background(), masterID)

	require.Error(t, err)
	require.Nil(t, res)
}

func TestMasterScheduleRepo_DeleteByMaster_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)

	masterID := uuid.New()

	mock.ExpectExec("DELETE FROM master_schedules").
		WithArgs(masterID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteByMaster(context.Background(), masterID)

	require.NoError(t, err)

}

func TestMasterScheduleRepo_DeleteByMaster_Error(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)

	masterID := uuid.New()

	mock.ExpectExec("DELETE FROM master_schedules").
		WithArgs(masterID).WillReturnError(errors.New("delete_filed"))

	err := repo.DeleteByMaster(context.Background(), masterID)

	require.Error(t, err)

}

func TestMasterScheduleRepo_Insert_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)

	masterID := uuid.New()

	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 20, 0, 0, 0, time.UTC)

	s := modelSchedule.Schedule{
		ID:        uuid.New(),
		MasterID:  masterID,
		DayOfWeek: 1,
		StartTime: start,
		EndTime:   end,
	}

	mock.ExpectExec("INSERT INTO master_schedules").
		WithArgs(s.ID, s.MasterID, s.DayOfWeek, s.StartTime, s.EndTime).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Insert(context.Background(), s)

	require.NoError(t, err)
}

func TestMasterScheduleRepo_Insert_Error(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)

	masterID := uuid.New()

	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 20, 0, 0, 0, time.UTC)

	s := modelSchedule.Schedule{
		ID:        uuid.New(),
		MasterID:  masterID,
		DayOfWeek: 1,
		StartTime: start,
		EndTime:   end,
	}

	mock.ExpectExec("INSERT INTO master_schedules").
		WithArgs(s.ID, s.MasterID, s.DayOfWeek, s.StartTime, s.EndTime).
		WillReturnError(errors.New("insert_error"))

	err := repo.Insert(context.Background(), s)

	require.Error(t, err)
}

func TestMasterScheduleRepo_Save_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)

	mock.ExpectBegin()

	masterID := uuid.New()

	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 20, 0, 0, 0, time.UTC)

	s := []modelSchedule.Schedule{
		{
			ID:        uuid.New(),
			MasterID:  masterID,
			DayOfWeek: 1,
			StartTime: start,
			EndTime:   end,
		},
	}

	mock.ExpectExec("INSERT INTO master_schedules").
		WithArgs(
			s[0].ID,
			s[0].MasterID,
			s[0].DayOfWeek,
			s[0].StartTime,
			s[0].EndTime,
		).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err := repo.Save(context.Background(), masterID, s)

	require.NoError(t, err)
}

func TestMasterScheduleRepo_Save_Error_Rollback(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)

	mock.ExpectBegin()

	masterID := uuid.New()

	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 20, 0, 0, 0, time.UTC)

	s := modelSchedule.Schedule{

		ID:        uuid.New(),
		MasterID:  masterID,
		DayOfWeek: 1,
		StartTime: start,
		EndTime:   end,
	}

	mock.ExpectExec("INSERT INTO master_schedules").
		WithArgs(
			s.ID,
			s.MasterID,
			s.DayOfWeek,
			s.StartTime,
			s.EndTime,
		).WillReturnError(errors.New("insert_error"))

	mock.ExpectRollback()

	err := repo.Save(context.Background(), masterID, []modelSchedule.Schedule{s})

	require.Error(t, err)
}

func TestMasterScheduleRepo_GetByWeekDay_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)
	masterID := uuid.New()

	start := time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)
	end := time.Date(2025, 1, 1, 20, 0, 0, 0, time.UTC)

	rows := sqlmock.NewRows([]string{"id", "master_id", "day_of_week", "start_time", "end_time"}).
		AddRow(uuid.New(), masterID, 1, start, end)
	mock.ExpectQuery("SELECT id,master_id,day_of_week,start_time,end_time FROM master_schedules").
		WithArgs(masterID, 1).WillReturnRows(rows)

	res, err := repo.GetByWeekDay(context.Background(), masterID, 1)

	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, 1, res.DayOfWeek)
}

func TestMasterScheduleRepo_GetByWeekDay_NotFound(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)
	masterID := uuid.New()

	mock.ExpectQuery("SELECT id,master_id,day_of_week,start_time,end_time FROM master_schedules").
		WithArgs(masterID, 1).WillReturnError(sql.ErrNoRows)

	res, err := repo.GetByWeekDay(context.Background(), masterID, 1)

	require.NoError(t, err)
	require.Nil(t, res)
}

func TestMasterScheduleRepo_GetByWeekDay_Error(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewMasterScheduleRepo(db)
	masterID := uuid.New()

	mock.ExpectQuery("SELECT id,master_id,day_of_week,start_time,end_time FROM master_schedules").
		WithArgs(masterID, 1).WillReturnError(errors.New("db_error"))

	res, err := repo.GetByWeekDay(context.Background(), masterID, 1)

	require.Error(t, err)
	require.Nil(t, res)
}
