// Package models extends the schema package for model management
package models

// StartUp loads all model dependencies
func StartUp() (err error) {

	// Load person templates into memory
	err = loadPersonEmails()

	return
}
