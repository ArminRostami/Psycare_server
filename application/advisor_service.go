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
	AdvStore  AdvisorStore
	RoleStore RoleStore
}

func (as *AdvisorService) CreateAdvisor(advisor *domain.Advisor) error {
	err := as.AdvStore.CreateAdvisor(advisor)
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
	return as.AdvStore.GetAdvisors(verified, limit, offset)
}

func (as *AdvisorService) AddSchedule(sch *domain.Schedule) error {
	return as.AdvStore.AddSchedule(sch)
}

func (as *AdvisorService) GetAvgRating(advisorID int64) (float64, error) {
	return as.AdvStore.GetAvgRating(advisorID)
}

func (as *AdvisorService) GetAdvisorWithID(id int64) (*domain.Advisor, error) {
	return as.AdvStore.GetAdvisorWithID(id)
}
