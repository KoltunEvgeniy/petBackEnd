package schedule

import (
	"context"
	"meawby/internal/model/modelSchedule"

	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AvailiabilityRepo struct {
	db *sqlx.DB
}

func NewAvailiabilityRepo(db *sqlx.DB) *AvailiabilityRepo {
	return &AvailiabilityRepo{db: db}
}

func (r *AvailiabilityRepo) Upset(ctx context.Context, a modelSchedule.Availiability) error {
	query := "INSERT INTO availability(id,master_id,date,is_available)VALUES($1,$2,$3,$4) ON CONFLICT(master_id,date) DO UPDATE SET is_available=EXCLUDED.is_available"
	_, err := r.db.ExecContext(ctx, query, a.ID, a.MasterID, a.Date, a.IsAvailable)
	return err
}

func (r *AvailiabilityRepo) GetRange(ctx context.Context, masterID uuid.UUID, from, to time.Time) ([]modelSchedule.Availiability, error) {
	var res []modelSchedule.Availiability
	query := "SELECT date,is_available FROM availability WHERE master_id=$1 and date BETWEEN $2 and $3 ORDER BY date"
	err := r.db.SelectContext(ctx, &res, query, masterID, from, to)
	return res, err
}

func (r *AvailiabilityRepo) GetByDay(ctx context.Context, masterID uuid.UUID, date time.Time) (*modelSchedule.Availiability, error) {
	var a modelSchedule.Availiability
	query := "SELECT id,master_id,date,is_available FROM availability WHERE master_id=$1 and date=$2"
	err := r.db.GetContext(ctx, &a, query, masterID, date)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (r *AvailiabilityRepo) Update(ctx context.Context, masterID uuid.UUID, date time.Time, avaible bool) error {
	query := "UPDATE availability SET is_available=$3 WHERE master_id = $1 and date=$2"
	_, err := r.db.ExecContext(ctx, query, masterID, date, avaible)
	return err
}
