package userHandler

import (
	"meawby/internal/handler/helper"
	"meawby/internal/handler/middleware"
	logger "meawby/internal/handler/middleware"
	"meawby/internal/model/modelUser"
	"meawby/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MasterHandler struct {
	helper.AccountResolver
	master        service.MasterService
	masterService service.MasterServicesService
}

func NewMasterHandler(service *service.Service) *MasterHandler {
	return &MasterHandler{master: service.MasterService, masterService: service.MasterServicesService}
}

func (h *MasterHandler) CreateMasterProfile(c *gin.Context) {
	var req modelUser.CreateMasterRequest
	if err := c.BindJSON(&req); err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	accountID, err := middleware.GetAccountID(c)
	if err != nil {
		return
	}
	err = h.master.Create(c.Request.Context(), accountID, req.Name)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, "PROFILE CREATED")
}

func (h *MasterHandler) GetMasterProfile(c *gin.Context) {
	accountID, err := middleware.GetAccountID(c)
	master, err := h.masterService.GetProfle(c.Request.Context(), accountID) //??
	if err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, master)
}

func (h *MasterHandler) AddService(c *gin.Context) {
	var req struct {
		ServiceID string `json:"service_id"`
	}
	if err := c.BindJSON(&req); err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	masterID, err := h.GetMasterByID(c)
	if err != nil {
		return
	}
	serviceID, err := uuid.Parse(req.ServiceID)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusNotFound, err.Error())
		return
	}

	if err := h.masterService.Add(c.Request.Context(), masterID, serviceID); err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (h *MasterHandler) RemoveService(c *gin.Context) {
	serviceIDstr := c.Param("id")
	serviceID, err := uuid.Parse(serviceIDstr)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusNotFound, err.Error())
		return
	}
	masterID, err := h.GetMasterByID(c)
	if err != nil {
		return
	}
	if err := h.masterService.Remove(c.Request.Context(), masterID, serviceID); err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (h *MasterHandler) ListService(c *gin.Context) {
	accountID, err := middleware.GetAccountID(c)
	if err != nil {
		return
	}
	_, services, err := h.masterService.GetList(c.Request.Context(), accountID)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, services)
}
