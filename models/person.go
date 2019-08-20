package models

import (
	"context"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/mrz1836/go-api/database"
	"github.com/mrz1836/go-api/models/schema"
	"github.com/mrz1836/go-sanitize"
	"github.com/volatiletech/sqlboiler/boil"
)

// Person extends the schema model
type Person struct {
	schema.Person
}

// NewPerson person model
func NewPerson() Person {
	return Person{schema.Person{}}
}

// BeforeValidate runs before validate (sanitizing, formatting, default values)
func (p *Person) BeforeValidate() {

	// Always sanitize
	p.Email = sanitize.Email(p.Email, false)
}

// Validate checks the model, struct and any custom validations
func (p *Person) Validate() error {

	// Runs before (sanitizing, formatting, default values)
	p.BeforeValidate()

	// Custom validations

	// Run the struct validations
	return validation.ValidateStruct(p,
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.FirstName, validation.Required),
		validation.Field(&p.LastName, validation.Required),
	)
}

// Save either inserts or updates a model
func (p *Person) Save(ctx context.Context, columns boil.Columns) (err error) {

	// Validate the model
	err = p.Validate()
	if err != nil {
		return
	}

	// Try to insert the model
	if p.ID == 0 {
		err = p.Insert(ctx, database.WriteDatabase, columns)
	} else {
		//var rowsAffected int64
		//rowsAffected, err = p.Update(context.Background(), database.WriteDatabase, columns)
		_, err = p.Update(ctx, database.WriteDatabase, columns)
	}

	return
}

/*func PersonExists(ctx context.Context, db *sql.DB, p *schema.Person) (bool, error) {
	return schema.PersonExists(ctx, db, p.ID)
}*/
