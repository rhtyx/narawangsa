package notifications

import "context"

func NewNotificationsService(name, email, mess string) INotifications {
	return &message{
		name:    name,
		email:   email,
		message: mess,
	}
}

func (s *message) SendNotifications(ctx context.Context, message message) error
