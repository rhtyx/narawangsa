package postgres

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateReadConfirmation(t *testing.T) {
	createBookListF()
	user := getUserF()
	booklist, _ := testQueries.ListBookList(context.Background(), user.ID)
	argReadConfirmation := CreateReadConfirmationParams{
		BookListID: booklist[0].ID,
		PagesRead:  5,
	}

	err := testQueries.CreateReadConfirmation(context.Background(), argReadConfirmation)
	require.NoError(t, err)
}

func TestListReadConfirmations(t *testing.T) {
	user := getUserF()
	booklist, _ := testQueries.ListBookList(context.Background(), user.ID)
	argListReadConfirmation := ListReadConfirmationsParams{
		BookListID: booklist[0].ID,
		Limit:      2,
	}

	listRC, err := testQueries.ListReadConfirmations(context.Background(), argListReadConfirmation)
	require.NoError(t, err)
	require.Len(t, listRC, 1)

	testQueries.truncate()
}
