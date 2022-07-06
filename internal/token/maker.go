package token

import "time"

type Maker interface {
	CreateToken(username string, userId int64, duration time.Duration) (string, error)
	CreateRefreshToken(userId int64) (string, error)
	VerifyToken(token string) (*Payload, error)
	VerifyRefreshToken(token string) (*Payload, error)
}
