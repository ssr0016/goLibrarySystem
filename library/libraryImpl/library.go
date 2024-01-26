package libraryimpl

import (
	"context"

	"github.com/ssr0016/library/data/request"
	"github.com/ssr0016/library/data/response"
	"github.com/ssr0016/library/helper"
	"github.com/ssr0016/library/library"
)

type BookServiceImpl struct {
	BookRepository library.BookRepository
}

func NewBookServiceImpl(bookRepository library.BookRepository) library.BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
	}
}

func (b *BookServiceImpl) Create(ctx context.Context, request request.BookCreateRequest) {
	book := library.Book{
		Name: request.Name,
	}

	b.BookRepository.Save(ctx, book)
}

func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicIfError(err)
	b.BookRepository.Delete(ctx, book.Id)
}

func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
	books := b.BookRepository.FindAll(ctx)

	var bookResp []response.BookResponse

	for _, value := range books {
		book := response.BookResponse{
			Id:   value.Id,
			Name: value.Name,
		}

		bookResp = append(bookResp, book)
	}

	return bookResp
}

func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {
	book, err := b.BookRepository.FindById(ctx, bookId)
	helper.PanicIfError(err)
	return response.BookResponse{
		Id:   book.Id,
		Name: book.Name,
	}
}

func (b *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) {
	book, err := b.BookRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	book.Name = request.Name
	b.BookRepository.Update(ctx, book)
}
