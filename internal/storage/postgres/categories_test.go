package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateCategory(t *testing.T) {
	err := testQueries.CreateCategory(context.Background(), "Fantasy")
	require.NoError(t, err)
}

func TestGetCategory(t *testing.T) {
	category, err := testQueries.GetCategory(context.Background(), "Fantasy")
	require.NoError(t, err)
	require.NotEmpty(t, category)

	require.Equal(t, "Fantasy", category.Name)
}

func TestListCategories(t *testing.T) {
	testQueries.CreateCategory(context.Background(), "Drama")
	categories, err := testQueries.ListCategories(context.Background(), ListCategoriesParams{
		Limit:  2,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Len(t, categories, 2)
}

func TestUpdateCategory(t *testing.T) {
	category, _ := testQueries.GetCategory(context.Background(), "Drama")
	err := testQueries.UpdateCategory(context.Background(), UpdateCategoryParams{
		Name:      "Korean Drama",
		UpdatedAt: time.Now(),
		ID:        category.ID,
	})
	require.NoError(t, err)
}

func TestDeleteCategory(t *testing.T) {
	categoryFantasy, _ := testQueries.GetCategory(context.Background(), "Fantasy")
	categoryKoreanDrama, _ := testQueries.GetCategory(context.Background(), "Korean Drama")
	err1 := testQueries.DeleteCategory(context.Background(), categoryFantasy.ID)
	err2 := testQueries.DeleteCategory(context.Background(), categoryKoreanDrama.ID)
	require.NoError(t, err1)
	require.NoError(t, err2)
}
