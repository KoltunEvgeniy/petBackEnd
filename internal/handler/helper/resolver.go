package helper

import (
	"errors"
	"meawby/internal/handler/middleware"
	"meawby/internal/model/modelErrors"
	"meawby/internal/service"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AccountResolver struct {
	client service.ClientService
	master service.MasterService
}

func NewAccountResolver(service *service.Service) *AccountResolver {
	return &AccountResolver{client: service.ClientService, master: service.MasterService}
}

func (r *AccountResolver) GetMasterByID(c *gin.Context) (uuid.UUID, error) {
	accountID, err := middleware.GetAccountID(c)
	if err != nil {
		return uuid.Nil, err
	}
	master, err := r.master.GetByAccountID(c.Request.Context(), accountID)
	if err != nil {
		middleware.NewErrorResponce(c, http.StatusNotFound, err.Error())
		return uuid.Nil, modelErrors.ErrAccountNotFound
	}
	return master.ID, nil

}

func (r *AccountResolver) GetClientByID(c *gin.Context) (uuid.UUID, error) {
	accountID, err := middleware.GetAccountID(c)
	if err != nil {
		return uuid.Nil, err
	}
	client, err := r.client.GetByAccountID(c.Request.Context(), accountID)
	if err != nil {
		middleware.NewErrorResponce(c, http.StatusNotFound, err.Error())
		return uuid.Nil, modelErrors.ErrAccountNotFound
	}
	return client.ID, nil

}

func DateStrToTime(c *gin.Context) (*time.Time, error) {
	dateStr := c.Query("date")
	if dateStr == "" {
		middleware.NewErrorResponce(c, http.StatusBadRequest, "Invalid_date")
		return nil, errors.New("Invalid_date")
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		middleware.NewErrorResponce(c, http.StatusBadRequest, err.Error())
		return nil, err
	}
	return &date, err
}
