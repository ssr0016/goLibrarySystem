package library

import (
	"context"
)

// Create Repository
// Storage of the entity/model bean in system

type BookRepository interface {
	Save(ctx context.Context, book Book)
	Update(ctx context.Context, book Book)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) (Book, error)
	FindAll(ctx context.Context) []Book
}
