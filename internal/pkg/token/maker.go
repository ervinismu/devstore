package token

import "time"

type Maker interface {
	GenerateRefreshToken(userID int) (string, time.Time, error)
	GenerateAccessToken(userID int) (string, time.Time, error)
	VerifyToken(tokenString string) (string, error)
}
