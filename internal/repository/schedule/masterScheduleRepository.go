package schedule

import (
	"context"
	"database/sql"
	"errors"
	"meawby/internal/model/modelSchedule"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MasterScheduleRepo struct {
	db *sqlx.DB
}

func NewMasterScheduleRepo(db *sqlx.DB) *MasterScheduleRepo {
	return &MasterScheduleRepo{db: db}
}

func (r *MasterScheduleRepo) ListByMaster(ctx context.Context, masterID uuid.UUID) ([]modelSchedule.Schedule, error) { //GetShedule
	var schedules []modelSchedule.Schedule
	query := "SELECT day_of_week,start_time,end_time from master_schedules where master_id = $1"
	err := r.db.SelectContext(ctx, &schedules, query, masterID)
	return schedules, err
}

func (r *MasterScheduleRepo) DeleteByMaster(ctx context.Context, masterID uuid.UUID) error {
	query := "DELETE FROM master_schedules WHERE master_id = $1"
	_, err := r.db.ExecContext(ctx, query, masterID)
	return err
}

func (r *MasterScheduleRepo) Insert(ctx context.Context, s modelSchedule.Schedule) error { //setSchedule
	query := "INSERT INTO master_schedules(id,master_id,day_of_week,start_time,end_time)VALUES($1,$2,$3,$4,$5)"
	_, err := r.db.ExecContext(ctx, query, s.ID, s.MasterID, s.DayOfWeek, s.StartTime, s.EndTime)
	return err
}

func (r *MasterScheduleRepo) Save(ctx context.Context, masterID uuid.UUID, s []modelSchedule.Schedule) error { //setSchedule
	tx := r.db.MustBeginTx(ctx, nil)
	query := "INSERT INTO master_schedules(id,master_id,day_of_week,start_time,end_time)VALUES($1,$2,$3,$4,$5)"
	for _, sc := range s {
		_, err := tx.ExecContext(ctx, query, sc.ID, sc.MasterID, sc.DayOfWeek, sc.StartTime, sc.EndTime)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()

}

func (r *MasterScheduleRepo) GetByWeekDay(ctx context.Context, masterID uuid.UUID, weekday int) (*modelSchedule.Schedule, error) {
	var sch modelSchedule.Schedule
	query := "SELECT id,master_id,day_of_week,start_time,end_time FROM master_schedules WHERE master_id = $1 and day_of_week=$2"
	err := r.db.GetContext(ctx, &sch, query, masterID, weekday)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &sch, nil
}
