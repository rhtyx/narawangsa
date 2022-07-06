package notifications

import "context"

type INotifications interface {
	SendNotifications(ctx context.Context) error
}
