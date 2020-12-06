package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTodos(storage Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		todos, err := storage.ListTodo()
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(todos)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func DeleteTodo(storage Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		err := storage.DeleteTodo(id)
		if err != nil {
			log.Fatal(err)
		}

	}
}

func AddTodo(storage Storage) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var t Todo

		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			fmt.Println("Marshalling Error logged")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		storage.CreateTodo(t)

	}
}
