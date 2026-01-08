package schedule

import (
	"context"

	"meawby/internal/model/modelSchedule"
	"meawby/internal/repository"
	"time"

	"github.com/google/uuid"
)

type AvailiabilityServ struct {
	schedules repository.MasterScheduleRepository
	// dayOffs       repository.MasterDayOfWeekRepository
	appountment   repository.AppointmentRepository
	service       repository.ServiceRepository
	availiability repository.AvailiabilityRepository
}

func NewAvailiabilityServ(repo *repository.Repository) *AvailiabilityServ {
	return &AvailiabilityServ{
		schedules: repo.MasterScheduleRepository,
		// dayOffs:       repo.MasterDayOfWeekRepository,
		appountment:   repo.AppointmentRepository,
		service:       repo.ServiceRepository,
		availiability: repo.AvailiabilityRepository,
	}
}

func (s *AvailiabilityServ) SetDate(ctx context.Context, masterID uuid.UUID, dateStr string, avible bool) error {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}

	_, err = s.availiability.GetByDay(ctx, masterID, date)
	if err != nil {
		a := modelSchedule.Availiability{
			ID:          uuid.New(),
			MasterID:    masterID,
			Date:        date,
			IsAvailable: avible,
		}
		return s.availiability.Upset(ctx, a)
	}
	return s.availiability.Update(ctx, masterID, date, avible)

}

func (s *AvailiabilityServ) GetRange(ctx context.Context, masterID uuid.UUID, from, to time.Time) ([]modelSchedule.AvailiabilityResp, error) {
	rows, err := s.availiability.GetRange(ctx, masterID, from, to)
	if err != nil {
		return nil, err
	}

	res := make([]modelSchedule.AvailiabilityResp, 0, len(rows))

	for _, r := range rows {
		res = append(res, modelSchedule.AvailiabilityResp{
			Date:        r.Date.Format("2006-01-02"),
			IsAvailable: r.IsAvailable,
		})
	}
	return res, nil
}
