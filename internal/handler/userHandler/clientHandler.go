package userHandler

import (
	"meawby/internal/handler/middleware"
	logger "meawby/internal/handler/middleware"
	"meawby/internal/model/modelUser"
	"meawby/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ClientHandler struct {
	client        service.ClientService
	master        service.MasterService
	masterService service.MasterServicesService
}

func NewClientHandler(service *service.Service) *ClientHandler {
	return &ClientHandler{client: service.ClientService, master: service.MasterService, masterService: service.MasterServicesService}
}

func (h *ClientHandler) CreateClientProfile(c *gin.Context) {
	var req modelUser.CreateClientRequest
	if err := c.BindJSON(&req); err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	accountID, err := middleware.GetAccountID(c)
	if err != nil {
		return
	}
	_, err = h.client.Create(c.Request.Context(), accountID, req.Name)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, "Cannot create client")
		return
	}
	c.JSON(201, "PROFILE CREATED")
}



func (h *ClientHandler) GetClientProfile(c *gin.Context) {
	accountID, err := middleware.GetAccountID(c)
	if err != nil {
		return
	}
	client, err := h.client.GetByAccountID(c.Request.Context(), accountID)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, modelUser.ClientResponse{ID: client.ID, Name: client.Name})
}

func (h *ClientHandler) GetAllActiveMasters(c *gin.Context) {
	masters, err := h.master.GetAll(c.Request.Context())
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, masters)
}

func (h *ClientHandler) GetMasterService(c *gin.Context) {
	masterID, _ := uuid.Parse("id")
	_, service, err := h.masterService.GetListForClient(c.Request.Context(), masterID)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, service)
}
