package app

import "psycare/domain"

type AdvisorStore interface {
	CreateAdvisor(advisor *domain.Advisor) error
	GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error)
	AddSchedule(sch *domain.Schedule) error
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
	// err = as.RoleStore.AddRole(advisor.ID, ROLE_ADVISOR)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (as *AdvisorService) GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error) {
	return as.AdvStore.GetAdvisors(verified, limit, offset)
}

func (as *AdvisorService) AddSchedule(sch *domain.Schedule) error {
	return as.AdvStore.AddSchedule(sch)
}
