package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Configuration struct {
	Port             int
	StorageType      string
	ConnectionString string
}

func main() {

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

	stg, err := NewStorage(StorageType(config.StorageType))
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
	log.Print(http.ListenAndServe(":"+strconv.Itoa(config.Port), nil))

	stg.DeleteStorage()
}
