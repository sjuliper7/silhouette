package models

import "time"

//User struct is return value as response
type User struct {
	ID        uint64    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Name      string    `json:"name" db:"name"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Profile   Profile   `json:"profile"`
}

//User struct is representative form table user
type UserTable struct {
	ID        uint64    `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Name      string    `db:"name"`
	Role      string    `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
