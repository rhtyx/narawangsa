package userlevels

import (
	"github.com/rhtyx/narawangsa/internal/domain/userlevels"
)

type handler struct {
	service userlevels.IUserLevels
}

func NewHandler(service userlevels.IUserLevels) *handler {
	return &handler{
		service: service,
	}
}
