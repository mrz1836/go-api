package models

import (
	"html/template"
	"path/filepath"

	"github.com/mrz1836/go-api/config"
	"github.com/mrz1836/go-api/notifications"
	"github.com/mrz1836/go-mail"
)

// Define the template vars
var (
	emailPersonExampleHTML *template.Template
	emailPersonExampleText *template.Template
)

// loadPersonEmails load the email templates (fired from StartUp)
func loadPersonEmails() (err error) {

	// Load the email
	email := notifications.Service.EmailService.NewEmail()
	currentDirectory := config.GetCurrentDir()

	// Parse the text version, store in local memory
	emailPersonExampleText, err = email.ParseTemplate(filepath.Join(currentDirectory, "..", "static", "views", "emails", "persons", "example_email.txt"))
	if err != nil {
		return
	}

	// Parse the html version, store in local memory
	emailPersonExampleHTML, err = email.ParseHTMLTemplate(filepath.Join(currentDirectory, "..", "static", "views", "emails", "persons", "example_email.html"))

	return
}

// EmailExampleData is the email struct for the data
type EmailExampleData struct {
	Person
	SupportEmail string `json:"support_email"`
}

// SendExampleEmail sends an example email
func (p *Person) SendExampleEmail() (err error) {

	// Create the data struct
	data := new(EmailExampleData)
	data.Person = *p
	data.SupportEmail = "support@example.com"

	// Start a new email
	email := notifications.Service.EmailService.NewEmail()
	email.Recipients = append(email.Recipients, data.Person.Email)
	email.FromName = "Acme"
	email.Subject = "Your example email subject line"
	email.Tags = append(email.Tags, "example_tag")
	email.TrackOpens = true

	// Apply the templates
	err = email.ApplyTemplates(emailPersonExampleHTML, emailPersonExampleText, data)
	if err != nil {
		return
	}

	// Send the email
	err = notifications.Service.EmailService.SendEmail(email, gomail.SMTP)

	return
}
