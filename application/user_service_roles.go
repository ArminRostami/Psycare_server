package app

const (
	ROLE_ADMIN   = "admin"
	ROLE_ADVISOR = "advisor"
)

type RoleStore interface {
	GetRoles(id int64) (*[]string, error)
	AddRole(id int64, role string) error
}

func (us *UserService) GetRoles(id int64) (*[]string, error) {
	return us.RoleStore.GetRoles(id)
}

func (us *UserService) AddRole(id int64, role string) error {
	return us.RoleStore.AddRole(id, role)
}

func (us *UserService) HasRole(id int64, roleName string) (bool, error) {
	roles, err := us.RoleStore.GetRoles(id)
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
