package models

import "time"

type ProfileTable struct {
	ID          uint64    `db:"id"`
	UserId      uint64    `db:"user_id" json:"user_id"`
	Address     string    `db:"address" json:"address"`
	WorkAt      string    `db:"work_at" json:"work_at"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Gender      string    `db:"gender" json:"gender"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
