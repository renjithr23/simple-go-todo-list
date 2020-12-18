package main

type StorageType string

const (
	Memory  StorageType = "Memory"
	Mongodb StorageType = "Mongodb"
)

type Storage interface {
	GetTodo(int) (Todo, error)
	CreateTodo(Todo) error
	DeleteTodo(int) error
	UpdateTodo(Todo) error
	ListTodo() ([]Todo, error)
	DeleteStorage()
}

func NewStorage(storageType StorageType) (Storage, error) {
	var stg Storage
	var err error

	switch storageType {
	case Memory:
		stg = new(StorageMemory)

	case Mongodb:
		stg, err = NewStorageMongodb()
	}
	return stg, err
}
