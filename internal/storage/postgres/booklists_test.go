package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createBookListF() error {
	createUserF()
	createBookAndCategoryF()
	user := getUserF()
	bookIDAndCategoryID := getBookIDandCategoryID()
	var argBookList = CreateBookListParams{
		UserID:    user.ID,
		BookID:    bookIDAndCategoryID.BookID,
		IsRead:    false,
		PagesRead: int32(0),
		EndDate:   time.Now(),
	}

	err := testQueries.CreateBookList(context.Background(), argBookList)
	return err
}

func TestCreateBookList(t *testing.T) {
	err := createBookListF()
	require.NoError(t, err)
}

func TestListBookList(t *testing.T) {
	user := getUserF()
	booklist, err := testQueries.ListBookList(context.Background(), user.ID)
	require.NoError(t, err)
	require.Len(t, booklist, 1)
}

func TestUpdateBookList(t *testing.T) {
	user := getUserF()
	booklist, _ := testQueries.ListBookList(context.Background(), user.ID)
	argUpdateBookList := UpdateBookListParams{
		IsRead:    false,
		PagesRead: 5,
		EndDate:   time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		BookID:    booklist[0].BookID,
	}

	err := testQueries.UpdateBookList(context.Background(), argUpdateBookList)
	require.NoError(t, err)
}

func TestDeleteBookList(t *testing.T) {
	bookIDAndCategoryID := getBookIDandCategoryID()
	err := testQueries.DeleteBookList(context.Background(), bookIDAndCategoryID.BookID)
	require.NoError(t, err)

	testQueries.truncate()
}
