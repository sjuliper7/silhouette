package models

//Profile ...
type Profile struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Address     string `json:"address"`
	Name        string `json:"name"`
	WorkAt      string `json:"work_at"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
