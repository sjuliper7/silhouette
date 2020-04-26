package model

// Email ...
type Email struct {
	Receivers   []string
	CCReceivers []string
	Subject     string
	Message     string
	Sender      string
}

// MailDial ...
type MailDial struct {
	SMTPHost string
	SMTPPort string
	Email    string
	Password string
}
