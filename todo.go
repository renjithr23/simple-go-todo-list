package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route GET /todos todos getTodos
// Return a list of Todos from the database
// responses:
//	200: todosResponse
//  404: errorResponse
// Consumes:
// - application/json
// produces:
// - application/json
// Schemes: http

// GetTodos handles GET requests and returns all current Todos
func GetTodos(storage Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		todos, err := storage.ListTodo()

		w.Header().Set("Content-Type", "application/json")
		response, err := json.Marshal(todos)
		if err != nil {
			log.Print(err)
			err_response, _ := json.Marshal(GenericError{"Server error"})

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(err_response)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// swagger:route DELETE /Todo/{id} todos deleteTodo
// Deletes a Todo
// Consumes:
// - application/json
// produces:
// - application/json
// Schemes: http
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func DeleteTodo(storage Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		err := storage.DeleteTodo(id)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			log.Print(err)
			err_response, _ := json.Marshal(GenericError{"Todo with ID " + strconv.Itoa(id) + " does not exist"})

			w.WriteHeader(http.StatusBadRequest)
			w.Write(err_response)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// swagger:route POST /todos todos addTodo
// Adds a new Todo Item
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse
// Consumes:
// - application/json
// produces:
// - application/json
// Schemes: http
func AddTodo(storage Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var t Todo
		err := json.NewDecoder(r.Body).Decode(&t)

		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		storage.CreateTodo(t)
	}
}

// Swagger documentation models

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// Data structure representing a single Todo List Item
// swagger:response todoResponse
type todoResponseWrapper struct {
	// Newly created Todo List Item
	// in: body
	Body Todo
}

// Data structure representing a List of Todo List Items
// swagger:response todosResponse
type todosResponseWrapper struct {
	// Newly created Todo List Item
	// in: body
	Body []Todo
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters deleteTodo
type productIDParamsWrapper struct {
	// The id of the Todo List Item for which the operation relates
	// in: path
	// required: true
	Id int `json:"id"`
}
