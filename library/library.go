package library

import (
	"context"

	"github.com/ssr0016/library/data/request"
	"github.com/ssr0016/library/data/response"
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

// Create Service
// Business logic
type BookService interface {
	Create(ctx context.Context, request request.BookCreateRequest)
	Update(ctx context.Context, request request.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) response.BookResponse
	FindAll(ctx context.Context) []response.BookResponse
}
