package app

import (
	"psycare/domain"
	"time"

	"github.com/pkg/errors"
)

const CANCELLATION_PENALTY = 0.3

type AppointmentStore interface {
	CreateAppointment(appt *domain.Appointment) error
	GetAppointments(id int64, forUser bool) (*[]domain.Appointment, error)
	GetAppointmentWithID(id int64) (*domain.Appointment, error)
	AddRating(rating *domain.Rating) error
	CancelAppointment(id int64) error
}

type AppointmentService struct {
	AppointmentStore
	AdvisorStore
	UserStore
}

func (as *AppointmentService) CreateAppointment(appt *domain.Appointment) error {
	if appt.UserID == appt.AdvisorID {
		return errors.New("failed to create appointment: you cannot book an appointment with yourself")
	}
	// TODO: assert that appointment can be booked

	cost, err := as.CalculateCost(appt)
	if err != nil {
		return errors.Wrap(err, "failed to create appointment")
	}
	err = as.UserStore.Pay(appt.UserID, appt.AdvisorID, cost)
	if err != nil {
		return errors.Wrap(err, "failed to create appointment")
	}

	return as.AppointmentStore.CreateAppointment(appt)
}

func (as *AppointmentService) AddRating(rating *domain.Rating) error {
	appt, err := as.GetAppointmentWithID(rating.AppointmentID)
	if err != nil {
		return errors.Wrap(err, "failed to add rating")
	}

	if appt.UserID != rating.UserID {
		return errors.New("failed to add rating: you can only rate your own appointments")
	}

	if appt.Cancelled {
		return errors.New("failed to add rating: you cannot rate cancelled appointments")
	}
	// FIXME: this should always be BEFORE:
	if time.Now().Before(appt.EndTime) {
		return errors.New("failed to add rating: cannot rate appointment before it's over")
	}

	return as.AppointmentStore.AddRating(rating)
}

func (as *AppointmentService) CalculateCost(appt *domain.Appointment) (int64, error) {
	adv, err := as.AdvisorStore.GetAdvisorWithID(appt.AdvisorID)
	if err != nil {
		return 0, errors.Wrap(err, "failed to calculate cost")
	}
	dur := appt.EndTime.Sub(appt.StartTime)
	return int64(dur.Hours() * float64(adv.HourlyFee)), nil
}

func (as *AppointmentService) CancelAppointment(uid, appointmentID int64) error {
	appt, err := as.AppointmentStore.GetAppointmentWithID(appointmentID)
	if err != nil {
		return errors.Wrap(err, "failed to cancel appointment")
	}

	if uid != appt.UserID {
		return errors.New("failed to cancel appointment: it does not belong to this user")
	}

	if appt.Cancelled {
		return errors.New("failed to cancel appointment: already cancelled")
	}

	if time.Now().After(appt.StartTime) {
		return errors.New("failed to cancel appointment: appointment already started or finished")
	}

	cost, err := as.CalculateCost(appt)
	if err != nil {
		return errors.Wrap(err, "failed to cancel appointment")
	}

	refundValue := int64(float64(cost) * (1 - CANCELLATION_PENALTY))
	err = as.UserStore.Pay(appt.AdvisorID, uid, refundValue)
	if err != nil {
		return errors.Wrap(err, "failed to cancel appointment: refund failed")
	}

	return as.AppointmentStore.CancelAppointment(appointmentID)
}

func (as *AppointmentService) GetAppointments(id int64, forUser bool) (*[]domain.Appointment, error) {
	return as.AppointmentStore.GetAppointments(id, forUser)
}
