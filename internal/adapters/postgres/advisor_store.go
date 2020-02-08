package postgres

import (
	"database/sql"
	"fmt"
	"psycare/internal/domain"
)

type AdvisorStore struct {
	DB *PDB
}

func (as *AdvisorStore) CreateAdvisor(advisor *domain.Advisor) error {
	return as.DB.namedExec(`INSERT into advisors (id, first_name, last_name, description)
						   VALUES (:id, :first_name, :last_name, :description)`, advisor)

}

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
