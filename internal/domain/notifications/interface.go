package notifications

import "context"

type INotifications interface {
	SendNotifications(ctx context.Context, name, email, quotes string) error
}
