// reference : https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-NewWithClaims-RegisteredClaims

package token

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var (
	key                   = []byte("hello")
	accessTokenExpiredAt  = time.Now().UTC().Add(time.Hour * 1)
	refreshTokenExpiredAt = time.Now().UTC().Add(time.Hour * 720)
)

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// validate expired time
		return nil
	}
	return err
}

func GenerateAccessToken(userID int) (string, time.Time, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(accessTokenExpiredAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", accessTokenExpiredAt, err
	}
	return tokenString, accessTokenExpiredAt, nil
}

func GenerateRefreshToken(userID int) (string, time.Time, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(refreshTokenExpiredAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", refreshTokenExpiredAt, err
	}
	return tokenString, refreshTokenExpiredAt, nil
}
