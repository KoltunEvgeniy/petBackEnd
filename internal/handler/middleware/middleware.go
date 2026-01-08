package middleware

import (
	"meawby/internal/model/modelErrors"
	"meawby/internal/service"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	headerAuthorization = "Authorization"
	ctxUserID           = "user_id"
	ctxRole             = "role"
)

type Middleware struct {
	auth service.AuthorizationService
}

func NewMiddleware(service *service.Service) *Middleware {
	return &Middleware{auth: service.AuthorizationService}
}

func (h *Middleware) UserIndentity(c *gin.Context) {
	header := c.GetHeader(headerAuthorization)
	if header == "" {
		NewErrorResponce(c, http.StatusUnauthorized, "No_auth_header")
		return
	}
	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		NewErrorResponce(c, http.StatusUnauthorized, "invalid_auth_header")
		return
	}
	claims, err := h.auth.ParseJWT(parts[1])
	if err != nil {
		NewErrorResponce(c, http.StatusUnauthorized, "invalid_token")
		return
	}
	c.Set(ctxUserID, claims.UserID)
	c.Set(ctxRole, claims.Role)
	c.Next()
}

func GetAccountID(c *gin.Context) (uuid.UUID, error) {
	v, ok := c.Get(ctxUserID)
	if !ok {
		NewErrorResponce(c, http.StatusBadRequest, modelErrors.ErrAccountNotFound.Error())
		return uuid.Nil, modelErrors.ErrAccountNotFound
	}
	aID, ok := v.(uuid.UUID)
	if !ok {
		NewErrorResponce(c, http.StatusBadRequest, modelErrors.ErrAccountNotFound.Error())
		return uuid.Nil, modelErrors.ErrAccountNotFound
	}
	return aID, nil
}

func (h *Middleware) RequereRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exist := c.Get(ctxRole)
		if !exist || userRole != role {
			NewErrorResponce(c, http.StatusForbidden, "FORBIDDEN")
			return
		}
		c.Next()
	}
}
