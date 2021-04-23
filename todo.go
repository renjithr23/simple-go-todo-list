package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
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
		response, _ := json.Marshal(todos)
		if err != nil {
			log.Print(err)
			err_response, _ := json.Marshal(InternalServerError("InternalServer Error"))

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(err_response)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// swagger:route GET /Todo/{id} todos getTodo
// Gets a Todo item from it's Id
// Consumes:
// - application/json
// produces:
// - application/json
// Schemes: http
// responses:
//	200: todosResponse
//  404: errorResponse
//  501: errorResponse

// GetTodo handles GET requests and returns a Todo based on an Id
func GetTodo(storage Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		todos, err := storage.GetTodo(id)

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(todos)
		if err != nil {
			log.Print(err)
			err_response, _ := json.Marshal(BadRequestError("Todo with ID " + strconv.Itoa(id) + " does not exist"))

			w.WriteHeader(http.StatusBadRequest)
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
			err_response, _ := json.Marshal(BadRequestError("Todo with ID " + strconv.Itoa(id) + " does not exist"))

			w.WriteHeader(http.StatusBadRequest)
			w.Write(err_response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(DeletedResponse("Todo with ID " + strconv.Itoa(id) + " is deleted"))
		w.Write(response)
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
		w.Header().Set("Content-Type", "application/json")
		var t Todo
		err := json.NewDecoder(r.Body).Decode(&t)

		errs := validation.ValidateStruct(&t,
			validation.Field(&t.Name, validation.Required),
			validation.Field(&t.Description, validation.Required),
		)
		if errs != nil {
			err_response, _ := json.Marshal(InvalidInputError(errs.(validation.Errors)))

			w.WriteHeader(http.StatusBadRequest)
			w.Write(err_response)
			return
		}

		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		storage.CreateTodo(t)

		response, _ := json.Marshal(CreatedResponse("Todo Created"))
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}
}
