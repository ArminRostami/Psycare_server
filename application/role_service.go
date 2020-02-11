package app

const (
	ROLE_ADMIN   = "admin"
	ROLE_ADVISOR = "advisor"
)

type RoleStore interface {
	GetRoles(id int64) (*[]string, error)
	AddRole(id int64, role string) error
}

type RoleService struct {
	Store RoleStore
}

func (rs *RoleService) GetRoles(id int64) (*[]string, error) {
	return rs.Store.GetRoles(id)
}

func (rs *RoleService) AddRole(id int64, role string) error {
	return rs.Store.AddRole(id, role)
}

func (rs *RoleService) HasRole(id int64, roleName string) (bool, error) {
	roles, err := rs.Store.GetRoles(id)
	if err != nil {
		return false, err
	}
	for _, role := range *roles {
		if role == roleName {
			return true, nil
		}
	}
	return false, nil
}
