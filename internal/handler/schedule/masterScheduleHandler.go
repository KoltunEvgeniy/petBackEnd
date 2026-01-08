package scheduleHandler

import (
	"meawby/internal/handler/helper"
	logger "meawby/internal/handler/middleware"
	"meawby/internal/model/modelSchedule"
	"meawby/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

type MasterScheduleHandl struct {
	masterSchedule service.MasterScheduleServive
	helper.AccountResolver
}

func NewMasterScheduleHandl(service *service.Service) *MasterScheduleHandl {
	return &MasterScheduleHandl{masterSchedule: service.MasterScheduleServive}
}

func (h *MasterScheduleHandl) GetMasterSchedule(c *gin.Context) {
	masterID, err := h.GetMasterByID(c)
	if err != nil {
		return
	}
	schedules, err := h.masterSchedule.Get(c.Request.Context(), masterID)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, schedules)
}

func (h *MasterScheduleHandl) AddMasterSchedule(c *gin.Context) {
	masterID, err := h.GetMasterByID(c)
	if err != nil {
		return
	}
	var req []modelSchedule.ScheduleReq
	if err := c.BindJSON(&req); err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.masterSchedule.Set(c.Request.Context(), masterID, req); err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
