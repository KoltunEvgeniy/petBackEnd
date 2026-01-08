package paymenthandler

import (
	logger "meawby/internal/handler/middleware"
	"meawby/internal/model/modelPayment"
	"meawby/internal/service"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentHandl struct {
	payment service.PaymentService
}

func NewPaymentHandl(service *service.Service) *PaymentHandl {
	return &PaymentHandl{payment: service.PaymentService}
}

func (h *PaymentHandl) CreatePayment(c *gin.Context) {
	appointmentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		logger.NewErrorResponce(c, http.StatusNotFound, err.Error())
		return
	}
	var req modelPayment.Payment
	if err := c.BindJSON(&req); err != nil {
		logger.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	req.ID = uuid.New()
	req.AppointmentID = appointmentID
	err = h.payment.Create(c.Request.Context(), &req)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *PaymentHandl) ListPayment(c *gin.Context) {
	appointmentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		logger.NewErrorResponce(c, http.StatusNotFound, err.Error())
		return
	}
	payments, err := h.payment.GetByAppointment(c.Request.Context(), appointmentID)
	if err != nil {
		logger.NewErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, payments)
}
