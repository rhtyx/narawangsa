package categorybooks

import (
	"github.com/rhtyx/narawangsa/internal/domain/categorybooks"
)

type handler struct {
	service categorybooks.ICategoryBooks
}

func NewHandler(service categorybooks.ICategoryBooks) *handler {
	return &handler{
		service: service,
	}
}
