package models

//User struct is represntative form table user
type User struct {
	ID       uint64 `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	LastName string `json:"last_name" db:"last_name"`
}
