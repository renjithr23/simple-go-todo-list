package main

type Todo struct {
	Id          int    `bson:"_id,omitempty"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
}

// func deleteTodo(id int, todos []Todo) []Todo {

// 	for j, todo := range todos {
// 		if todo.Id == id {
// 			copy(todos[j:], todos[j+1:])
// 			todos = todos[:len(todos)-1]
// 			break
// 		}
// 	}

// 	return todos

// }

// func addTodo(todos []Todo, id int, name string, description string) []Todo {

// 	todos = append(todos, Todo{id, name, description, time.Now()})
// 	return todos

// }

// func appendSampleData(todos []Todo) []Todo {

// 	todos = append(todos, Todo{1, "name1", "desc1", time.Now()})
// 	todos = append(todos, Todo{2, "name2", "desc2", time.Now()})
// 	todos = append(todos, Todo{3, "name3", "desc3", time.Now()})
// 	todos = append(todos, Todo{4, "name4", "desc4", time.Now()})
// 	todos = append(todos, Todo{5, "name5", "desc5", time.Now()})

// 	return todos

// }
