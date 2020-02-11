package domain

type User struct {
	UserName string    `db:"username" json:"username" validate:"required"`
	Email    string    `db:"email" json:"email" validate:"required,email"`
	Password string    `db:"password" json:"password" validate:"required"`
	ID       int64     `db:"id"`
	Credit   int64     `db:"credit"`
	Roles    *[]string `json:"roles"`
}
