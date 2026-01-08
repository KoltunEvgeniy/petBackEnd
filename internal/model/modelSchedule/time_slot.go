package modelSchedule

import "time"

type TimeSlot struct {
	StartAt time.Time `json:"start_at" db:"start_at"`
	EndAt   time.Time `json:"end_at" db:"end_at"`
}
