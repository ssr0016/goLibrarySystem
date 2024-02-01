package libraryimpl

import (
	"context"

	"github.com/ssr0016/library/helper"
	"github.com/ssr0016/library/library"
)

type service struct {
	BookStore library.BookStore
	store     *store
}

func NewService(bookStore library.BookStore) library.Service {
	return &service{
		BookStore: bookStore,
	}
}

func (s *service) Create(ctx context.Context, cmd *library.CreateBookCommand) error {
	result, err := s.store.bookTaken(ctx, 0, cmd.Title)
	helper.PanicIfError(err)

	if len(result) > 0 {
		return library.ErrNameEmpty
	}

	err = s.BookStore.Create(ctx, &library.CreateBookCommand{
		Title:      cmd.Title,
		AuthorID:   cmd.AuthorID,
		CategoryID: cmd.CategoryID,
	})

	helper.PanicIfError(err)

	return nil
}
func (s *service) Search(ctx context.Context, query *library.SearchBookQuery) (*library.SearchBookResult, error) {
	result, err := s.BookStore.Search(ctx, query)
	helper.PanicIfError(err)

	return result, nil
}

func (s *service) GetById(ctx context.Context, bookId int64) (*library.BookDTO, error) {
	result, err := s.BookStore.GetById(ctx, bookId)
	helper.PanicIfError(err)

	if result == nil {
		return nil, library.ErrBookNotFound
	}

	return result, nil
}

func (s *service) Update(ctx context.Context, cmd *library.UpdateBookCommand) error {
	result, err := s.store.bookTaken(ctx, cmd.ID, cmd.Title)
	helper.PanicIfError(err)

	if len(result) > 1 {
		return library.ErrBookNameExisting
	}

	entity := &library.UpdateBookCommand{
		ID:         cmd.ID,
		Title:      cmd.Title,
		AuthorID:   cmd.AuthorID,
		CategoryID: cmd.CategoryID,
	}

	err = s.BookStore.Update(ctx, entity)
	helper.PanicIfError(err)

	return nil
}

// type BookServiceImpl struct {
// 	BookRepository library.BookRepository
// }

// func NewBookServiceImpl(bookRepository library.BookRepository) library.BookService {
// 	return &BookServiceImpl{
// 		BookRepository: bookRepository,
// 	}
// }

// func (b *BookServiceImpl) Create(ctx context.Context, request request.BookCreateRequest) {
// 	book := library.Book{
// 		Name: request.Name,
// 	}

// 	b.BookRepository.Save(ctx, book)
// }

// func (b *BookServiceImpl) Delete(ctx context.Context, bookId int) {
// 	book, err := b.BookRepository.FindById(ctx, bookId)
// 	helper.PanicIfError(err)
// 	b.BookRepository.Delete(ctx, book.Id)
// }

// func (b *BookServiceImpl) FindAll(ctx context.Context) []response.BookResponse {
// 	books := b.BookRepository.FindAll(ctx)

// 	var bookResp []response.BookResponse

// 	for _, value := range books {
// 		book := response.BookResponse{
// 			Id:   value.Id,
// 			Name: value.Name,
// 		}

// 		bookResp = append(bookResp, book)
// 	}

// 	return bookResp
// }

// func (b *BookServiceImpl) FindById(ctx context.Context, bookId int) response.BookResponse {
// 	book, err := b.BookRepository.FindById(ctx, bookId)
// 	helper.PanicIfError(err)
// 	return response.BookResponse{
// 		Id:   book.Id,
// 		Name: book.Name,
// 	}
// }

// func (b *BookServiceImpl) Update(ctx context.Context, request request.BookUpdateRequest) {
// 	book, err := b.BookRepository.FindById(ctx, request.Id)
// 	helper.PanicIfError(err)

// 	book.Name = request.Name
// 	b.BookRepository.Update(ctx, book)
// }
