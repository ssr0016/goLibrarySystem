package library

import (
	"time"

	"github.com/ssr0016/library/api/errors"
)

// Model/Entity typically represents a table in the application's database

// type Book struct {
// 	Id   int
// 	Name string
// }

var (
	ErrNameEmpty         = errors.New("book.name-empty", "Name is empty")
	ErrBookNameExisting  = errors.New("book.book-existing", "Name already exists")
	ErrAuthorIDInvalid   = errors.New("book.author-id-invalid", "Author ID is invalid")
	ErrCategoryIDInvalid = errors.New("book.category-id-invalid", "Category is invalid")
	ErrBookNotFound      = errors.New("book.book-not-found", "Book not found")
)

type Book struct {
	ID          int64  `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	AuthorID    string `db:"author_id" json:"author_id"`
	CategoryID  string `db:"category_id" json:"category_id"`
	PublishedAt string `db:"published_at" json:"published_at"`
}

type BookDTO struct {
	ID          int64       `id:"id" json:"id"`
	Title       string      `db:"title" json:"title"`
	AuthorID    string      `db:"author_id" json:"author_id"`
	CategoryID  string      `db:"category_id" json:"category_id"`
	PublishedAt string      `db:"published_at" json:"published_at"`
	Author      []*Author   `json:"author,omitempty"`
	Category    []*Category `json:"category,omitempty"`
}

type Category struct {
	ID   int64  `db:"category_id" json:"category_id"`
	Name string `db:"name" json:"name"`
}

type CategoryDTO struct {
	ID    int64   `db:"category_id" json:"category_id"`
	Name  string  `db:"name" json:"name"`
	Books []*Book `json:"books"`
}

type Author struct {
	ID   int64  `db:"author_id" json:"author_id"`
	Name string `db:"name" json:"name"`
}

type AuthorDTO struct {
	ID    int64   `db:"author_id" json:"author_id"`
	Name  string  `db:"name" json:"name"`
	Books []*Book `json:"books"`
}

type CreateBookCommand struct {
	Title      string `json:"name"`
	AuthorID   string `json:"author_id"`
	CategoryID string `json:"category_id"`
}

type SearchBookCommand struct {
	Title      string `schema:"title"`
	AuthorID   int64  `schema:"author_id"`
	CategoryID int64  `schema:"category_id"`
}

type SearchBookQuery struct {
	Title      string     `schema:"title"`
	AuthorID   int64      `schema:"author_id"`
	CategoryID int64      `schema:"category_id"`
	DateFrom   *time.Time `schema:"date_from"`
	Dateto     *time.Time `schema:"date_to"`
	Page       int        `schema:"page"`
	PerPage    int        `schema:"page_size"`
}

type SearchBookResult struct {
	TotalCount int64   `json:"total_count"`
	Books      []*Book `json:"result"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}

type UpdateBookCommand struct {
	ID         int64
	Title      string `json:"name"`
	AuthorID   string `json:"author_id"`
	CategoryID string `json:"category_id"`
}

func (cmd *CreateBookCommand) Validate() error {
	if len(cmd.Title) == 0 {
		return ErrNameEmpty
	}

	if cmd.AuthorID == "" {
		return ErrAuthorIDInvalid
	}

	if cmd.CategoryID == "" {
		return ErrCategoryIDInvalid
	}

	return nil
}

func (cmd *UpdateBookCommand) Validate() error {
	if len(cmd.Title) == 0 {
		return ErrNameEmpty
	}

	if cmd.AuthorID == "" {
		return ErrAuthorIDInvalid
	}

	if cmd.CategoryID == "" {
		return ErrCategoryIDInvalid
	}

	return nil

}
