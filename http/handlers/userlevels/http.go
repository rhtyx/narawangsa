package userlevels

import (
	"github.com/rhtyx/narawangsa/internal/domain/userlevels"
	"github.com/rhtyx/narawangsa/internal/domain/users"
	"github.com/rhtyx/narawangsa/internal/token"
)

type handler struct {
	service     userlevels.IUserLevels
	userService users.IUsers
	token       token.Maker
}

func NewHandler(service userlevels.IUserLevels, userService users.IUsers, token token.Maker) *handler {
	return &handler{
		service:     service,
		userService: userService,
		token:       token,
	}
}
