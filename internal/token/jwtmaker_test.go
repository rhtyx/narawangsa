package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker("qwertyuiopasdfghjklzxcvbnmqwerty", "qwertyuiopasdfghjklzxcvbnmqwerkl")
	require.NoError(t, err)

	username := "asep"
	userId := 1
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, int64(userId), duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, *payload.Username)
	require.Equal(t, int64(userId), payload.UserId)
	require.WithinDuration(t, issuedAt, *payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, *payload.ExpiredAt, time.Second)
}

func TestJWTRefreshMaker(t *testing.T) {
	maker, err := NewJWTMaker("qwertyuiopasdfghjklzxcvbnmqwerty", "qwertyuiopasdfghjklzxcvbnmqwerkl")
	require.NoError(t, err)

	userId := 1
	token, err := maker.CreateRefreshToken(int64(userId))
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyRefreshToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, int64(userId), payload.UserId)
}

func TestFailNewJWT(t *testing.T) {
	_, err := NewJWTMaker("qwerty", "qwerty")
	require.EqualError(t, err, "invalid key size: must be at least 32 characters")
}
