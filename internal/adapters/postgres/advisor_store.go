package postgres

import (
	"database/sql"
	"fmt"
	"psycare/internal/domain"

	"github.com/jmoiron/sqlx"
)

// AdvisorStore _
type AdvisorStore struct {
	DB *sqlx.DB
}

// CreateAdvisor _
func (as *AdvisorStore) CreateAdvisor(advisor *domain.Advisor) error {
	tx, err := as.DB.Beginx()
	if err != nil {
		return fmt.Errorf("transaction start failed: %w", err)
	}
	_, err = tx.NamedExec(`INSERT into advisors (id, first_name, last_name, description)
				 		   VALUES (:id, :first_name, :last_name, :description)`, advisor)
	if err != nil {
		return fmt.Errorf("failed to insert advisor: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to insert advisor: %w", err)
	}
	return nil
}

// GetAdvisors _
func (as *AdvisorStore) GetAdvisors(verified bool, limit, offset int) (*[]domain.Advisor, error) {
	advisors := &[]domain.Advisor{}
	err := as.DB.Select(advisors, `SELECT id, first_name, last_name, description 
								   FROM advisors WHERE verified=$1 
								   LIMIT $2 OFFSET $3`, verified, limit, offset)

	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to receive advisors: %w", err)
	}
	return advisors, nil
}
