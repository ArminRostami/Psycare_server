package app

import "psycare/domain"

type AdvisorStore interface {
	CreateAdvisor(advisor *domain.Advisor) error
	GetAdvisorWithID(id int64) (*domain.Advisor, error)
	GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error)
	AddSchedule(sch *domain.Schedule) error
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
		return err
	}
	return nil
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
