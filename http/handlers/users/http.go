package users

import (
	"github.com/rhtyx/narawangsa/internal/domain/users"
	"github.com/rhtyx/narawangsa/internal/token"
)

type handler struct {
	token   token.Maker
	service users.IUsers
}

func NewHandler(service users.IUsers, token token.Maker) *handler {
	return &handler{
		service: service,
		token:   token,
	}
}
