package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID        uuid.UUID  `json:"id"`
	Username  *string    `json:"username"`
	UserId    int64      `json:"user_id"`
	IssuedAt  *time.Time `json:"issued_at"`
	ExpiredAt *time.Time `json:"expired_at"`
}

func NewPayload(username *string, userId int64, duration *time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	if username != nil && duration != nil {
		now := time.Now()
		later := time.Now().Add(*duration)

		payload := &Payload{
			ID:        tokenID,
			UserId:    userId,
			Username:  username,
			IssuedAt:  &now,
			ExpiredAt: &later,
		}

		return payload, nil
	}

	payload := &Payload{
		ID:     tokenID,
		UserId: userId,
	}

	return payload, nil
}

func (p *Payload) Valid() error {
	if p.ExpiredAt != nil && time.Now().After(*p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
