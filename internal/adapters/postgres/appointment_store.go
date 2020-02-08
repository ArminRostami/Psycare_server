package postgres

import "psycare/internal/domain"

type AppointmentStore struct {
	DB *PDB
}

func (as *AppointmentStore) CreateAppointment(appt *domain.Appointment) error {
	return as.DB.namedExec(`INSERT INTO 
	appointments (user_id, advisor_id, start_datetime, end_datetime)
	VALUES (:user_id, :advisor_id, :start_datetime, :end_datetime)`, appt)
}
