package main

// Todo defines the structure for a Item in the Todo list
// swagger:model
type Todo struct {
	// the id for the Todo Item
	//
	// required: false
	// min: 1
	Id int `bson:"_id,omitempty"`

	// the name for this Todo Item
	//
	// required: true
	// max length: 255
	Name string `bson:"name" validate:"required"`

	// the descirption of the Todo item
	//
	// required: true
	// max length: 255
	// min length: 1
	Description string `bson:"description" validate:"required"`
}
