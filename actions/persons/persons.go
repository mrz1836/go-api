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
	"github.com/mrz1836/go-api/models/schema"
	"github.com/volatiletech/sqlboiler/boil"
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

	// Save will insert a new person
	err := person.Save(context.Background(), boil.Greylist(schema.PersonColumns.Email, schema.PersonColumns.FirstName, schema.PersonColumns.LastName))
	if err != nil {
		apirouter.ReturnResponse(w, http.StatusExpectationFailed, fmt.Sprintf("failed to save person: %s", err), false)
		return
	}

	// Encode the model for return
	b, err := json.Marshal(person)
	if err != nil {
		apirouter.ReturnResponse(w, http.StatusExpectationFailed, fmt.Sprintf("encoding person failed: %s", err), false)
		return
	}

	// Send the model back
	apirouter.ReturnResponse(w, http.StatusCreated, string(b), true)
}
