package modelAuth

import (
	"time"

	"github.com/google/uuid"
)

type SMSCode struct {
	ID        uuid.UUID `db:"id"`
	AccountID uuid.UUID `db:"account_id"`
	Code      string    `db:"code"`
	ExpiresAt time.Time `db:"expires_at"`
}

type RefreshToken struct {
	ID        uuid.UUID `db:"id"`
	AccountID uuid.UUID `db:"account_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `db:"expires_at"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type VerifySMSRequest struct {
	Phone string `db:"phone" binding:"required"`
	Code  string `db:"code" binding:"required"`
}

type VerifySMSResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type RefreshTokenResponse struct {
	AccessToken      string `json:"access_token"`
	AccessExpiresAt  int64  `json:"access_expires_at"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresAt int64  `json:"refresh_expires_at"`
}
