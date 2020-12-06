package main

type StorageType int

const (
	Memory StorageType = iota
	Mongodb
)

type Storage interface {
	GetTodo(int) (Todo, error)
	CreateTodo(Todo) error
	DeleteTodo(int) error
	UpdateTodo(Todo) error
	ListTodo() ([]Todo, error)
}

func NewStorage(storageType StorageType) (Storage, error) {
	var stg Storage
	var err error

	switch storageType {
	case Memory:
		stg = new(StorageMemory)

		// case Mongodb:
		// 	stg = new(StorageMongoDB)
	}

	return stg, err

}
