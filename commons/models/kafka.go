package models

//MessageKa fkaAfterRegistration is struct for handling message as payload message
type MessageKafkaAfterRegistration struct {
	UserID      int64  `json:"user_id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	WorkAt      string `json:"work_at"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
}
