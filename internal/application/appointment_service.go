package app

import "psycare/internal/domain"

type AppointmentStore interface {
	CreateAppointment(appt *domain.Appointment) error
}

type AppointmentService struct {
	Store AppointmentStore
}

func (as *AppointmentService) CreateAppointment(appt *domain.Appointment) (err error) {
	err = as.Store.CreateAppointment(appt)
	return
}
