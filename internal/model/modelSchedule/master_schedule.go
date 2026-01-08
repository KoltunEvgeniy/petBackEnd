package modelSchedule

import (
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID        uuid.UUID `db:"id" json:"id"`
	MasterID  uuid.UUID `db:"master_id" json:"-"`
	DayOfWeek int       `db:"day_of_week" json:"day_of_week"`
	StartTime time.Time `db:"start_time" json:"start_time"`
	EndTime   time.Time `db:"end_time" json:"end_time"`
}

type ScheduleReq struct {
	DayOfWeek int    `db:"day_of_week" json:"day_of_week"`
	StartTime string `db:"start_time" json:"start_time"`
	EndTime   string `db:"end_time" json:"end_time"`
}

type Availiability struct {
	ID          uuid.UUID `db:"id"`
	MasterID    uuid.UUID `db:"master_id"`
	Date        time.Time `db:"date"`
	IsAvailable bool      `db:"is_available"`
}

type AvailiabilityReq struct {
	Date        string `db:"date"`
	IsAvailable bool   `db:"is_available"`
}

type AvailiabilityResp struct {
	Date        string `db:"date"`
	IsAvailable bool   `db:"is_available"`
}

type MasterScheduleReq struct {
	DayOfWeek int    `db:"day_of_week" json:"day_of_week"`
	StartTime string `db:"start_time" json:"start_time"`
	EndTime   string `db:"end_time" json:"end_time"`
}

type MasterScheduleResp struct {
	DayOfWeek int    `db:"day_of_week" json:"day_of_week"`
	StartTime string `db:"start_time" json:"start_time"`
	EndTime   string `db:"end_time" json:"end_time"`
}

type MasterDayOff struct {
	ID       uuid.UUID `db:"id"`
	MasterID uuid.UUID `db:"master_id"`
	Date     time.Time `db:"date"`
	Reason   string    `db:"reason"`
}

type SlotwWithStatus struct {
	StartTime string `json:"start_time"`
	Status    string `json:"status"`
}

type MasterDaySlotsResp struct {
	Date  string            `json:"date"`
	Slots []SlotwWithStatus `json:"slots"`
}
