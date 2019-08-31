// Package persons are the actions associated with the person model
package persons

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-api-router"
	"github.com/mrz1836/go-api/config"
	"github.com/mrz1836/go-api/models"
	"github.com/mrz1836/go-api/models/schema"
	"github.com/mrz1836/go-logger"
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
	params, _ := apirouter.GetParams(req)

	// Create the model
	person := models.NewPerson()

	// Set the values
	person.Email = params.Get(schema.PersonColumns.Email)
	person.FirstName = params.Get(schema.PersonColumns.FirstName)
	person.MiddleName = params.Get(schema.PersonColumns.MiddleName)
	person.LastName = params.Get(schema.PersonColumns.LastName)

	// Save will insert a new person since we are creating a new model
	_, err := person.Save(context.Background(), models.PersonCreateColumns)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, "error in save method", fmt.Sprintf("error creating person: %s", err.Error()), http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, apiError)
		return
	}

	// This should not fail on the encode
	_ = apirouter.ReturnJSONEncode(w, http.StatusCreated, json.NewEncoder(w), person, models.PersonAllFields)
}

// updatePerson modifies an existing model
func updatePerson(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Get the parameters
	params, _ := apirouter.GetParams(req)

	// Permit only fields allowed to update
	apirouter.PermitParams(params, models.PersonUpdateColumns.Cols)

	// Get the model by ID
	//todo: replace with params.GetUint64()
	id, err := strconv.ParseUint(params.Get(schema.PersonColumns.ID), 10, 64)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error parsing uint in params: %s", err.Error()), "unable to update person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, apiError)
		return
	}
	person, err := models.GetPersonByID(id)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error getting related person: %s", err.Error()), "unable to update person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, apiError)
		return
	}

	// Test to see if deleted
	if person.IsDeleted.Bool {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("person is marked as deleted: %d", id), "unable to update a deleted record", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, apiError)
		return
	}

	// todo:  api router - imbue model

	// todo: using imbue, set the edited fields

	person.FirstName = params.Get(schema.PersonColumns.FirstName)

	// Save will update an exiting person
	affected, err := person.Save(context.Background(), models.PersonUpdateColumns)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, "error in save method, updating person failed", fmt.Sprintf("error updating person: %s", err.Error()), http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, apiError)
		return
	}

	logger.Printf("affected %d rows", affected)

	// This should not fail on the encode
	_ = apirouter.ReturnJSONEncode(w, http.StatusOK, json.NewEncoder(w), person, models.PersonAllFields)
}

// deletePerson will mark a record as deleted
func deletePerson(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Get the parameters
	params, _ := apirouter.GetParams(req)

	//todo: api router - permit fields

	// Get the model by ID
	//todo: replace with params.GetUint64()
	id, err := strconv.ParseUint(params.Get(schema.PersonColumns.ID), 10, 64)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error parsing uint in params: %s", err.Error()), "unable to delete person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, apiError)
		return
	}
	person, err := models.GetPersonByID(id)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, fmt.Sprintf("error getting related person: %s", err.Error()), "unable to delete person", http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, apiError)
		return
	}

	// Already deleted?
	if person.IsDeleted.Bool {
		_ = apirouter.ReturnJSONEncode(w, http.StatusOK, json.NewEncoder(w), person, models.PersonAllFields)
		return
	}

	// Not deleted, let's update
	person.IsDeleted.Bool = true

	// Save will update an exiting person
	_, err = person.Save(context.Background(), models.PersonDeleteColumns)
	if err != nil {
		apiError := apirouter.ErrorFromRequest(req, "error in save method, deleting person failed", fmt.Sprintf("error deleting person: %s", err.Error()), http.StatusExpectationFailed, "")
		apirouter.ReturnResponse(w, req, http.StatusExpectationFailed, apiError)
		return
	}

	// This should not fail on the encode
	_ = apirouter.ReturnJSONEncode(w, http.StatusOK, json.NewEncoder(w), person, models.PersonAllFields)
}
