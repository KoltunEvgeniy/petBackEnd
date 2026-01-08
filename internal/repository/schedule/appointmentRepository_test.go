package schedule

import (
	"context"
	"database/sql"
	"errors"
	"meawby/internal/model/modelSchedule"
	"regexp"
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

func TestAppointmentRepo_Create_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAppointmentRepo(db)

	appt := &modelSchedule.Appointment{
		Id:        uuid.New(),
		ClientID:  uuid.New(),
		MasterID:  uuid.New(),
		ServiceID: uuid.New(),
		StartAt:   time.Now().Add(time.Hour),
		EndAt:     time.Now().Add(2 * time.Hour),
		Price:     10,
		CreatedAt: time.Now(),
	}
	mock.ExpectBegin()

	mock.ExpectQuery("SELECT EXISTS").
		WithArgs(appt.MasterID, appt.StartAt, appt.EndAt).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO appointments")).
		WithArgs(appt.Id, appt.ClientID, appt.MasterID, appt.ServiceID, appt.StartAt, appt.EndAt, appt.Status, appt.Price).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Create(context.Background(), appt)
	require.NoError(t, err)
}

func TestAppointmentRepo_Create_SlotBooked(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAppointmentRepo(db)

	appt := &modelSchedule.Appointment{

		MasterID: uuid.New(),

		StartAt: time.Now().Add(time.Hour),
		EndAt:   time.Now().Add(2 * time.Hour),
	}
	mock.ExpectBegin()

	mock.ExpectQuery("SELECT EXISTS").
		WithArgs(appt.MasterID, appt.StartAt, appt.EndAt).WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	mock.ExpectRollback()
	err := repo.Create(context.Background(), appt)
	require.Error(t, err)
	require.Equal(t, "Slot_Alredy_Booked", err.Error())
}

func TestAppointmentRepo_Create_ExistQuerryError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAppointmentRepo(db)

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT EXISTS").WillReturnError(errors.New("Error"))

	mock.ExpectRollback()
	err := repo.Create(context.Background(), &modelSchedule.Appointment{})
	require.Error(t, err)
}

func TestAppointmentRepo_Create_InsertError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAppointmentRepo(db)

	appt := &modelSchedule.Appointment{
		Id:       uuid.New(),
		MasterID: uuid.New(),
		StartAt:  time.Now(),
		EndAt:    time.Now().Add(time.Hour),
	}
	mock.ExpectBegin()

	mock.ExpectQuery("SELECT EXISTS").WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO appointments")).WillReturnError(errors.New("insert_failed"))
	mock.ExpectRollback()

	err := repo.Create(context.Background(), appt)
	require.Error(t, err)
}

func TestAppointmentRepo_GetByID_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAppointmentRepo(db)

	id := uuid.New()

	rows := sqlmock.NewRows([]string{
		"id", "client_id", "master_id", "service_id", "start_at", "end_at", "status", "price", "created_at",
	}).AddRow(id, uuid.New(), uuid.New(), uuid.New(), time.Now(), time.Now().Add(time.Hour), "booked", 30, time.Now())

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id,client_id,master_id,service_id,start_at,end_at,status,price,created_at FROM appointments")).WithArgs(id).WillReturnRows(rows)

	app, err := repo.GetByID(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, id, app.Id)
}

func TestSmsCodeRepo_GetByID_CantFind(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAppointmentRepo(db)

	id := uuid.New()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id,client_id,master_id,service_id,start_at,end_at,status,price,created_at FROM appointments")).WithArgs(id).WillReturnError(sql.ErrNoRows)

	_, err := repo.GetByID(context.Background(), id)
	require.Error(t, err)
	require.Equal(t, sql.ErrNoRows, err)
}

func TestSmsCodeRepo_Exist_True(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewAppointmentRepo(db)

	mock.ExpectQuery("SELECT COUNT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
	exist, err := repo.Exist(context.Background(), uuid.New(), time.Now())
	require.NoError(t, err)
	require.True(t, exist)
}
