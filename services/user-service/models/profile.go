package models

type Profile struct {
	ID          int64  `json:"id"`
	Address     string `json:"address"`
	WorkAt      string `json:"work_at"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
}
