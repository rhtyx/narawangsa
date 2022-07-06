package postgres

import (
	"context"
	"testing"

	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/stretchr/testify/require"
)

func refreshToken() string {
	maker, _ := token.NewJWTMaker("qwertyuiopasdfghjklzxcvbnmqwerty", "qwertyuiopasdfghjklzxcvbnmqwerkl")

	userId := 1
	token, _ := maker.CreateRefreshToken(int64(userId))
	return token
}

func TestCreateRefreshToken(t *testing.T) {
	testQueries.truncate()
	token := refreshToken()
	err := testQueries.CreateRefreshToken(context.Background(), token)
	require.NoError(t, err)
}

func TestGetRefreshToken(t *testing.T) {
	testQueries.truncate()
	token := refreshToken()
	_ = testQueries.CreateRefreshToken(context.Background(), token)

	returnedToken, err := testQueries.GetRefreshToken(context.Background(), token)
	require.NoError(t, err)

	require.Equal(t, token, returnedToken)
}

func TestDeleteRefreshToken(t *testing.T) {
	testQueries.truncate()
	token := refreshToken()
	_ = testQueries.CreateRefreshToken(context.Background(), token)

	err := testQueries.DeleteRefreshToken(context.Background(), token)
	require.NoError(t, err)
}
