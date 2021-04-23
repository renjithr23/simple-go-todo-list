package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

type Configuration struct {
	Port             int
	StorageType      string
	ConnectionString string
}

func main() {

	// Reading and decoding configuraiotn files
	var config Configuration

	file, err := os.Open("./conf.json")
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	// Initializeing Storage Type.
	stg, err := NewStorage(StorageType(config.StorageType))
	if err != nil {
		log.Fatal(err)
	}

	// Defyning Routes for Todo Resource
	r := mux.NewRouter()

	r.HandleFunc("/todos", GetTodos(stg)).Methods("GET")
	r.HandleFunc("/todos/{id}", GetTodo(stg)).Methods("GET")
	r.HandleFunc("/todos/{id}", DeleteTodo(stg)).Methods("DELETE")
	r.HandleFunc("/todos", AddTodo(stg)).Methods("POST")
	http.Handle("/", r)

	// Defining routes that serve the swagger UI
	getR := r.Methods(http.MethodGet).Subrouter()
	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// Starting the Server
	fmt.Println("The beer server is on tap at http://localhost:8080.")
	log.Print(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))

	// Cleanig up
	stg.DeleteStorage()
}
