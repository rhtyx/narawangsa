package users

import (
	"github.com/rhtyx/narawangsa/internal/domain/users"
	"github.com/rhtyx/narawangsa/internal/token"
	"github.com/rhtyx/narawangsa/lib"
)

type handler struct {
	config  lib.Config
	token   token.Maker
	service users.IUsers
}

func NewHandler(service users.IUsers, token token.Maker, config lib.Config) *handler {
	return &handler{
		service: service,
		token:   token,
		config:  config,
	}
}
