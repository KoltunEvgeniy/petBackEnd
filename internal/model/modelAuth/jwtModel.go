package modelAuth

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWTclaims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.StandardClaims
}
