package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	stg, err := NewStorage(Memory)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/todos", GetTodos(stg)).Methods("GET")
	r.HandleFunc("/todos/delete/{id}", DeleteTodo(stg)).Methods("GET")
	r.HandleFunc("/todos/add", AddTodo(stg)).Methods("POST")
	http.Handle("/", r)

	// Starting the Server
	fmt.Println("The beer server is on tap at http://localhost:8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
