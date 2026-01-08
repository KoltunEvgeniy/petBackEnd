package auth

import (
	"meawby/internal/model/modelAuth"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateRefreshToken(secret string, ttl time.Duration) (string, error) {
	claims := modelAuth.JWTclaims{
		UserID: uuid.New(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return refreshToken.SignedString([]byte(secret))
}
