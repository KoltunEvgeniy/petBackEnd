package authHandler

import (
	logger "meawby/internal/handler/middleware"
	"meawby/internal/model/modelAuth"
	"meawby/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandl struct {
	auth service.AuthorizationService
}

func NewAuthHandler(service *service.Service) *AuthHandl {
	return &AuthHandl{auth: service.AuthorizationService}
}

type SendSmsCodeRequest struct {
	Phone string `json:"phone" binding:"required"`
}

func (h *AuthHandl) SendSMS(c *gin.Context) {
	var req SendSmsCodeRequest
	if err := c.BindJSON(&req); err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	if req.Phone == "" {
		logger.NewErrorResponce(c, http.StatusBadRequest, "phone required")
	}
	if err := h.auth.SendSMS(c.Request.Context(), req.Phone); err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (h *AuthHandl) Verify(c *gin.Context) {
	var req modelAuth.VerifySMSRequest
	if err := c.BindJSON(&req); err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.auth.VerifySMS(c.Request.Context(), req.Phone, req.Code)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AuthHandl) Refresh(c *gin.Context) {
	var req modelAuth.RefreshRequest
	if err := c.BindJSON(&req); err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.auth.Refresh(c.Request.Context(), req.RefreshToken)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
