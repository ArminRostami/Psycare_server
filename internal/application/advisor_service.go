package app

import "psycare/internal/domain"

// AdvisorStore _
type AdvisorStore interface {
	CreateAdvisor(advisor *domain.Advisor) error
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
