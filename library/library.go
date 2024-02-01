package library

import (
	"context"
)

// Create Repository
// Storage of the entity/model bean in system

type BookStore interface {
	Create(ctx context.Context, cmd *CreateBookCommand) error
	Search(ctx context.Context, query *SearchBookQuery) (*SearchBookResult, error)
	GetById(ctx context.Context, bookId int64) (*BookDTO, error)
	Update(ctx context.Context, cmd *UpdateBookCommand) error
	// FindAll(ctx context.Context) []Book
}

// Create Service
// Business logic
type Service interface {
	Create(ctx context.Context, cmd *CreateBookCommand) error
	Search(ctx context.Context, query *SearchBookQuery) (*SearchBookResult, error)
	GetById(ctx context.Context, bookId int64) (*BookDTO, error)
	Update(ctx context.Context, cmd *UpdateBookCommand) error
	// Delete(ctx context.Context, bookId int)
}
