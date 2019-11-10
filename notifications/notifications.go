/*
Package notifications handles all outbound notifications (email, sms, push, etc)
*/
package notifications

import "github.com/mrz1836/go-mail"

// notificationService is the configuration and services for all notifications
type notificationService struct {
	EmailService *gomail.MailService `json:"email_service"`
}

var (
	// Service is the default settings and connections
	Service *notificationService
)

// StartUp all notification services
func StartUp() (err error) {

	// Start the new service
	Service = new(notificationService)

	// load the email service
	Service.EmailService, err = loadEmailService()

	return
}
