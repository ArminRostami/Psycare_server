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

	// check with advisor schedule:
	err := as.checkWithSchedule(appt)
	if err != nil {
		return errors.Wrap(err, "failed to create appointment")
	}

	// check with advisor appointments:
	err = as.checkWithAppointments(appt, appt.AdvisorID, false)
	if err != nil {
		return errors.Wrap(err, "failed to create appointment")
	}

	// check with user appointments:
	err = as.checkWithAppointments(appt, appt.UserID, true)
	if err != nil {
		return errors.Wrap(err, "failed to create appointment")
	}

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

func (as *AppointmentService) checkWithAppointments(appt *domain.Appointment, id int64, forUser bool) error {
	var target string
	if forUser {
		target = "user"
	} else {
		target = "advisor"
	}

	bookedAppts, err := as.AppointmentStore.GetAppointments(id, forUser)
	if err != nil {
		return errors.Wrap(err, "failed to check with other appointments")
	}

	for _, ref := range *bookedAppts {
		conflict := !(appt.EndTime.Before(ref.StartTime) || appt.StartTime.After(ref.EndTime))
		if conflict {
			return errors.Errorf("appointment is incompatible with other %s appointments", target)
		}
	}
	return nil
}

func (as *AppointmentService) checkWithSchedule(appt *domain.Appointment) error {
	sch, err := as.AdvisorStore.GetSchedule(appt.AdvisorID)
	if err != nil {
		return errors.Wrap(err, "failed to get schedule")
	}

	for _, period := range sch.Periods {
		if *period.DayOfWeek == int(appt.StartTime.Weekday()) &&
			after(appt.StartTime, period.StartTime) &&
			after(period.EndTime, appt.EndTime) {
			return nil
		}
	}
	return errors.New("appointment is incompatible with schedule")
}

func after(a, b time.Time) bool {
	if a.Hour() != b.Hour() {
		return a.Hour() > b.Hour()
	}
	if a.Minute() != b.Minute() {
		return a.Minute() > b.Minute()
	}
	if a.Second() != b.Second() {
		return a.Second() > b.Second()
	}
	return false
}
