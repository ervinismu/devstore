// reference : https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-NewWithClaims-RegisteredClaims

package token

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

type JWTMaker struct {
	accessTokenKey  string
	refreshTokenKey string
}

var (
	accessTokenKey        = []byte("hello")
	refreshTokenKey       = []byte("hello")
	accessTokenExpiredAt  = time.Now().UTC().Add(time.Hour * 1)
	refreshTokenExpiredAt = time.Now().UTC().Add(time.Hour * 720)
)

func NewJWTMaker(accessTokenKey string, refreshTokenKey string) *JWTMaker {
	return &JWTMaker{
		accessTokenKey:  accessTokenKey,
		refreshTokenKey: refreshTokenKey,
	}
}

func (maker *JWTMaker) GenerateRefreshToken(userID int, duration time.Duration) (string, time.Time, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(refreshTokenExpiredAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(refreshTokenKey)
	if err != nil {
		return "", refreshTokenExpiredAt, err
	}
	return tokenString, refreshTokenExpiredAt, nil
}

func (maker *JWTMaker) GenerateAccessToken(userID int) (string, time.Time, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(accessTokenExpiredAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   fmt.Sprintf("%d", userID),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(accessTokenKey)
	if err != nil {
		return "", accessTokenExpiredAt, err
	}
	return tokenString, accessTokenExpiredAt, nil
}

func (maker *JWTMaker) VerifyToken(tokenString string) (string, error) {

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return accessTokenKey, nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		log.Error(fmt.Errorf("verify token : %w", err))
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		sub := fmt.Sprint(claims["sub"])
		return sub, nil
	}
	return "", err
}

// var (
// 	accessTokenKey        = []byte("hello")
// 	refreshTokenKey       = []byte("hello")
// 	accessTokenExpiredAt  = time.Now().UTC().Add(time.Hour * 1)
// 	refreshTokenExpiredAt = time.Now().UTC().Add(time.Hour * 720)
// )

// func VerifyToken(tokenString string) (string, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return accessTokenKey, nil
// 	})
// 	if err != nil {
// 		log.Error(fmt.Errorf("verify token : %w", err))
// 		return "", err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		sub := fmt.Sprint(claims["sub"])
// 		return sub, nil
// 	}
// 	return "", err
// }

// func GenerateAccessToken(userID int) (string, time.Time, error) {
// 	claims := jwt.RegisteredClaims{
// 		ExpiresAt: jwt.NewNumericDate(accessTokenExpiredAt),
// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
// 		Subject:   fmt.Sprintf("%d", userID),
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(accessTokenKey)
// 	if err != nil {
// 		return "", accessTokenExpiredAt, err
// 	}
// 	return tokenString, accessTokenExpiredAt, nil
// }

// func GenerateRefreshToken(userID int) (string, time.Time, error) {
// 	claims := jwt.RegisteredClaims{
// 		ExpiresAt: jwt.NewNumericDate(refreshTokenExpiredAt),
// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
// 		Subject:   fmt.Sprintf("%d", userID),
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(refreshTokenKey)
// 	if err != nil {
// 		return "", refreshTokenExpiredAt, err
// 	}
// 	return tokenString, refreshTokenExpiredAt, nil
// }
