package models

import "time"

type ProfileTable struct {
	ID          uint64    `db:"id"`
	UserId      uint64    `db:"user_id"`
	Address     string    `db:"address"`
	WorkAt      string    `db:"work_at"`
	PhoneNumber string    `db:"phone_number"`
	Gender      string    `db:"gender"`
	IsActive    bool      `db:"is_active"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
