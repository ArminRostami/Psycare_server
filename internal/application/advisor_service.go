package app

import "psycare/internal/domain"

type AdvisorStore interface {
	CreateAdvisor(advisor *domain.Advisor) error
	GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error)
}

type AdvisorService struct {
	Store AdvisorStore
}

func (as *AdvisorService) CreateAdvisor(advisor *domain.Advisor) (err error) {
	err = as.Store.CreateAdvisor(advisor)
	return
}

func (as *AdvisorService) GetAdvisors(verified bool, limit, offset int) (advisors *[]domain.Advisor, err error) {
	advisors, err = as.Store.GetAdvisors(verified, limit, offset)
	return
}
