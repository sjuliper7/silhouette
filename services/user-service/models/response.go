package models

//Response ...
type Response struct {
	Code    int         `json:"code,omitempty"`
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
