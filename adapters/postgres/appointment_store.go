package postgres

import (
	"fmt"
	"psycare/domain"

	"github.com/pkg/errors"
)

type AppointmentStore struct {
	DB *PDB
}

func (as *AppointmentStore) CreateAppointment(appt *domain.Appointment) error {
	err := as.DB.namedExec(`
	INSERT INTO appointments (user_id, advisor_id, start_datetime, end_datetime) 
	VALUES (:user_id, :advisor_id, :start_datetime, :end_datetime)`, appt)
	return errors.Wrap(err, "failed to create appointment")
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
	return appts, errors.Wrap(err, "failed to get appointments")
}

func (as *AppointmentStore) AddRating(rating *domain.Rating) error {
	err := as.DB.namedExec(`
	INSERT INTO ratings(user_id, appointment_id, score) 
	VALUES (:user_id, :appointment_id, :score)`, rating)
	return errors.Wrap(err, "failed to add rating")
}

func (as *AppointmentStore) GetAppointmentWithID(id int64) (*domain.Appointment, error) {
	appt := &domain.Appointment{}
	err := as.DB.Con.Get(appt, `
		SELECT * FROM appointments WHERE id=$1`, id)
	return appt, errors.Wrap(err, "failed to get appointment info")
}

func (as *AppointmentStore) CancelAppointment(id int64) error {
	err := as.DB.exec(`
	UPDATE appointments SET cancelled=true WHERE id=$1`, id)
	return errors.Wrap(err, "failed to cancel appointment")
}
