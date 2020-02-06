package postgres

import (
	"database/sql"
	"fmt"
	"psycare/internal/domain"
)

// AdvisorStore _
type AdvisorStore struct {
	DB *DB
}

// CreateAdvisor _
func (as *AdvisorStore) CreateAdvisor(advisor *domain.Advisor) error {
	return as.DB.namedExec(`INSERT into advisors (id, first_name, last_name, description)
						   VALUES (:id, :first_name, :last_name, :description)`, advisor)
}

// GetAdvisors _
func (as *AdvisorStore) GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error) {
	advisors := &[]domain.Advisor{}
	err := as.DB.Con.Select(advisors, `SELECT id, first_name, last_name, description 
								   FROM advisors WHERE verified=$1 
								   LIMIT $2 OFFSET $3`, verified, limit, offset)

	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to receive advisors: %w", err)
	}
	return advisors, nil
}
