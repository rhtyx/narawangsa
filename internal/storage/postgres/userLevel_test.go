package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateUserLevel(t *testing.T) {
	testQueries.CreateUser(context.Background(), argUser)
	user, _ := testQueries.GetUser(context.Background(), argUpdateUser.Username)
	err := testQueries.CreateUserLevel(context.Background(), user.ID)
	require.NoError(t, err)
}

func TestGetUserLevel(t *testing.T) {
	user, _ := testQueries.GetUser(context.Background(), argUpdateUser.Username)
	userLevel, err := testQueries.GetUserLevel(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, userLevel)

	require.Equal(t, int32(0), userLevel.Level)
}

func TestUpdateUserLevelsLevel(t *testing.T) {
	user, _ := testQueries.GetUser(context.Background(), argUpdateUser.Username)
	err := testQueries.UpdateUserLevelsLevel(context.Background(), UpdateUserLevelsLevelParams{
		Level:     1,
		UpdatedAt: time.Now(),
		UserID:    user.ID,
	})
	require.NoError(t, err)

	testQueries.DeleteUser(context.Background(), argUser.Username)

	testQueries.truncate()
}
