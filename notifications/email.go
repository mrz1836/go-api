package notifications

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/mrz1836/go-api/config"
	gomail "github.com/mrz1836/go-mail"
)

// loadEmailService loads the email service
func loadEmailService() (service *gomail.MailService, err error) {

	// Define your service configuration
	service = new(gomail.MailService)
	service.FromDomain = config.Values.Email.FromDomain
	service.FromName = config.Values.Email.FromName
	service.FromUsername = config.Values.Email.FromUsername

	// Mandrill
	service.MandrillAPIKey = config.Values.Email.MandrillAPIKey

	// AWS SES
	service.AwsSesAccessID = config.Values.Email.AwsSesAccessID
	service.AwsSesSecretKey = config.Values.Email.AwsSesSecretKey

	// Postmark
	service.PostmarkServerToken = config.Values.Email.PostmarkServerToken

	// SMTP
	service.SMTPHost = config.Values.Email.SMTPHost
	service.SMTPPassword = config.Values.Email.SMTPPassword
	service.SMTPPort = config.Values.Email.SMTPPort
	service.SMTPUsername = config.Values.Email.SMTPUsername

	// Set default styles
	service.EmailCSS, err = ioutil.ReadFile(filepath.Join(config.GetCurrentDir(), "..", "static", "css", "email.css"))
	if err != nil {
		err = fmt.Errorf("error loading email styles: %w", err)
		return
	}

	// Start and return the service
	err = service.StartUp()

	return
}
