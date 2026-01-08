package user

import (
	"context"
	"errors"
	"meawby/internal/model/modelAuth"
	"meawby/internal/model/modelErrors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestRefreshTokenRepo_Create_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewRefreshTokenRepo(db)

	token := &modelAuth.RefreshToken{
		ID:        uuid.New(),
		AccountID: uuid.New(),
		Token:     "qefimunefwnufe",
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	mock.ExpectExec("INSERT INTO refresh_tokens").WithArgs(token.ID, token.AccountID, token.Token, token.ExpiresAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(context.Background(), token)

	require.NoError(t, err)
}

func TestRefreshTokenRepo_Create_DBError(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewRefreshTokenRepo(db)

	token := &modelAuth.RefreshToken{
		ID:        uuid.New(),
		AccountID: uuid.New(),
		Token:     "qefimunefwnufe",
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	mock.ExpectExec("INSERT INTO refresh_tokens").WithArgs(token.ID, token.AccountID, token.Token, token.ExpiresAt).
		WillReturnError(errors.New("db_error"))

	err := repo.Create(context.Background(), token)

	require.Error(t, err)
}

func TestRefreshTokenRepo_GetNyToken_Success(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewRefreshTokenRepo(db)

	token := "eopkwecncw"

	rows := sqlmock.NewRows([]string{"id", "account_id", "token", "expires_at"}).AddRow(uuid.New(), uuid.New(), token, time.Now().Add(5*time.Minute))
	mock.ExpectQuery("SELECT id,account_id,token,expires_at FROM refresh_tokens").WithArgs(token).WillReturnRows(rows)

	tokenInf, err := repo.GetByToken(context.Background(), token)

	require.NoError(t, err)
	require.Equal(t, token, tokenInf.Token)
}

func TestRefreshTokenRepo_GetNyToken_NotFound(t *testing.T) {
	db, mock := newMockDB(t)
	repo := NewRefreshTokenRepo(db)

	token := "eopkwecncw"

	mock.ExpectQuery("SELECT id,account_id,token,expires_at FROM refresh_tokens").WithArgs(token).WillReturnError(modelErrors.ErrAccountNotFound)

	_, err := repo.GetByToken(context.Background(), token)

	require.Error(t, err)

}
