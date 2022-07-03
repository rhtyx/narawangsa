package booklists

import (
	"github.com/rhtyx/narawangsa/internal/domain/booklists"
)

type handler struct {
	service booklists.IBooklists
}

func NewHandler(service booklists.IBooklists) *handler {
	return &handler{
		service: service,
	}
}
