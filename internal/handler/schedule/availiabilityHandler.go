package scheduleHandler

import (
	"meawby/internal/handler/helper"
	"meawby/internal/handler/middleware"
	logger "meawby/internal/handler/middleware"

	"meawby/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AvailiabilityHandl struct {
	helper.AccountResolver
	availiability service.AvailiabilityService
}

func NewAvailiabilityHandl(service *service.Service) *AvailiabilityHandl {
	return &AvailiabilityHandl{availiability: service.AvailiabilityService}
}

func (h *AvailiabilityHandl) GetAvailiabilityByMaster(c *gin.Context) {
	masterID, err := h.GetMasterByID(c)
	if err != nil {
		return
	}
	from, _ := time.Parse("2006-01-02", c.Query("from"))
	to, _ := time.Parse("2006-01-02", c.Query("to"))
	slots, err := h.availiability.GetRange(c.Request.Context(), masterID, from, to)
	if err != nil {
		logger.NewErrorResponce(c, 500, err.Error())
		return
	}
	c.JSON(200, slots)
}


func (h *AvailiabilityHandl) GetAvailiabilityByClient(c *gin.Context) {
	masterID, _ := uuid.Parse(c.Param("id"))

	from, _ := time.Parse("2006-01-02", c.Query("from"))
	to, _ := time.Parse("2006-01-02", c.Query("to"))
	slots, err := h.availiability.GetRange(c.Request.Context(), masterID, from, to)
	if err != nil {
		logger.NewErrorResponce(c, 500, err.Error())
		return
	}
	c.JSON(200, slots)
}

func (h *AvailiabilityHandl) SetAvailiability(c *gin.Context) {
	masterID, err := h.GetMasterByID(c)
	if err != nil {
		return
	}

	var req struct {
		Date      string `json:"date" binding:"required"`
		Available bool   `json:"ava"`
	}
	if err := c.BindJSON(&req); err != nil {
		middleware.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.availiability.SetDate(c.Request.Context(), masterID, req.Date, req.Available); err != nil {
		middleware.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
