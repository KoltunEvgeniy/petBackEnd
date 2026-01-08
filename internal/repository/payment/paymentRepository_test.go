package payment

import (
	"context"
	"errors"
	"meawby/internal/model/modelPayment"
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

func TestSmsCodeRepo_Create_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewPaymentRepo(db)

	payment := &modelPayment.Payment{
		ID:            uuid.New(),
		AppointmentID: uuid.New(),
		Amount:        1377,
		Status:        "paid",
	}
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO payments(id,appointment_id,amount,status)VALUES($1,$2,$3,$4)")).
		WithArgs(payment.ID, payment.AppointmentID, payment.Amount, payment.Status).WillReturnResult(sqlmock.NewResult(1, 1))
	err := repo.Create(context.Background(), payment)
	require.NoError(t, err)
}

func TestSmsCodeRepo_Create_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewPaymentRepo(db)

	payment := &modelPayment.Payment{
		ID:            uuid.New(),
		AppointmentID: uuid.New(),
		Amount:        1377,
		Status:        "paid",
	}
	dberr := errors.New("InsertFailed")
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO payments(id,appointment_id,amount,status)VALUES($1,$2,$3,$4)")).
		WithArgs(payment.ID, payment.AppointmentID, payment.Amount, payment.Status).WillReturnError(dberr)
	err := repo.Create(context.Background(), payment)
	require.Error(t, err)
	require.Equal(t, dberr, err)
}

func TestSmsCodeRepo_GetByAppointmentID_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewPaymentRepo(db)
	appointmentID := uuid.New()

	rows := sqlmock.NewRows([]string{"id", "appointment_id", "amount", "status", "created_at"}).AddRow(uuid.New(), appointmentID, 1377, "paid", time.Now())
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id,appointment_id,amount,status,created_at FROM payments")).WithArgs(appointmentID).WillReturnRows(rows)
	payments, err := repo.GetByAppointmentID(context.Background(), appointmentID)
	require.NoError(t, err)
	require.Len(t, payments, 1)
	require.Equal(t, "paid", payments[0].Status)
}

func TestSmsCodeRepo_GetByAppointmentID_Empty(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewPaymentRepo(db)
	appointmentID := uuid.New()

	rows := sqlmock.NewRows([]string{"id", "appointment_id", "amount", "status", "created_at"})

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id,appointment_id,amount,status,created_at FROM payments")).WithArgs(appointmentID).WillReturnRows(rows)
	payments, err := repo.GetByAppointmentID(context.Background(), appointmentID)
	require.NoError(t, err)
	require.Empty(t, payments)
}

func TestSmsCodeRepo_GetByAppointmentID_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewPaymentRepo(db)
	appointmentID := uuid.New()
	dberr := errors.New("db_err")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id,appointment_id,amount,status,created_at FROM payments")).WithArgs(appointmentID).WillReturnError(dberr)
	payments, err := repo.GetByAppointmentID(context.Background(), appointmentID)
	require.Error(t, err)
	require.Nil(t, payments)
	require.Equal(t, dberr, err)
}
