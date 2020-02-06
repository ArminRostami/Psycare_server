package app

import "psycare/internal/domain"

// AdvisorStore _
type AdvisorStore interface {
	CreateAdvisor(advisor *domain.Advisor) error
	GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error)
}

// AdvisorService _
type AdvisorService struct {
	Store AdvisorStore
}

// CreateAdvisor _
func (as *AdvisorService) CreateAdvisor(advisor *domain.Advisor) (err error) {
	err = as.Store.CreateAdvisor(advisor)
	return
}

// GetAdvisors _
func (as *AdvisorService) GetAdvisors(verified bool, limit, offset int) (advisors *[]domain.Advisor, err error) {
	advisors, err = as.Store.GetAdvisors(verified, limit, offset)
	return
}
