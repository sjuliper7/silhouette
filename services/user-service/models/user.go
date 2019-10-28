package models

//User struct is represntative form table user
type User struct {
	ID       uint64  `json:"id" db:"id"`
	Username string  `json:"username" db:"username"`
	Email    string  `json:"email" db:"email"`
	Name     string  `json:"name" db:"name"`
	Role     string  `json:"role" db:"role"`
	Profile  Profile `json:"profile"`
}
