package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vjftw/orchestrate/master/managers"
	"github.com/vjftw/orchestrate/master/models"
)

// UserController - Handles actions that can be performed on Users
type UserController struct {
	UserManager managers.IManager `inject:"manager user"`
}

// AddRoutes - Adds the routes assosciated to this controller
func (uC *UserController) AddRoutes(r *mux.Router) {
	r.
		HandleFunc("/v1/users", uC.postHandler).
		Name("POSTusers").
		Methods("POST")
}

func (uC *UserController) postHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Unmarshal request into user variable
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// 400 on Error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// validate the user variable
	err = uC.UserManager.Validate(user)
	if err != nil {
		// 400 on Error
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Persist the user variable
	uC.UserManager.Save(&user)

	// write the user variable to output and set http header to 201
	Respond(w, http.StatusCreated, user)
}
