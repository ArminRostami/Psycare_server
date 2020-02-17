package app

import (
	"psycare/domain"

	"github.com/pkg/errors"
)

type AdvisorStore interface {
	CreateAdvisor(advisor *domain.Advisor) error
	GetAdvisorWithID(id int64) (*domain.Advisor, error)
	GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error)
	AddSchedule(sch *domain.Schedule) error
	GetSchedule(id int64) (*domain.Schedule, error)
	DeleteScheduleWithID(id int64) error
	GetAvgRating(advisorID int64) (float64, error)
}

type AdvisorService struct {
	AdvisorStore
	RoleStore
}

func (as *AdvisorService) CreateAdvisor(advisor *domain.Advisor) error {
	err := as.AdvisorStore.CreateAdvisor(advisor)
	if err != nil {
		return err
	}
	err = as.RoleStore.AddRole(advisor.ID, ROLE_ADVISOR)
	if err != nil {
		return errors.Wrap(err, "advisor created")
	}
	return nil
}

func (as *AdvisorService) DeleteScheduleWithID(id int64) error {
	return as.AdvisorStore.DeleteScheduleWithID(id)
}

func (as *AdvisorService) GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error) {
	return as.AdvisorStore.GetAdvisors(verified, limit, offset)
}

func (as *AdvisorService) AddSchedule(sch *domain.Schedule) error {
	return as.AdvisorStore.AddSchedule(sch)
}

func (as *AdvisorService) GetAvgRating(advisorID int64) (float64, error) {
	return as.AdvisorStore.GetAvgRating(advisorID)
}

func (as *AdvisorService) GetAdvisorWithID(id int64) (*domain.Advisor, error) {
	return as.AdvisorStore.GetAdvisorWithID(id)
}

func (as *AdvisorService) GetSchedule(id int64) (*domain.Schedule, error) {
	return as.AdvisorStore.GetSchedule(id)
}
