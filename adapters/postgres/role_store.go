package postgres

import (
	"github.com/pkg/errors"
)

type RoleStore struct {
	DB *PDB
}

func (r *RoleStore) GetRoles(id int64) (*[]string, error) {
	roles := &[]string{}
	err := r.DB.Con.Select(roles, `
	SELECT name 
	FROM (SELECT role_id FROM user_roles WHERE user_id=$1) as ur 
	INNER JOIN roles on ur.role_id=roles.id `, id)
	return roles, errors.Wrap(err, "failed to get roles")
}

func (r *RoleStore) AddRole(id int64, role string) error {
	err := r.DB.Execute(`
	INSERT INTO user_roles (user_id, role_id) 
	VALUES ($1, (SELECT id FROM roles WHERE name=$2))`, id, role)
	return errors.Wrap(err, "failed to add role")
}
