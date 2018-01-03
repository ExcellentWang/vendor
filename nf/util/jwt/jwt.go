package jwt

import (
	"encoding/base64"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

var (
	key []byte
)

func init() {
	var err error
	key, err = base64.StdEncoding.DecodeString("AAaVBFj6LXY5L/u+wccacWSl1ek54v7mY1mj17U9iXi5zPZ/+6QIfTL7qwCGtn0ch8bql6zexm9Xv9+1deZjBQ==")
	if err != nil {
		panic(err)
	}
}

func GenToken(subject string, expiration time.Time) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims["sub"] = subject
	token.Claims["exp"] = expiration.Unix()
	return token.SignedString(key)
}

func ParseToken(tokenString string) (claims claims, err error) {
	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}

type claims map[string]interface{}

func (claims claims) GetSubject() (string, error) {
	if val, ok := claims["sub"].(string); ok {
		return val, nil
	}
	return "", fmt.Errorf("the subject %v is not a string", claims["sub"])
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return key, nil
}
