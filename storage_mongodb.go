package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageMongodb struct {
	client *mongo.Client
	ctx    context.Context
}

func NewStorageMongodb() (StorageMongodb, error) {
	var mongodb StorageMongodb
	var err error

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

	mongodb.client, err = mongo.NewClient(options.Client().ApplyURI(config.ConnectionString))
	if err != nil {
		fmt.Println("Error detected")
		log.Fatal(err)
	}
	mongodb.ctx = context.Background()
	err = mongodb.client.Connect(mongodb.ctx)
	if err != nil {
		log.Fatal(err)
	}
	return mongodb, err
}

func (mongodb StorageMongodb) DeleteStorage() {
	log.Print("Deleting mongo Client")
	defer mongodb.client.Disconnect(mongodb.ctx)
}

func (mongodb StorageMongodb) GetTodo(id int) (Todo, error) {

	todoDatabase := mongodb.client.Database("tododb")
	todoCollection := todoDatabase.Collection("todos")

	var todo Todo
	err := todoCollection.FindOne(mongodb.ctx, bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}
	return todo, err

}

func (mongodb StorageMongodb) DeleteTodo(id int) error {

	todoDatabase := mongodb.client.Database("tododb")
	todoCollection := todoDatabase.Collection("todos")

	_, err := todoCollection.DeleteMany(mongodb.ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}

	return err

}

func (mongodb StorageMongodb) CreateTodo(todo Todo) error {

	todoDatabase := mongodb.client.Database("tododb")
	todoCollection := todoDatabase.Collection("todos")

	result, err := todoCollection.InsertOne(mongodb.ctx, todo)

	if err == nil {
		fmt.Println(result.InsertedID)
	} else {
		fmt.Println(err)
	}

	return err

}

func (mongodb StorageMongodb) UpdateTodo(update_todo Todo) error {

	todoDatabase := mongodb.client.Database("tododb")
	todoCollection := todoDatabase.Collection("todos")

	result_update, err := todoCollection.UpdateMany(
		mongodb.ctx,
		bson.M{"_id": update_todo.Id},
		bson.D{
			{"$set", bson.D{{"description", update_todo.Description}, {"name", update_todo.Name}}},
		},
	)

	if err == nil {
		fmt.Println("Updated Count ", result_update.ModifiedCount)
	} else {
		fmt.Println(err)
	}

	return nil

}

func (mongodb StorageMongodb) ListTodo() ([]Todo, error) {

	todoDatabase := mongodb.client.Database("tododb")
	todoCollection := todoDatabase.Collection("todos")

	var todolist []Todo

	cursor, err := todoCollection.Find(
		mongodb.ctx,
		bson.D{{}},
	)

	if err != nil {
		log.Fatal(err)
	} else {
		for cursor.Next(mongodb.ctx) {
			var todo Todo
			err := cursor.Decode(&todo)
			if err != nil {
				log.Fatal(err)
			}
			todolist = append(todolist, todo)
		}
	}
	return todolist, err

}
