package booklists

import (
	"github.com/rhtyx/narawangsa/internal/domain/booklists"
	"github.com/rhtyx/narawangsa/internal/token"
)

type handler struct {
	service booklists.IBooklists
	token   token.Maker
}

func NewHandler(service booklists.IBooklists, token token.Maker) *handler {
	return &handler{
		service: service,
		token:   token,
	}
}
