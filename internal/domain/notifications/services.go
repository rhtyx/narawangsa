package notifications

import (
	"context"

	"github.com/trycourier/courier-go/v2"
)

func NewNotificationsService() INotifications {
	return &message{}
}

func (s *message) SendNotifications(ctx context.Context, name, email, quotes string) error {
	client := courier.CreateClient("pk_prod_5MWJ36B934M1VJHA86AZ3CYPRXP9", nil)

	_, err := client.SendMessage(
		ctx,
		courier.SendMessageRequestBody{
			Message: map[string]interface{}{
				"to": map[string]string{
					"email": email,
				},
				"template": "3RR2FE9B2X4BXXQ4G15XDAYXF05A",
				"data": map[string]string{
					"name":   name,
					"quotes": quotes,
				},
			},
		},
	)

	return err
}
