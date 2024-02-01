package libraryimpl

import (
	"context"
	"database/sql"

	"github.com/ssr0016/library/helper"
	"github.com/ssr0016/library/library"
)

type store struct {
	Db *sql.DB
}

func NewStore(Db *sql.DB) *store {
	return &store{
		Db: Db,
	}
}

func (s *store) Create(ctx context.Context, entity *library.CreateBookCommand) error {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	rawSQL := `
		INSERT INTO book(
			title,
			author_id,
			category_id
			published_at
		)
		VALUES(
			:title,
			:author_id,
			:category_id,
			:published_at
		)
		`
	_, err = tx.ExecContext(ctx, rawSQL, entity)
	helper.PanicIfError(err)
	return nil

}

func (s *store) Search(ctx context.Context, query *library.SearchBookQuery) (*library.SearchBookResult, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	rawSQL := `
		SELECT
			id,
			title,
			author,
			category_id,
			published_at
		FROM book
	`
	result, errQuery := tx.QueryContext(ctx, rawSQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var books []library.SearchBookQuery

	for result.Next() {
		book := library.SearchBookQuery{}
		err := result.Scan(query)
		helper.PanicIfError(err)

		books = append(books, book)
	}

	searchResult := &library.SearchBookResult{
		TotalCount: int64(len(books)),
		Page:       query.Page,
		PerPage:    query.PerPage,
	}

	return searchResult, nil
}

func (s *store) GetById(ctx context.Context, bookId int64) (*library.BookDTO, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	rawSQL := `
		SELECT 
			id,
			title,
			author_id,
			category_id,
			published_at
		FROM book
		WHERE id = ?
	`

	result, errQuery := tx.QueryContext(ctx, rawSQL, bookId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	book := library.BookDTO{}

	if result.Next() {
		err := result.Scan(&book)
		helper.PanicIfError(err)
		return &book, nil
	} else {

		return nil, library.ErrBookNotFound
	}
}

func (s *store) Update(ctx context.Context, cmd *library.UpdateBookCommand) error {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	rawSQL := `
		UPDATE book
		SET title = :title,
		author_id = :author_id,
		category_id = :category_id,
		published_at = :published_at
		WHERE id = :id
	`
	_, err = tx.ExecContext(ctx, rawSQL, cmd)
	helper.PanicIfError(err)

	return nil
}

func (s *store) bookTaken(ctx context.Context, id int64, title string) ([]*library.Book, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	rawSQL := `
		SELECT
			id,
			title,
			author_id,
			category_id,
			published_at
		FROM book
		WHERE 
			id = ? OR
			title = ?
	`

	rows, errQuery := tx.QueryContext(ctx, rawSQL, id, title)
	helper.PanicIfError(errQuery)
	defer rows.Close()

	var result []*library.Book

	for rows.Next() {
		book := library.Book{}
		err := rows.Scan(&book)
		helper.PanicIfError(err)
		result = append(result, &book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
