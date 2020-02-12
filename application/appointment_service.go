package app

import "psycare/domain"

type AppointmentStore interface {
	CreateAppointment(appt *domain.Appointment) error
	GetAppointments(id int64, forUser bool) (*[]domain.Appointment, error)
	AddRating(rating *domain.Rating) error
}

type AppointmentService struct {
	Store AppointmentStore
}

func (as *AppointmentService) CreateAppointment(appt *domain.Appointment) error {
	return as.Store.CreateAppointment(appt)
}

func (as *AppointmentService) GetAppointments(id int64, forUser bool) (*[]domain.Appointment, error) {
	return as.Store.GetAppointments(id, forUser)
}

func (as *AppointmentService) AddRating(rating *domain.Rating) error {
	return as.Store.AddRating(rating)
}
