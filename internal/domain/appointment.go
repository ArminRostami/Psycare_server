package domain

import "time"

type Appointment struct {
	UasUrl    int64
	AdvisorID int64
	StartTime time.Time
	EndTime   time.Time
}
