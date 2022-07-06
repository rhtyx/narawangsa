package notifications

import "github.com/rhtyx/narawangsa/internal/domain/notifications"

type handler struct {
	service notifications.INotifications
}

func NewService(service notifications.INotifications) *handler {
	return &handler{
		service: service,
	}
}
