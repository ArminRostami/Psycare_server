package postgres

import (
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
