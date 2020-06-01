package models

import (
	"database/sql"
	"time"
)

type ProfileTable struct {
	ID          int64     `db:"id" json:"id"`
	UserID      int64     `db:"user_id" json:"user_id"`
	Name        string    `db:"name" json:"name"`
	Address     string    `db:"address" json:"address"`
	WorkAt      string    `db:"work_at" json:"work_at"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	DateOfBirth string    `db:"date_of_birth" json:"date_of_birth"`
	Gender      string    `db:"gender" json:"gender"`
	IsActive    bool      `db:"is_active" json:"is_active"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type PortfolioTableScanner struct {
	ID          sql.NullInt64  `db:"id"`
	UserID      sql.NullInt64  `db:"user_id"`
	Name        sql.NullString `db:"name"`
	Address     sql.NullString `db:"address"`
	WorkAt      sql.NullString `db:"work_at"`
	PhoneNumber sql.NullString `db:"phone_number"`
	DateOfBirth sql.NullString `db:"date_of_birth"`
	Gender      sql.NullString `db:"gender"`
	IsActive    sql.NullBool   `db:"is_active"`
	CreatedAt   sql.NullString `db:"created_at"`
	UpdatedAt   sql.NullString `db:"updated_at"`
}

type OutputKafkaProfile struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	WorkAt      string `json:"work_at"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	IsActive    bool   `json:"is_active"`
}
