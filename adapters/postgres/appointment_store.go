package postgres

import (
	"fmt"
	"psycare/domain"
	"time"

	"github.com/pkg/errors"
)

type AppointmentStore struct {
	DB *PDB
}

func (as *AppointmentStore) CreateAppointment(appt *domain.Appointment) error {
	return as.DB.namedExec(`
	INSERT INTO appointments (user_id, advisor_id, start_datetime, end_datetime) 
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

func (as *AppointmentStore) AddRating(rating *domain.Rating) error {
	appt := &domain.Appointment{}
	err := as.DB.Con.Get(appt, `
	SELECT * FROM appointments 
	WHERE id=$1 AND user_id=$2 AND cancelled=false`, rating.AppointmentID, rating.UserID)
	if err != nil {
		return errors.Wrap(err, "cannot get appointment: it does not exist or does not belong to this user")
	}

	if time.Now().After(appt.EndTime) {
		return errors.New("cannot rate appoinment before it's over")
	}

	err = as.DB.namedExec(`
	INSERT INTO ratings(user_id, appointment_id, score) 
	VALUES (:user_id, :appointment_id, :score)`, rating)
	if err != nil {
		return errors.Wrap(err, "failed to add rating")
	}

	return nil
}

func (as *AppointmentStore) GetAppointmentWithID(id int64) (*domain.Appointment, error) {
	appt := &domain.Appointment{}
	err := as.DB.Con.Get(appt, `
		SELECT * FROM appointments WHERE id=$1`, id)
	return appt, errors.Wrap(err, "failed to get appointment info")
}

func (as *AppointmentStore) CancelAppointment(id int64) error {
	return as.DB.exec(`
	UPDATE appointments SET cancelled=true WHERE id=$1`, id)
}
