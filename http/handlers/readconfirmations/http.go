package readconfirmations

import (
	"github.com/rhtyx/narawangsa/internal/domain/booklists"
	"github.com/rhtyx/narawangsa/internal/domain/readconfirmations"
)

type handler struct {
	service         readconfirmations.IReadConfirmations
	booklistService booklists.IBooklists
}

func NewHandler(service readconfirmations.IReadConfirmations, booklistService booklists.IBooklists) *handler {
	return &handler{
		service:         service,
		booklistService: booklistService,
	}
}
