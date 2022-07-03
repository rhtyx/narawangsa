package token

import "time"

type Maker interface {
	CreateToken(username string, userId int64, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
