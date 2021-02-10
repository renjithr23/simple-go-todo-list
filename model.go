package main

type Todo struct {
	Id          int    `bson:"_id,omitempty"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
}
