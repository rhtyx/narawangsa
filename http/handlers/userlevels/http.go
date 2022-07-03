package userlevels

import (
	"github.com/rhtyx/narawangsa/internal/domain/userlevels"
	"github.com/rhtyx/narawangsa/internal/token"
)

type handler struct {
	service userlevels.IUserLevels
	token   token.Maker
}

func NewHandler(service userlevels.IUserLevels, token token.Maker) *handler {
	return &handler{
		service: service,
		token:   token,
	}
}
