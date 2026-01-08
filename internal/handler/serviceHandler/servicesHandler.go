package serviceHandler

import (
	"fmt"
	logger "meawby/internal/handler/middleware"
	"meawby/internal/model/modelService"
	"meawby/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type ServiceHandl struct {
	services service.ServicesService
}

func NewServiceHandl(service *service.Service) *ServiceHandl {
	return &ServiceHandl{services: service.ServicesService}
}

func (h *ServiceHandl) CreateService(c *gin.Context) {
	var req modelService.CreateServiceRequest
	if err := c.BindJSON(&req); err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Create(c.Request.Context(), req.Title, req.DurationMin, req.Price)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("CREATED SERVICE %v", req.Title))
}

func (h *ServiceHandl) GetServices(c *gin.Context) {
	services, err := h.services.GetAll(c.Request.Context())
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, services)
}
