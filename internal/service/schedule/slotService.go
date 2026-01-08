package schedule

import (
	"context"
	"database/sql"
	"errors"
	"meawby/internal/model/modelSchedule"
	"meawby/internal/repository"
	"time"

	"github.com/google/uuid"
)

const SlotDuration = 2 * time.Hour

type SlotServ struct {
	schedule      repository.MasterScheduleRepository
	availiability repository.AvailiabilityRepository
	appointment   repository.AppointmentRepository
}

func NewSlotServ(repo *repository.Repository) *SlotServ {
	return &SlotServ{schedule: repo.MasterScheduleRepository, availiability: repo.AvailiabilityRepository, appointment: repo.AppointmentRepository}
}

func (s *SlotServ) GetForMaster(ctx context.Context, masterID uuid.UUID, date time.Time) (*modelSchedule.MasterDaySlotsResp, error) {
	avl, err := s.availiability.GetByDay(ctx, masterID, date)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	isAva := true
	if avl != nil {
		isAva = avl.IsAvailable
	}
	if !isAva {
		return &modelSchedule.MasterDaySlotsResp{
			Date:  date.Format("2006-01-02"),
			Slots: []modelSchedule.SlotwWithStatus{},
		}, nil
	}

	day := int(date.Weekday())
	if day == 0 {
		day = 7
	}

	schedule, err := s.schedule.GetByWeekDay(ctx, masterID, day)
	if err != nil {
		return nil, err
	}
	if schedule == nil {
		return &modelSchedule.MasterDaySlotsResp{
			Date:  date.Format("2006-01-02"),
			Slots: []modelSchedule.SlotwWithStatus{},
		}, nil
	}

	dayStart := time.Date(date.Year(), date.Month(), date.Day(),
		schedule.StartTime.Hour(), schedule.StartTime.Minute(), 0, 0, time.UTC)
	dayEnd := time.Date(date.Year(), date.Month(), date.Day(),
		schedule.EndTime.Hour(), schedule.EndTime.Minute(), 0, 0, time.UTC)

	slots := []modelSchedule.SlotwWithStatus{}
	t := dayStart
	for !t.Add(SlotDuration).After(dayEnd) {
		booked, _ := s.appointment.Exist(ctx, masterID, t)
		status := "free"
		if booked {
			status = "booked"
		}
		slots = append(slots, modelSchedule.SlotwWithStatus{
			StartTime: t.Format("15:04"),
			Status:    status,
		})
		t = t.Add(SlotDuration)
	}
	return &modelSchedule.MasterDaySlotsResp{
		Date:  date.Format("2006-01-02"),
		Slots: slots,
	}, nil
}

func (s *SlotServ) GetFreeForClient(ctx context.Context, masterID uuid.UUID, date time.Time) (*modelSchedule.MasterDaySlotsResp, error) {
	res, err := s.GetForMaster(ctx, masterID, date)
	if err != nil {
		return nil, err
	}
	free := make([]modelSchedule.SlotwWithStatus, 0)

	for _, slot := range res.Slots {
		if slot.Status == "free" {
			free = append(free, slot)
		}
	}
	res.Slots = free
	return res, nil
}
