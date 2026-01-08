package auth

import (
	"meawby/internal/model/modelAuth"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func generateJWT(userID uuid.UUID, role string, secret string, ttl time.Duration) (string, int64) {
	expiresAt := time.Now().Add(ttl).Unix()
	claims := modelAuth.JWTclaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString([]byte(secret))
	return signed, expiresAt
}

func generateRefresh(secret string, ttl time.Duration) (string, error) {
	expiresAt := time.Now().Add(ttl).Unix()
	id := uuid.New()
	claims := modelAuth.JWTclaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Id:        id.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
