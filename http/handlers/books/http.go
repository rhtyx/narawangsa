package books

import "github.com/rhtyx/narawangsa/internal/domain/books"

type handler struct {
	service books.IBooks
}

func Newhandler(service books.IBooks) *handler {
	return &handler{
		service: service,
	}
}
