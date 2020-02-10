package postgres

import (
	"fmt"
	"psycare/internal/domain"
)

type AppointmentStore struct {
	DB *PDB
}

func (as *AppointmentStore) CreateAppointment(appt *domain.Appointment) error {
	return as.DB.namedExec(`INSERT INTO 
	appointments (user_id, advisor_id, start_datetime, end_datetime)
	VALUES (:user_id, :advisor_id, :start_datetime, :end_datetime)`, appt)
}

func (as *AppointmentStore) GetAppointments(id int64, forUser bool) (*[]domain.Appointment, error) {
	var fieldName string
	if forUser {
		fieldName = "user_id"
	} else {
		fieldName = "advisor_id"
	}
	query := fmt.Sprintf(`SELECT * from appointments WHERE %s=$1`, fieldName)
	appts := &[]domain.Appointment{}
	err := as.DB.Con.Select(appts, query, id)
	if err != nil {
		return nil, err
	}
	return appts, nil
}
