package categories

import "github.com/rhtyx/narawangsa/internal/domain/categories"

type handler struct {
	service categories.ICategories
}

func NewHandler(service categories.ICategories) *handler {
	return &handler{
		service: service,
	}
}
