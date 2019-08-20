// Package persons are the actions associated with the person model
package persons

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-api-router"
	"github.com/mrz1836/go-api/models"
)

// RegisterRoutes register all the package specific routes
func RegisterRoutes(router *apirouter.Router) {
	router.HTTPRouter.POST("/persons", router.Request(create))
	router.HTTPRouter.OPTIONS("/persons", router.SetCrossOriginHeaders)
}

// create makes a new model
func create(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	// Get the parameters
	params := apirouter.GetParams(req)

	// Create the model
	person := models.NewPerson()

	// Set the values
	person.Email = params.Get("email")
	person.FirstName = params.Get("first_name")
	person.LastName = params.Get("last_name")

	// Save will insert a new person since we are creating a new model
	err := person.Save(context.Background(), models.PersonCreateColumns)
	if err != nil {
		apirouter.ReturnResponse(w, http.StatusExpectationFailed, fmt.Sprintf("failed to save person: %s", err), false)
		return
	}

	// This should not fail on the encode
	_ = apirouter.ReturnJSONEncode(w, http.StatusCreated, json.NewEncoder(w), person, models.PersonAllFields)
}
