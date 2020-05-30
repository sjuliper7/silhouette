package email

import (
	"github.com/sirupsen/logrus"
	"github.com/sjuliper7/silhouette/services/notification-service/model"
	"github.com/sjuliper7/silhouette/services/notification-service/repository"
	"gopkg.in/gomail.v2"
)

type emailRepository struct {
	dialer *gomail.Dialer
}

// NewEmailRepository ...
func NewEmailRepository(dialer *gomail.Dialer) repository.EmailRepository {
	return &emailRepository{dialer: dialer}
}

func (repo *emailRepository) SentEmail(email model.Email) error {
	mailer := gomail.NewMessage()

	mailer.SetHeader("From", email.Sender)
	mailer.SetHeader("To", email.Receivers...)
	mailer.SetHeader("Subject", email.Subject)
	mailer.SetBody("text/html", email.Message)

	err := repo.dialer.DialAndSend(mailer)
	if err != nil {
		logrus.Errorf("[repository][email_repository] error when sent email: %v", err)
		return err
	}

	logrus.Infof("Yay email have sent to, %v", email.Receivers)

	return nil
}
