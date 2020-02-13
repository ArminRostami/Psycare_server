package postgres

import (
	"database/sql"
	"fmt"
	"psycare/domain"
	"time"

	"github.com/pkg/errors"
)

type AdvisorStore struct {
	DB *PDB
}

func (as *AdvisorStore) CreateAdvisor(advisor *domain.Advisor) error {
	return as.DB.namedExec(`
	INSERT into advisors (id, first_name, last_name, description) 
	VALUES (:id, :first_name, :last_name, :description)`, advisor)
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
	if err != nil {
		return nil, errors.Wrap(err, "failed to get advisor")
	}
	return advisor, nil
}

func (as *AdvisorStore) AddSchedule(sch *domain.Schedule) error {
	var errs string
	for _, p := range sch.Periods {
		err := as.DB.exec(`
		INSERT INTO schedules (advisor_id, day_of_week, start_time, end_time) 
		VALUES ($1, $2, $3, $4)`, sch.AdvisorID, p.DayOfWeek, getTime(p.StartTime), getTime(p.EndTime))
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

func getTime(src time.Time) string {
	h, m, s := src.Clock()
	zone, _ := src.Zone()
	return fmt.Sprintf("%d:%d:%d%s", h, m, s, zone)
}

func (as *AdvisorStore) GetAvgRating(advisorID int64) (float64, error) {
	avg := new(float64)
	err := as.DB.Con.Get(avg, `
	SELECT AVG(score) 
	FROM (SELECT id FROM appointments WHERE advisor_id=$1) as aps 
	INNER JOIN ratings ON aps.id=ratings.appointment_id
	`, advisorID)
	if err != nil {
		return 0, errors.Wrap(err, "could not get average rating")
	}
	return *avg, nil
}
