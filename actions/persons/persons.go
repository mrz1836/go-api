// Package persons are the actions associated with the person model
package persons

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-api-router"
	"github.com/mrz1836/go-api/config"
	"github.com/mrz1836/go-api/database"
	"github.com/mrz1836/go-api/models"
	"github.com/mrz1836/go-api/models/schema"
)

// RegisterRoutes register all the package specific routes
func RegisterRoutes(router *apirouter.Router) {
	router.HTTPRouter.POST("/persons", router.BasicAuth(router.Request(createPerson), config.Values.BasicAuth.User, config.Values.BasicAuth.Password, config.Values.UnauthorizedError))
	router.HTTPRouter.PUT("/persons", router.BasicAuth(router.Request(updatePerson), config.Values.BasicAuth.User, config.Values.BasicAuth.Password, config.Values.UnauthorizedError))
	router.HTTPRouter.DELETE("/persons", router.BasicAuth(router.Request(deletePerson), config.Values.BasicAuth.User, config.Values.BasicAuth.Password, config.Values.UnauthorizedError))
	router.HTTPRouter.OPTIONS("/persons", router.SetCrossOriginHeaders)
}

// createPerson makes a new model
func createPerson(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Get the parameters
	params := apirouter.GetParams(req)

	// Create the model
	person := models.NewPerson()

	// Set the values
	person.Email = params.GetString(schema.PersonColumns.Email)
	person.FirstName = params.GetString(schema.PersonColumns.FirstName)
	person.MiddleName = params.GetString(schema.PersonColumns.MiddleName)
	person.LastName = params.GetString(schema.PersonColumns.LastName)

	// Check missing value
	if len(person.Email) == 0 {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("missing field: %s", schema.PersonColumns.Email), fmt.Sprintf("error creating person - missing field: %s", schema.PersonColumns.Email), http.StatusBadRequest, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Get existing person?
	existingPerson, err := models.GetPersonByEmail(person.Email)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, "error getting existing person", fmt.Sprintf("error getting existing offer: %s", err.Error()), http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	} else if existingPerson != nil && existingPerson.IsDeleted.Bool {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("person has been deleted: %s", person.Email), "account has been disabled", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Start a new transaction
	var tx *sql.Tx
	tx, _, err = database.NewTx(config.DatabaseDefaultTxTimeout)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error creating tx: %s", err.Error()), "error creating person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Show the existing person
	if existingPerson != nil && existingPerson.ID > 0 {
		// This should not fail on the encode
		_ = apirouter.ReturnJSONEncode(w, http.StatusCreated, json.NewEncoder(w), existingPerson, models.PersonAllFields)
		return
	}

	// Save will insert a new person since we are creating a new model
	_, err = person.Save(models.PersonCreateColumns, tx)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error creating person: %s", err.Error()), fmt.Sprintf("error creating person: %s", err.Error()), http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Commit!
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error creating person: %s", err.Error()), "error creating person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// This should not fail on the encode
	_ = apirouter.ReturnJSONEncode(w, http.StatusCreated, json.NewEncoder(w), person, models.PersonAllFields)
}

// updatePerson modifies an existing model
func updatePerson(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Get the parameters
	params := apirouter.GetParams(req)

	// Get the model by ID
	id := params.GetUint64(schema.PersonColumns.ID)

	person, err := models.GetPersonByID(id)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error getting related person: %s", err.Error()), "unable to update person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Test to see if deleted
	if person.IsDeleted.Bool {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("person is marked as deleted: %d", id), "unable to update a deleted record", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Set the first name
	person.FirstName = params.GetString(schema.PersonColumns.FirstName)

	// Start a new transaction
	var tx *sql.Tx
	tx, _, err = database.NewTx(config.DatabaseDefaultTxTimeout)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error updating person: %s", err.Error()), "error updating person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Save will update an exiting person
	//var affected int64
	_, err = person.Save(models.PersonUpdateColumns, tx)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error saving person: %s", err.Error()), fmt.Sprintf("error updating person: %s", err.Error()), http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Commit
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error in commit creating person: %s", err.Error()), "error creating person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// This should not fail on the encode
	_ = apirouter.ReturnJSONEncode(w, http.StatusOK, json.NewEncoder(w), person, models.PersonAllFields)
}

// deletePerson will mark a record as deleted
func deletePerson(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Get the parameters
	params := apirouter.GetParams(req)

	// Get the model by ID
	id := params.GetUint64(schema.PersonColumns.ID)

	person, err := models.GetPersonByID(id)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error getting related person: %s", err.Error()), "unable to delete person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Already deleted?
	if person.IsDeleted.Bool {
		_ = apirouter.ReturnJSONEncode(w, http.StatusOK, json.NewEncoder(w), person, models.PersonAllFields)
		return
	}

	// Not deleted, let's update
	person.IsDeleted.Bool = true

	// Start a new transaction
	var tx *sql.Tx
	tx, _, err = database.NewTx(config.DatabaseDefaultTxTimeout)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error creating tx: %s", err.Error()), "error deleting person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Save will update an exiting person
	_, err = person.Save(models.PersonDeleteColumns, tx)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error saving person: %s", err.Error()), fmt.Sprintf("error deleting person: %s", err.Error()), http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// Commit
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error in commit: %s", err.Error()), "error deleting person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, apiError.Code, apiError)
		return
	}

	// This should not fail on the encode
	_ = apirouter.ReturnJSONEncode(w, http.StatusOK, json.NewEncoder(w), person, models.PersonAllFields)
}
