package adminHandler

import (
	"fmt"
	"meawby/internal/handler/middleware"
	logger "meawby/internal/handler/middleware"
	"meawby/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AdminHandl struct {
	admin service.AdminService
}

func NewAdminHandler(service *service.Service) *AdminHandl {
	return &AdminHandl{admin: service.AdminService}
}

func (h *AdminHandl) UpdateRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accountIDstr := c.Param("id")
		accountID, err := uuid.Parse(accountIDstr)
		if err != nil {
			logger.NewErrorResponce(c, http.StatusNotFound, err.Error())
			return
		}
		err = h.admin.UpdateRole(c.Request.Context(), accountID, role)
		if err != nil {
			logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("%v ROLE UPDATED", role))
	}
}


func (h *AdminHandl) UpdateRoleNow(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accountId, err := middleware.GetAccountID(c)
		if err != nil {
			return
		}
		err = h.admin.UpdateRole(c.Request.Context(), accountId, role)
		if err != nil {
			logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, fmt.Sprintf("%v ROLE UPDATED", role))
	}
}

func (h *AdminHandl) GetAll(c *gin.Context) {
	accounts, err := h.admin.GetAll(c)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, accounts)
}
