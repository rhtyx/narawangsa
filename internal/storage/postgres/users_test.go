package postgres

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var argUser = CreateUserParams{
	Name:     "Tony",
	Username: "tony",
	Email:    "tony@mail.com",
	Password: "tony123",
}

var argUserPassword = UpdatePasswordUserParams{
	Password:  "tony321",
	UpdatedAt: time.Now(),
	Username:  "tony",
}

var argUpdateUser = UpdateUserParams{
	Name:      "Tony T",
	Email:     "tonyT@mail.com",
	UpdatedAt: time.Now(),
	Username:  "tony",
}

func createUserF() {
	_, err := testQueries.CreateUser(context.Background(), argUser)
	if err != nil {
		fmt.Println(err)
	}
}

func getUserF() User {
	user, _ := testQueries.GetUser(context.Background(), argUser.Username)
	return user
}

func TestCreateUser(t *testing.T) {
	id, err := testQueries.CreateUser(context.Background(), argUser)
	require.NoError(t, err)

	user, _ := testQueries.GetUser(context.Background(), argUser.Username)
	require.Equal(t, user.ID, id)
}

func TestGetUser(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), argUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, argUser.Name, user.Name)
	require.Equal(t, argUser.Username, user.Username)
	require.Equal(t, argUser.Email, user.Email)
}

func TestUpdatePasswordUser(t *testing.T) {
	err := testQueries.UpdatePasswordUser(context.Background(), argUserPassword)
	require.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	user, err := testQueries.UpdateUser(context.Background(), argUpdateUser)
	require.NoError(t, err)
	require.Equal(t, argUpdateUser.Name, user.Name)
	require.Equal(t, argUpdateUser.Email, user.Email)
}

func TestDeleteUser(t *testing.T) {
	err := testQueries.DeleteUser(context.Background(), argUser.Username)
	require.NoError(t, err)
}
