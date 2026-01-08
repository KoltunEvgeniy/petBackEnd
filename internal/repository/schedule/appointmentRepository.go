package schedule

import (
	"context"
	"errors"
	"meawby/internal/model/modelSchedule"

	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AppointmentRepo struct {
	db *sqlx.DB
}

func NewAppointmentRepo(db *sqlx.DB) *AppointmentRepo {
	return &AppointmentRepo{db: db}
}

func (r *AppointmentRepo) Create(ctx context.Context, appt *modelSchedule.Appointment) error { //insert
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var exist bool
	query1 := "SELECT EXISTS(SELECT 1 FROM appointments WHERE master_id=$1 AND start_at < $3 and end_at >$2 FOR UPDATE)"
	err = tx.QueryRowContext(ctx, query1, appt.MasterID, appt.StartAt, appt.EndAt).Scan(&exist)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("Slot_Alredy_Booked")
	}
	query := "INSERT INTO appointments(id,client_id,master_id,service_id,start_at,end_at,status,price) values($1,$2,$3,$4,$5,$6,$7,$8)"
	_, err = tx.ExecContext(ctx, query, appt.Id, appt.ClientID, appt.MasterID, appt.ServiceID, appt.StartAt, appt.EndAt, appt.Status, appt.Price)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *AppointmentRepo) GetByID(ctx context.Context, apptID uuid.UUID) (*modelSchedule.Appointment, error) {
	var appt modelSchedule.Appointment
	query := "SELECT id,client_id,master_id,service_id,start_at,end_at,status,price,created_at FROM appointments WHERE id = $1"
	err := r.db.GetContext(ctx, &appt, query, apptID)
	return &appt, err
}

func (r *AppointmentRepo) ListByClient(ctx context.Context, clientID uuid.UUID) ([]modelSchedule.Appointment, error) {
	var appts []modelSchedule.Appointment
	query := "SELECT id,client_id,master_id,service_id,start_at,end_at,status,price,created_at FROM appointments WHERE client_id = $1"
	err := r.db.SelectContext(ctx, &appts, query, clientID)
	return appts, err
}

func (r *AppointmentRepo) ListByMaster(ctx context.Context, masterID uuid.UUID) ([]modelSchedule.Appointment, error) {
	var appts []modelSchedule.Appointment
	query := "SELECT id,client_id,master_id,service_id,start_at,end_at,status,price,created_at FROM appointments WHERE master_id = $1"
	err := r.db.SelectContext(ctx, &appts, query, masterID)
	return appts, err
}

func (r *AppointmentRepo) Exist(ctx context.Context, masterID uuid.UUID, startAt time.Time) (bool, error) {
	var count int
	query := "SELECT COUNT(*) FROM appointments WHERE master_id=$1 AND start_at=$2"
	err := r.db.GetContext(ctx, &count, query, masterID, startAt)
	return count > 0, err
}
