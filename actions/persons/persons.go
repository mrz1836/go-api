// Package persons are the actions associated with the person model
package persons

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-api-router"
	"github.com/mrz1836/go-api/config"
	"github.com/mrz1836/go-api/models"
	"github.com/mrz1836/go-api/models/schema"
)

// RegisterRoutes register all the package specific routes
func RegisterRoutes(router *apirouter.Router) {
	router.HTTPRouter.POST("/persons", router.BasicAuth(router.Request(create), config.Values.BasicAuth.User, config.Values.BasicAuth.Password, config.Values.UnauthorizedError))
	router.HTTPRouter.OPTIONS("/persons", router.SetCrossOriginHeaders)
}

// create makes a new model
func create(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Get the parameters
	params := apirouter.GetParams(req)

	// Create the model
	person := models.NewPerson()

	// Set the values
	person.Email = params.Get(schema.PersonColumns.Email)
	person.FirstName = params.Get(schema.PersonColumns.FirstName)
	person.MiddleName = params.Get(schema.PersonColumns.MiddleName)
	person.LastName = params.Get(schema.PersonColumns.LastName)

	// Save will insert a new person since we are creating a new model
	err := person.Save(context.Background(), models.PersonCreateColumns)
	if err != nil {
		apirouter.ReturnResponse(w, http.StatusExpectationFailed, fmt.Sprintf("failed to save person: %s", err), false)
		return
	}

	// This should not fail on the encode
	_ = apirouter.ReturnJSONEncode(w, http.StatusCreated, json.NewEncoder(w), person, models.PersonAllFields)
}
