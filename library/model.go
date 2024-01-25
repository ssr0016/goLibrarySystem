package library

import "errors"

// Model/Entity typically represents a table in the application's database

var (
	ErrBookIdNotFound = errors.New("book id not found")
)

type Book struct {
	Id   int
	Name string
}
