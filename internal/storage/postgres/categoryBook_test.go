package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createBookAndCategoryF() {
	testQueries.CreateBook(context.Background(), argBook1)
	testQueries.CreateCategory(context.Background(), "Fantasy")
}

func getBookIDandCategoryID() CreateBookCategoryParams {
	book, _ := testQueries.GetBook(context.Background(), "Harry Potter")
	category, _ := testQueries.GetCategory(context.Background(), "Fantasy")

	return CreateBookCategoryParams{
		BookID:     book.ID,
		CategoryID: category.ID,
	}
}

func TestCreateBookCategory(t *testing.T) {
	createBookAndCategoryF()
	bookCategoryParams := getBookIDandCategoryID()

	err := testQueries.CreateBookCategory(context.Background(), bookCategoryParams)
	require.NoError(t, err)
}

func TestDeleteBookCategory(t *testing.T) {
	bookCategoryParams := getBookIDandCategoryID()
	err := testQueries.DeleteBookCategory(context.Background(), DeleteBookCategoryParams(bookCategoryParams))
	require.NoError(t, err)

	testQueries.truncate()
}
