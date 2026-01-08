package schedule

import (
	"context"
	"log"
	"meawby/internal/model/modelSchedule"
	"meawby/internal/repository"
	"time"

	"github.com/google/uuid"
)

type AppointmentServ struct {
	appointment repository.AppointmentRepository
	service     repository.ServiceRepository
}

func NewAppointmentServ(repo *repository.Repository) *AppointmentServ {
	return &AppointmentServ{appointment: repo.AppointmentRepository, service: repo.ServiceRepository}
}

func (s *AppointmentServ) Create(ctx context.Context, clientID uuid.UUID, req *modelSchedule.AppointmentRequest) (*modelSchedule.Appointment, error) {
	service, err := s.service.GetById(ctx, req.ServiceId)
	if err != nil {
		return nil, err
	}
	log.Printf("GetByID service called : %v", service)
	startAt, _ := parseDateTime(req.Date, req.StartTime)
	endAt := startAt.Add(SlotDuration)
	a := &modelSchedule.Appointment{
		Id:        uuid.New(),
		ClientID:  clientID,
		MasterID:  req.MasterID,
		ServiceID: req.ServiceId,
		StartAt:   startAt,
		EndAt:     endAt,
		Status:    "booked",
		Price:     service.Price,
	}
	if err := s.appointment.Create(ctx, a); err != nil {
		return nil, err
	}
	return a, nil
}

func (s *AppointmentServ) GetByID(ctx context.Context, apptID uuid.UUID) (*modelSchedule.Appointment, error) {
	return s.appointment.GetByID(ctx, apptID)
}

func (s *AppointmentServ) ListByClient(ctx context.Context, clientID uuid.UUID) ([]modelSchedule.Appointment, error) {
	return s.appointment.ListByClient(ctx, clientID)
}

func (s *AppointmentServ) ListByMaster(ctx context.Context, masterID uuid.UUID) ([]modelSchedule.Appointment, error) {
	return s.appointment.ListByMaster(ctx, masterID)
}

func parseDateTime(date string, timestr string) (time.Time, error) {

	layout := "2006-01-02 15:04"
	return time.ParseInLocation(layout, date+" "+timestr, time.UTC)
}
