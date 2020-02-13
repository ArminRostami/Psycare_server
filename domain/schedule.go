package domain

import "time"

type Schedule struct {
	AdvisorID int64    `json:"advisor_id"`
	Periods   []Period `json:"periods" validate:"required,dive"`
}

type Period struct {
	DayOfWeek *int      `db:"day_of_week" json:"day_of_week" validate:"required,min=0,max=6"`
	StartTime time.Time `db:"start_time" json:"start_time" validate:"required"`
	EndTime   time.Time `db:"end_time" json:"end_time" validate:"required,gtfield=StartTime"`
}
