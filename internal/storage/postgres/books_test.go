package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var argBook1 = CreateBookParams{
	Title:    "Harry Potter",
	Author:   "J. K. Rowling",
	Year:     "2003",
	Pages:    1004,
	Synopsis: "sadkljfbasdkjfbaskldfansdknlfjasbdfkaljsbdflkasdjbfiaewufbaerogabnerlkj",
}

var argBook2 = CreateBookParams{
	Title:    "Harry Potter2",
	Author:   "J. K. Rowling",
	Year:     "2004",
	Pages:    1039,
	Synopsis: "asdlkjfbasdjkfbaeofbaslkdjfbasdkbfalskdjfbasdk gaskldj fkasdjc aslkdcalkesubfawoe",
}

func TestCreateBook(t *testing.T) {
	err := testQueries.CreateBook(context.Background(), argBook1)
	require.NoError(t, err)
}

func TestGetBook(t *testing.T) {
	book, err := testQueries.GetBook(context.Background(), "Harry Potter")
	require.NoError(t, err)
	require.NotEmpty(t, book)

	require.Equal(t, argBook1.Title, book.Title)
	require.Equal(t, argBook1.Author, book.Author)
	require.Equal(t, argBook1.Year, book.Year)
	require.Equal(t, argBook1.Pages, book.Pages)
	require.Equal(t, argBook1.Synopsis, book.Synopsis)
}

func TestUpdateBook(t *testing.T) {
	book, _ := testQueries.GetBook(context.Background(), "Harry Potter")
	var argUpdateBook = UpdateBookParams{
		Title:     "Harry Potter",
		Author:    "J. K. Rowling",
		Year:      "2005",
		Pages:     2000,
		Synopsis:  "asdlkjfbasdjkfbaeofbaslkdjfbasdkbfalskdjfbasdk gaskldj fkasdjc aslkdcalkesubfawoe",
		UpdatedAt: time.Now(),
		ID:        book.ID,
	}
	err := testQueries.UpdateBook(context.Background(), argUpdateBook)
	require.NoError(t, err)
}

func TestListBooks(t *testing.T) {
	testQueries.CreateBook(context.Background(), argBook2)
	books, err := testQueries.ListBooks(context.Background(), 2)
	require.NoError(t, err)
	require.Len(t, books, 2)
}

func TestDeleteBook(t *testing.T) {
	var err error
	books, _ := testQueries.ListBooks(context.Background(), 2)
	for _, book := range books {
		err = testQueries.DeleteBook(context.Background(), book.ID)
	}
	require.NoError(t, err)

	testQueries.truncate()
}
