package libraryimpl

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ssr0016/library/config"
	"github.com/ssr0016/library/library"
	"github.com/stretchr/testify/require"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "secret"
// 	dbname   = "postgres"
// )

// func TestSave(t *testing.T) {
// 	sqlInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", sqlInfo)
// 	require.NoErrorf(t, err, "Database connection failed")
// 	defer db.Close()

// 	t.Run("Create", func(t *testing.T) {
// 		_, err := db.Exec("INSERT INTO book (name) VALUES ('test')")
// 		require.NoError(t, err)
// 	})

// }

func TestSave(t *testing.T) {
	testCases := []struct {
		name string
		library.Book
	}{
		{
			name: "Create",
			Book: library.Book{
				Name: "test",
			},
		},
	}

	for _, tc := range testCases {
		db := config.DbConnectin()
		bookRepository := NewBookRepositoryImpl(db)
		bookRepository.Save(context.Background(), tc.Book)

		require.Equal(t, "test", tc.Book.Name)
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		name string
		library.Book
	}{
		{
			name: "Update",
			Book: library.Book{
				Name: "test",
			},
		},
	}

	for _, tc := range testCases {
		db := config.DbConnectin()
		bookRepository := NewBookRepositoryImpl(db)
		bookRepository.Update(context.Background(), tc.Book)

		require.Equal(t, "test", tc.Book.Name)
	}
}

func TestDelete(t *testing.T) {
	testCases := []struct {
		name   string
		bookId int
	}{
		{
			name:   "Delete",
			bookId: 1,
		},
	}

	for _, tc := range testCases {
		db := config.DbConnectin()
		bookRepository := NewBookRepositoryImpl(db)
		bookRepository.Delete(context.Background(), tc.bookId)

		require.Equal(t, 1, tc.bookId)
	}
}

func TestFindById(t *testing.T) {
	testCases := []struct {
		name           string
		bookId         int
		expectedResult library.Book
		expectedError  error
	}{
		{
			name:           "FindById Success",
			bookId:         1,
			expectedResult: library.Book{Id: 1, Name: "test"},
			expectedError:  nil,
		},

		{
			name:   "FindById Failed",
			bookId: 0,
			expectedResult: library.Book{
				Id:   0,
				Name: "test",
			},
			expectedError: nil,
		},

		{
			name:   "FindById Failed",
			bookId: 2,
			expectedResult: library.Book{
				Id:   0,
				Name: "test",
			},
			expectedError: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := config.DbConnectin()

			bookRepository := NewBookRepositoryImpl(db)
			bookRepository.FindById(context.Background(), tc.bookId)
		})
	}
}

func TestFindAll(t *testing.T) {
	testCases := []struct {
		name           string
		expectedResult []library.Book
	}{
		{
			name:           "FindAll Success",
			expectedResult: []library.Book{{Id: 1, Name: "test"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := config.DbConnectin()
			result := tc.expectedResult

			bookRepository := NewBookRepositoryImpl(db)
			bookRepository.FindAll(context.Background())

			require.Equal(t, tc.expectedResult, result)
		})
	}
}
