package person

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deliverydudes/go-library/utils/logger"
	"github.com/julienschmidt/httprouter"
	"github.com/mrz1836/go-api-router"
	"github.com/mrz1836/go-api/database"
	"github.com/mrz1836/go-api/models/schema"
	"github.com/satori/go.uuid"
	"github.com/volatiletech/sqlboiler/boil"
)

// RegisterRoutes register all the package specific routes
func RegisterRoutes(router *apirouter.Router) {
	router.HTTPRouter.POST("/person", router.Request(create))
	router.HTTPRouter.OPTIONS("/person", router.SetCrossOriginHeaders)
}

// create makes a new model
func create(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	// Create the model
	email := uuid.NewV4().String()
	u := &schema.Person{Email: email + "@gmail.com"}

	// Try to insert the model
	err := u.Insert(context.Background(), database.WriteConnection, boil.Greylist(schema.PersonColumns.Email))
	if err != nil {
		logger.Data(2, logger.ERROR, fmt.Sprintf("error: %s", err))
		apirouter.ReturnResponse(w, http.StatusExpectationFailed, fmt.Sprintf("error creating person: %s", err), true)
		return
	}

	// Good!
	_, _ = fmt.Fprint(w, "Created!\n")
}
