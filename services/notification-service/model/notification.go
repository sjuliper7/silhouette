package model

type (
	//Notification ...
	Notification struct {
		Type  string `json:"type"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}
)
