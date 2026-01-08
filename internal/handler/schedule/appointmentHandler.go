package scheduleHandler

import (
	"meawby/internal/handler/helper"
	"meawby/internal/handler/middleware"

	"meawby/internal/model/modelSchedule"
	"meawby/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AppointmentHandl struct {
	appointment service.AppointmentService
	helper.AccountResolver
}

func NewAppointmentHandler(service *service.Service) *AppointmentHandl {
	return &AppointmentHandl{appointment: service.AppointmentService}
}

func (h *AppointmentHandl) CreateAppointment(c *gin.Context) {
	var req modelSchedule.AppointmentRequest
	if err := c.BindJSON(&req); err != nil {
		middleware.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	clientID, err := h.GetClientByID(c)
	if err != nil {
		return
	}

	res, err := h.appointment.Create(c.Request.Context(), clientID, &req)
	if err != nil {
		middleware.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *AppointmentHandl) ListAppointmentClient(c *gin.Context) {
	clientID, err := h.GetClientByID(c)
	if err != nil {
		return
	}
	appointmets, err := h.appointment.ListByClient(c.Request.Context(), clientID)
	if err != nil {
		middleware.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, appointmets)
}

func (h *AppointmentHandl) ListAppointmentMaster(c *gin.Context) {
	masterID, err := h.GetMasterByID(c)
	if err != nil {
		return
	}
	appointmets, err := h.appointment.ListByMaster(c.Request.Context(), masterID)
	if err != nil {
		middleware.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, appointmets)
}

