package schedule

import (
	"context"

	"meawby/internal/model/modelSchedule"
	"meawby/internal/repository"
	"time"

	"github.com/google/uuid"
)

type MasterScheduleServ struct {
	schedule repository.MasterScheduleRepository
}

func NewMasterScheduleServ(repo *repository.Repository) *MasterScheduleServ {
	return &MasterScheduleServ{schedule: repo.MasterScheduleRepository}
}

func (s *MasterScheduleServ) Set(ctx context.Context, masterID uuid.UUID, req []modelSchedule.ScheduleReq) error {
	var schedules []modelSchedule.Schedule

	for _, r := range req {
		start, _ := time.Parse("15:04", r.StartTime)
		end, _ := time.Parse("15:04", r.EndTime)
		schedules = append(schedules, modelSchedule.Schedule{
			ID:        uuid.New(),
			MasterID:  masterID,
			DayOfWeek: r.DayOfWeek,
			StartTime: start,
			EndTime:   end,
		})
	}

	return s.schedule.Save(ctx, masterID, schedules)
}

func (s *MasterScheduleServ) Get(ctx context.Context, masterID uuid.UUID) ([]modelSchedule.MasterScheduleResp, error) {
	schedules, err := s.schedule.ListByMaster(ctx, masterID)
	if err != nil {
		return nil, err
	}
	res := make([]modelSchedule.MasterScheduleResp, 0, len(schedules))
	for _, sch := range schedules {
		res = append(res, modelSchedule.MasterScheduleResp{
			DayOfWeek: sch.DayOfWeek,
			StartTime: sch.StartTime.Format("15:04"),
			EndTime:   sch.EndTime.Format("15:04"),
		})
	}
	return res, nil
}
