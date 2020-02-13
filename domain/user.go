package domain

type User struct {
	UserName string    `db:"username" json:"username" validate:"required"`
	Email    string    `db:"email" json:"email" validate:"required,email"`
	Password string    `db:"password" json:"password" validate:"required"`
	ID       int64     `db:"id" json:"id,omitempty"`
	Credit   int64     `db:"credit" json:"credit,omitempty"`
	Roles    *[]string `json:"roles,omitempty"`
}
