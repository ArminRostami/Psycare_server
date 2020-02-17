package postgres

import (
	"database/sql"
	"fmt"
	"psycare/domain"

	"github.com/pkg/errors"
)

type AdvisorStore struct {
	DB *PDB
}

func (as *AdvisorStore) CreateAdvisor(advisor *domain.Advisor) error {
	err := as.DB.NamedExecute(`
	INSERT into advisors (id, first_name, last_name, description, hourly_fee) 
	VALUES (:id, :first_name, :last_name, :description, :hourly_fee)`, advisor)
	return errors.Wrap(err, "failed to create advisor")
}

func (as *AdvisorStore) GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error) {
	advisors := &[]domain.Advisor{}
	err := as.DB.Con.Select(advisors, `
	SELECT id, first_name, last_name, description 
	FROM advisors WHERE verified=$1 
	LIMIT $2 OFFSET $3`, verified, limit, offset)

	if err != nil && err != sql.ErrNoRows {
		return nil, errors.Wrap(err, "failed to receive advisors")
	}
	return advisors, nil
}

func (as *AdvisorStore) GetAdvisorWithID(id int64) (*domain.Advisor, error) {
	advisor := &domain.Advisor{}
	err := as.DB.Con.Get(advisor, `
	SELECT * FROM advisors WHERE id=$1`, id)
	return advisor, errors.Wrap(err, "failed to get advisor info")
}

func (as *AdvisorStore) AddSchedule(sch *domain.Schedule) error {
	var errs string
	for _, p := range sch.Periods {
		err := as.DB.Execute(`
		INSERT INTO schedules (advisor_id, day_of_week, start_time, end_time) 
		VALUES ($1, $2, $3, $4)`, sch.AdvisorID, p.DayOfWeek, p.StartTime, p.EndTime)
		if err != nil {
			fmt.Println(err)
			errs += fmt.Sprintf("failed to add %v: %v\n", p, err)
		}
	}
	if errs != "" {
		return errors.Errorf("failed to add schedule: %s", errs)
	}
	return nil
}

func (as *AdvisorStore) GetSchedule(id int64) (*domain.Schedule, error) {
	periods := &[]domain.Period{}
	err := as.DB.Con.Select(periods, `
	SELECT day_of_week, start_time, end_time FROM schedules WHERE advisor_id=$1 `, id)
	return &domain.Schedule{AdvisorID: id, Periods: *periods}, errors.Wrap(err, "failed to get schedule")
}

func (as *AdvisorStore) GetAvgRating(advisorID int64) (float64, error) {
	avg := new(float64)
	err := as.DB.Con.Get(avg, `
	SELECT AVG(score) 
	FROM (SELECT id FROM appointments WHERE advisor_id=$1) as aps 
	INNER JOIN ratings ON aps.id=ratings.appointment_id
	`, advisorID)
	return *avg, errors.Wrap(err, "failed to get average rating")
}

func (as *AdvisorStore) DeleteScheduleWithID(id int64) error {
	err := as.DB.Execute(`
	DELETE FROM schedules
	WHERE advisor_id=$1`, id)
	return errors.Wrap(err, "failed to delete schedule")
}
