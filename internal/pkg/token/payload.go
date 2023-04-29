package token

import (
	"strconv"
	"time"
)

type Payload struct {
	Subject   string
	ExpiredAt time.Time
	IssuedAt  time.Time
}

func NewPayload(userID int, duration time.Duration) (*Payload, error) {
	strUserID := strconv.Itoa(userID)
	payload := &Payload{
		Subject:   strUserID,
		ExpiredAt: time.Now().Add(duration),
		IssuedAt:  time.Now(),
	}

	return payload, nil
}
