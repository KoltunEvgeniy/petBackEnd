package scheduleHandler

import (
	"meawby/internal/handler/helper"
	logger "meawby/internal/handler/middleware"
	"meawby/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SlotHandl struct {
	slots service.SlotServive
	helper.AccountResolver
}

func NewSlotHandl(service *service.Service) *SlotHandl {
	return &SlotHandl{slots: service.SlotServive}
}

func (h *SlotHandl) GetMasterSlotsM(c *gin.Context) {
	masterID, err := h.GetMasterByID(c)
	if err != nil {
		return
	}
	date, err := helper.DateStrToTime(c)
	if err != nil {
		return
	}
	res, err := h.slots.GetForMaster(c.Request.Context(), masterID, *date)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *SlotHandl) GetMasterSlotsC(c *gin.Context) {
	masterId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		logger.NewErrorResponce(c, http.StatusNotFound, err.Error())
		return
	}
	date, err := helper.DateStrToTime(c)
	if err != nil {
		return
	}
	res, err := h.slots.GetFreeForClient(c.Request.Context(), masterId, *date)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
