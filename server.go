package main 

import (
	"fmt"	
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func getTodoHandler(todos *[]Todo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){
		
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(*todos)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func deleteTodoHandler(todos *[]Todo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){

		vars := mux.Vars(r)
		id , _ := strconv.Atoi(vars["id"])


		*todos = deleteTodo(id, *todos)
		

		for _ , todo := range *todos {
			fmt.Println(todo.Id)
		}

		fmt.Fprint(w,"Todo Deleted")

	}
}


func addTodoHandler(todos *[]Todo) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){

		var t Todo

		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			fmt.Println("Error logged")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		*todos = addTodo(*todos, t.Id, t.Name, t.Description)

		fmt.Fprint(w,"Todo Added")

	}
}

func main() {  

	todos := make([]Todo,0)
	todos = appendSampleData(todos)

	r := mux.NewRouter()


	r.HandleFunc("/todos", getTodoHandler(&todos)).Methods("GET")

	r.HandleFunc("/todos/delete/{id}", deleteTodoHandler(&todos)).Methods("GET")

	r.HandleFunc("/todos/add", addTodoHandler(&todos)).Methods("POST")

	http.Handle("/", r)


	log.Fatal(http.ListenAndServe(":8080", nil))


}