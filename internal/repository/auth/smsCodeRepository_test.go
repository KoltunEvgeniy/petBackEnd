package auth

import (
	"context"
	"database/sql"
	"errors"
	"meawby/internal/model/modelAuth"
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
	repo := NewSmsCodeRepo(db)

	code := &modelAuth.SMSCode{
		ID:        uuid.New(),
		AccountID: uuid.New(),
		Code:      "1377",
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO sms_codes(id,account_id,code,expires_at)VALUES($1,$2,$3,$4)")).
		WithArgs(code.ID, code.AccountID, code.Code, code.ExpiresAt).WillReturnResult(sqlmock.NewResult(1, 1))
	err := repo.Create(context.Background(), code)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())

}

func TestSmsCodeRepo_Create_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSmsCodeRepo(db)

	code := &modelAuth.SMSCode{
		ID:        uuid.New(),
		AccountID: uuid.New(),
		Code:      "1377",
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO sms_codes(id,account_id,code,expires_at)VALUES($1,$2,$3,$4)")).
		WillReturnError(errors.New("DB_error"))
	err := repo.Create(context.Background(), code)
	require.Error(t, err)
	require.EqualError(t, err, "DB_error")

}

func TestSmsCodeRepo_GetValidCode_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSmsCodeRepo(db)

	accountID := uuid.New()
	code := "1377"

	expected := modelAuth.SMSCode{
		ID:        uuid.New(),
		AccountID: accountID,
		Code:      code,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	rows := sqlmock.NewRows([]string{"id", "account_id", "code", "expires_at"}).
		AddRow(expected.ID, expected.AccountID, expected.Code, expected.ExpiresAt)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id,account_id,code,expires_at FROM sms_codes WHERE account_id=$1 and code=$2 and expires_at>now()")).
		WithArgs(accountID, code).WillReturnRows(rows)
	result, err := repo.GetValidCode(context.Background(), accountID, code)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, expected.ID, result.ID)
}

func TestSmsCodeRepo_GetValidCode_NotFound(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSmsCodeRepo(db)

	accountID := uuid.New()
	code := "1377"

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id,account_id,code,expires_at FROM sms_codes WHERE account_id=$1 and code=$2 and expires_at>now()")).
		WithArgs(accountID, code).WillReturnError(sql.ErrNoRows)
	result, err := repo.GetValidCode(context.Background(), accountID, code)
	require.Error(t, err)
	require.Nil(t, result)
	require.Equal(t, sql.ErrNoRows, err)
}

func TestSmsCodeRepo_GetValidCode_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSmsCodeRepo(db)

	accountID := uuid.New()
	code := "1377"

	dberr := errors.New("DB_error")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id,account_id,code,expires_at FROM sms_codes WHERE account_id=$1 and code=$2 and expires_at>now()")).
		WithArgs(accountID, code).WillReturnError(dberr)
	result, err := repo.GetValidCode(context.Background(), accountID, code)
	require.Error(t, err)
	require.Nil(t, result)
	require.Equal(t, dberr, err)
}

func TestSmsCodeRepo_Delete_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSmsCodeRepo(db)
	id := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM sms_codes WHERE id = $1")).WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 1))
	err := repo.Delete(context.Background(), id)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSmsCodeRepo_Delete_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewSmsCodeRepo(db)
	id := uuid.New()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM sms_codes WHERE id = $1")).
		WillReturnError(errors.New("Delete_filed"))
	err := repo.Delete(context.Background(), id)
	require.Error(t, err)
	require.EqualError(t, err, "Delete_filed")
}
