package readconfirmations

import "github.com/rhtyx/narawangsa/internal/domain/readconfirmations"

type handler struct {
	service readconfirmations.IReadConfirmations
}

func NewHandler(service readconfirmations.IReadConfirmations) *handler {
	return &handler{
		service: service,
	}
}
