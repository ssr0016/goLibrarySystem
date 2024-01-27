package libraryimpl

import (
	"context"
	"database/sql"

	"github.com/ssr0016/library/helper"
	"github.com/ssr0016/library/library"
)

type BookRepositoryImpl struct {
	Db *sql.DB
}

func NewBookRepositoryImpl(Db *sql.DB) library.BookRepository {
	return &BookRepositoryImpl{
		Db: Db,
	}
}

func (b *BookRepositoryImpl) Save(ctx context.Context, book library.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	SQL := `
		INSERT INTO
		book
		(name)
		VALUES
		($1)
	`
	_, err = tx.ExecContext(ctx, SQL, book.Name)
	helper.PanicIfError(err)
}
func (b *BookRepositoryImpl) Update(ctx context.Context, book library.Book) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	SQL := `
		UPDATE
		book
		SET
		name = $1
		WHERE
		id = $2
	`
	_, err = tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helper.PanicIfError(err)
}
func (b *BookRepositoryImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	SQL := `
		DELETE
		FROM
		book
		WHERE
		id = $1
	`
	_, errExec := tx.ExecContext(ctx, SQL, bookId)
	helper.PanicIfError(errExec)

}
func (b *BookRepositoryImpl) FindById(ctx context.Context, bookId int) (library.Book, error) {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	SQL := `
		SELECT 
		id, name
		FROM
		book
		WHERE
		id = $1
	`

	result, errQuery := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	book := library.Book{}

	if result.Next() {
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)
		return book, nil
	} else {

		return book, library.ErrBookIdNotFound
	}
}

func (b *BookRepositoryImpl) FindAll(ctx context.Context) []library.Book {
	tx, err := b.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	SQL := `
		SELECT
		id, name
		FROM
		book
	`
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var books []library.Book

	for result.Next() {
		book := library.Book{}
		err := result.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)

		books = append(books, book)
	}

	return books

}
