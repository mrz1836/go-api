package models

import (
	"context"
	"database/sql"

	"github.com/friendsofgo/errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/mrz1836/go-api/database"
	"github.com/mrz1836/go-api/models/schema"
	"github.com/mrz1836/go-sanitize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var (
	// PersonCreateColumns columns only allowed in create
	PersonCreateColumns = boil.Greylist(
		schema.PersonColumns.Email,
		schema.PersonColumns.FirstName,
		schema.PersonColumns.LastName,
		schema.PersonColumns.MiddleName,
	)

	// PersonUpdateColumns columns only allowed in update
	PersonUpdateColumns = boil.Whitelist(
		schema.PersonColumns.Email,
		schema.PersonColumns.FirstName,
		schema.PersonColumns.LastName,
		schema.PersonColumns.MiddleName,
	)

	// PersonDeleteColumns columns only allowed in delete
	PersonDeleteColumns = boil.Whitelist(
		schema.PersonColumns.IsDeleted,
	)

	// PersonAllFields all fields that can be displayed
	PersonAllFields = []string{
		schema.PersonColumns.CreatedAt,
		schema.PersonColumns.Email,
		schema.PersonColumns.FirstName,
		schema.PersonColumns.ID,
		schema.PersonColumns.LastName,
		schema.PersonColumns.MiddleName,
		schema.PersonColumns.ModifiedAt,
	}
)

// Person extends the schema model
type Person struct {
	schema.Person
}

// NewPerson creates an empty person model
func NewPerson() *Person {
	return &Person{schema.Person{}}
}

// NewPersonUsingSchema creates a person model using a schema
func NewPersonUsingSchema(person schema.Person) *Person {
	return &Person{person}
}

// NewPersonsUsingSchema creates a new model using a schema
func NewPersonsUsingSchema(personSlice schema.PersonSlice) []Person {
	persons := make([]Person, len(personSlice))
	for _, person := range personSlice {
		persons = append(persons, Person{*person})
	}
	return persons
}

// GetPersonByID gets a person by ID
func GetPersonByID(id uint64) (person *Person, err error) {

	// Start with a schema
	var p *schema.Person

	// Find the associated record
	p, err = schema.FindPerson(context.Background(), database.ReadDatabase, id) // todo: turn slice of strings into variadic
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return
	}

	// Create a new model with existing schema
	person = NewPersonUsingSchema(*p)

	return
}

// GetPersonByEmail gets a person by email address
func GetPersonByEmail(email string) (person *Person, err error) {

	// Start with a schema
	var p *schema.Person

	// Find the associated record
	p, err = schema.Persons(qm.Where(schema.PersonColumns.Email+" = ?", email)).One(context.Background(), database.ReadDatabase)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return
	}

	// Create a new model with existing schema
	person = NewPersonUsingSchema(*p)

	return
}

// GetPersons gets an array of person  //todo: temporary for now
func GetPersons() (persons []Person, err error) {

	// Start with a schema
	var p schema.PersonSlice

	// Find the associated record
	p, err = schema.Persons(
		qm.Where(schema.PersonColumns.IsDeleted+" = ?", 0)).All(context.Background(), database.ReadDatabase)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return
	}

	// Create a new model with existing schema
	persons = NewPersonsUsingSchema(p)

	return
}

// BeforeValidate runs before validate (sanitizing, formatting, default values)
func (p *Person) BeforeValidate() {

	// Always sanitize
	p.Email = sanitize.Email(p.Email, false)

	// Treat names as "formal"
	p.FirstName = sanitize.FormalName(p.FirstName)
	p.MiddleName = sanitize.FormalName(p.MiddleName)
	p.LastName = sanitize.FormalName(p.LastName)
}

// Validate checks the model, struct and any custom validations
func (p *Person) Validate() error {

	// Runs before (sanitizing, formatting, default values)
	p.BeforeValidate()

	// Custom validations

	// Run the struct validations
	return validation.ValidateStruct(p,
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.FirstName, validation.Length(0, 50)),
		validation.Field(&p.LastName, validation.Length(0, 50)),
		validation.Field(&p.MiddleName, validation.Length(0, 50)),
	)
}

// Save either inserts or updates a model
func (p *Person) Save(columns boil.Columns, tx *sql.Tx) (rowsAffected int64, err error) {

	// Validate the model
	err = p.Validate()
	if err != nil {
		return
	}

	// Try to insert the model
	if p.ID == 0 {
		rowsAffected = 1
		err = p.Insert(context.Background(), tx, columns)
	} else {
		rowsAffected, err = p.Update(context.Background(), tx, columns)
	}

	return
}
