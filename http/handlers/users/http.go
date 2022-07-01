package users

import "github.com/rhtyx/narawangsa/internal/domain/users"

type handler struct {
	service users.IUsers
}

func NewHandler(service users.IUsers) *handler {
	return &handler{
		service: service,
	}
}
