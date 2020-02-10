package app

import "psycare/domain"

type AdvisorStore interface {
	CreateAdvisor(advisor *domain.Advisor) error
	GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error)
	AddSchedule(sch *domain.Schedule) error
}

type AdvisorService struct {
	Store AdvisorStore
}

func (as *AdvisorService) CreateAdvisor(advisor *domain.Advisor) error {
	return as.Store.CreateAdvisor(advisor)
}

func (as *AdvisorService) GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error) {
	return as.Store.GetAdvisors(verified, limit, offset)
}

func (as *AdvisorService) AddSchedule(sch *domain.Schedule) error {
	return as.Store.AddSchedule(sch)
}
